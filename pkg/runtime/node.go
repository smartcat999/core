package runtime

import (
	"context"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/tkeel-io/core/pkg/dispatch"
	zfield "github.com/tkeel-io/core/pkg/logger"
	"github.com/tkeel-io/core/pkg/mapper"
	"github.com/tkeel-io/core/pkg/placement"
	"github.com/tkeel-io/core/pkg/repository/dao"
	"github.com/tkeel-io/core/pkg/types"
	"github.com/tkeel-io/core/pkg/util"
	xkafka "github.com/tkeel-io/core/pkg/util/kafka"
	"github.com/tkeel-io/kit/log"
	"go.uber.org/zap"
)

type NodeConf struct {
	Sources []string
}

type Node struct {
	runtimes        map[string]*Runtime
	dispatch        dispatch.Dispatcher
	resourceManager types.ResourceManager
	mappers         map[string]mapper.Mapper

	lock   sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc
}

func NewNode(ctx context.Context, resourceManager types.ResourceManager, dispatcher dispatch.Dispatcher) *Node {
	ctx, cacel := context.WithCancel(ctx)
	return &Node{
		ctx:             ctx,
		cancel:          cacel,
		lock:            sync.RWMutex{},
		dispatch:        dispatcher,
		resourceManager: resourceManager,
		runtimes:        make(map[string]*Runtime),
		mappers:         make(map[string]mapper.Mapper),
	}
}

func (n *Node) Start(cfg NodeConf) error {
	log.Info("start node...")

	n.initializeMetadata()
	for index := range cfg.Sources {
		var err error
		var sourceIns *xkafka.Pubsub
		if sourceIns, err = xkafka.NewKafkaPubsub(cfg.Sources[index]); nil != err {
			return errors.Wrap(err, "create source instance")
		} else if err = sourceIns.Received(n.ctx, n); nil != err {
			return errors.Wrap(err, "consume source")
		}

		placement.Global().Append(placement.Info{ID: sourceIns.ID(), Flag: true})
	}

	return nil
}

func (n *Node) HandleMessage(ctx context.Context, msg *sarama.ConsumerMessage) error {
	rid := msg.Topic
	if _, has := n.runtimes[rid]; !has {
		log.Info("create container", zfield.ID(rid))
		rt := NewRuntime(n.ctx, rid, n.dispatch)
		for _, mp := range n.mapperSlice() {
			if mc, has := n.mapper(mp)[rt.ID()]; has {
				rt.AppendMapper(*mc)
			}
		}
		n.runtimes[rid] = rt
	}

	// load runtime spec.
	runtime := n.runtimes[rid]
	runtime.DeliveredEvent(context.Background(), msg)
	return nil
}

func (n *Node) initializeMetadata() {
	n.listMetadata()
	go n.watchMetadata()
}

// initialize runtime environments.
func (n *Node) listMetadata() {
	elapsedTime := util.NewElapsed()
	ctx, cancel := context.WithTimeout(n.ctx, 30*time.Second)
	defer cancel()

	repo := n.resourceManager.Repo()
	revision := repo.GetLastRevision(context.Background())
	log.Info("initialize actor manager, mapper loadding...")
	repo.RangeMapper(ctx, revision, func(mappers []dao.Mapper) {
		// 将mapper加入每一个 runtime.
		for _, mp := range mappers {
			// parse mapper.
			mpIns, err := mapper.NewMapper(mp, 1)
			if nil != err {
				log.Error("parse mapper", zap.Error(err),
					zfield.Eid(mp.EntityID), zfield.Mid(mp.ID))
				continue
			}
			log.Debug("parse mapper", zfield.Eid(mp.EntityID), zfield.Mid(mp.ID))
			n.mappers[mp.ID] = mpIns
		}
	})

	log.Debug("runtime.Environment initialized", zfield.Elapsedms(elapsedTime.Elapsed()))
}

// watchResource watch resources.
func (n *Node) watchMetadata() {
	repo := n.resourceManager.Repo()
	repo.WatchMapper(context.Background(),
		repo.GetLastRevision(context.Background()),
		func(et dao.EnventType, mp dao.Mapper) {
			// parse mapper.
			var err error
			var mpIns mapper.Mapper
			log.Info("parse mapper", zfield.Eid(mp.EntityID), zfield.Mid(mp.ID))
			if mpIns, err = mapper.NewMapper(mp, 1); nil != err {
				log.Error("parse mapper", zap.Error(err), zfield.Eid(mp.EntityID), zfield.Mid(mp.ID))
				return
			}

			// cache mapper.
			n.mappers[mp.ID] = mpIns
			for rtID, mc := range n.mapper(mpIns) {
				if rt, has := n.runtimes[rtID]; has {
					rt.AppendMapper(*mc)
				}
			}
		})
}

func (n *Node) mapperSlice() []mapper.Mapper {
	mps := []mapper.Mapper{}
	for _, mp := range n.mappers {
		mps = append(mps, mp)
	}
	return mps
}

func (n *Node) mapper(mp mapper.Mapper) map[string]*MCache {
	res := make(map[string]*MCache)
	for eid, tentacles := range mp.Tentacles() {
		// select runtime.
		info := placement.Global().Select(eid)
		if _, exists := res[info.ID]; !exists {
			res[info.ID] = &MCache{
				ID:     mp.ID(),
				Mapper: mp}
		}
		// append tentacles.
		res[info.ID].Tentacles =
			append(res[info.ID].Tentacles, tentacles...)
	}

	return res
}

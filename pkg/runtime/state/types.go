/*
Copyright 2021 The tKeel Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package state

import (
	"context"

	"github.com/tkeel-io/core/pkg/mapper"
	"github.com/tkeel-io/core/pkg/repository/dao"
	"github.com/tkeel-io/core/pkg/runtime/message"
)

type Machiner interface {
	// GetID return state machine id.
	GetID() string
	// GetEntity returns this.Entity.
	GetEntity() *dao.Entity
	// Context return state context.
	Context() *StateContext
	// OnMessage recv message from pubsub.
	Invoke(ctx context.Context, msgCtx message.Context) Result
	// Flush flush entity data.
	Flush(ctx context.Context) error
}

type Result struct {
	Err    error
	Status Status
}

const (
	MCreated   Status = "Created"
	MDeleted   Status = "Deleted"
	MFailured  Status = "Failured"
	MCompleted Status = "Completed"
)

type WatchKey = mapper.WatchKey

type Status string

type MessageHandler = func(context.Context, message.Context) []WatchKey

type Mapper struct {
	ID          string `json:"id" msgpack:"id"`
	TQL         string `json:"tql" msgpack:"tql"`
	Name        string `json:"name" msgpack:"name"`
	Description string `json:"description" msgpack:"description"`
}

type RawData struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Mark      string `json:"mark"`
	Path      string `json:"path"`
	Values    string `json:"values"`
	Timestamp int64  `json:"ts"` //nolint
}

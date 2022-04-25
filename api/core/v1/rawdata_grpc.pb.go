// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RawdataClient is the client API for Rawdata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RawdataClient interface {
	GetRawdata(ctx context.Context, in *GetRawdataRequest, opts ...grpc.CallOption) (*GetRawdataResponse, error)
}

type rawdataClient struct {
	cc grpc.ClientConnInterface
}

func NewRawdataClient(cc grpc.ClientConnInterface) RawdataClient {
	return &rawdataClient{cc}
}

func (c *rawdataClient) GetRawdata(ctx context.Context, in *GetRawdataRequest, opts ...grpc.CallOption) (*GetRawdataResponse, error) {
	out := new(GetRawdataResponse)
	err := c.cc.Invoke(ctx, "/api.core.v1.Rawdata/GetRawdata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RawdataServer is the server API for Rawdata service.
// All implementations must embed UnimplementedRawdataServer
// for forward compatibility
type RawdataServer interface {
	GetRawdata(context.Context, *GetRawdataRequest) (*GetRawdataResponse, error)
	mustEmbedUnimplementedRawdataServer()
}

// UnimplementedRawdataServer must be embedded to have forward compatible implementations.
type UnimplementedRawdataServer struct {
}

func (UnimplementedRawdataServer) GetRawdata(context.Context, *GetRawdataRequest) (*GetRawdataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRawdata not implemented")
}
func (UnimplementedRawdataServer) mustEmbedUnimplementedRawdataServer() {}

// UnsafeRawdataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RawdataServer will
// result in compilation errors.
type UnsafeRawdataServer interface {
	mustEmbedUnimplementedRawdataServer()
}

func RegisterRawdataServer(s grpc.ServiceRegistrar, srv RawdataServer) {
	s.RegisterService(&Rawdata_ServiceDesc, srv)
}

func _Rawdata_GetRawdata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRawdataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawdataServer).GetRawdata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.core.v1.Rawdata/GetRawdata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawdataServer).GetRawdata(ctx, req.(*GetRawdataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rawdata_ServiceDesc is the grpc.ServiceDesc for Rawdata service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rawdata_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.core.v1.Rawdata",
	HandlerType: (*RawdataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRawdata",
			Handler:    _Rawdata_GetRawdata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/core/v1/rawdata.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package worker

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

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerClient interface {
	CommmandMessageLink(ctx context.Context, opts ...grpc.CallOption) (Worker_CommmandMessageLinkClient, error)
	Result(ctx context.Context, in *ResultMessageRequest, opts ...grpc.CallOption) (*ResultMessageResponse, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) CommmandMessageLink(ctx context.Context, opts ...grpc.CallOption) (Worker_CommmandMessageLinkClient, error) {
	stream, err := c.cc.NewStream(ctx, &Worker_ServiceDesc.Streams[0], "/worker.Worker/CommmandMessageLink", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerCommmandMessageLinkClient{stream}
	return x, nil
}

type Worker_CommmandMessageLinkClient interface {
	Send(*ExecutorMessage) error
	Recv() (*WorkerMessage, error)
	grpc.ClientStream
}

type workerCommmandMessageLinkClient struct {
	grpc.ClientStream
}

func (x *workerCommmandMessageLinkClient) Send(m *ExecutorMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerCommmandMessageLinkClient) Recv() (*WorkerMessage, error) {
	m := new(WorkerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerClient) Result(ctx context.Context, in *ResultMessageRequest, opts ...grpc.CallOption) (*ResultMessageResponse, error) {
	out := new(ResultMessageResponse)
	err := c.cc.Invoke(ctx, "/worker.Worker/Result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
// All implementations must embed UnimplementedWorkerServer
// for forward compatibility
type WorkerServer interface {
	CommmandMessageLink(Worker_CommmandMessageLinkServer) error
	Result(context.Context, *ResultMessageRequest) (*ResultMessageResponse, error)
	mustEmbedUnimplementedWorkerServer()
}

// UnimplementedWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (UnimplementedWorkerServer) CommmandMessageLink(Worker_CommmandMessageLinkServer) error {
	return status.Errorf(codes.Unimplemented, "method CommmandMessageLink not implemented")
}
func (UnimplementedWorkerServer) Result(context.Context, *ResultMessageRequest) (*ResultMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedWorkerServer) mustEmbedUnimplementedWorkerServer() {}

// UnsafeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServer will
// result in compilation errors.
type UnsafeWorkerServer interface {
	mustEmbedUnimplementedWorkerServer()
}

func RegisterWorkerServer(s grpc.ServiceRegistrar, srv WorkerServer) {
	s.RegisterService(&Worker_ServiceDesc, srv)
}

func _Worker_CommmandMessageLink_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServer).CommmandMessageLink(&workerCommmandMessageLinkServer{stream})
}

type Worker_CommmandMessageLinkServer interface {
	Send(*WorkerMessage) error
	Recv() (*ExecutorMessage, error)
	grpc.ServerStream
}

type workerCommmandMessageLinkServer struct {
	grpc.ServerStream
}

func (x *workerCommmandMessageLinkServer) Send(m *WorkerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerCommmandMessageLinkServer) Recv() (*ExecutorMessage, error) {
	m := new(ExecutorMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Worker_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.Worker/Result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Result(ctx, req.(*ResultMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Worker_ServiceDesc is the grpc.ServiceDesc for Worker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Worker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "worker.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Result",
			Handler:    _Worker_Result_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CommmandMessageLink",
			Handler:       _Worker_CommmandMessageLink_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "worker.proto",
}
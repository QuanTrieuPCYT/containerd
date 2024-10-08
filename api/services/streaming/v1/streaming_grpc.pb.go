//go:build !no_grpc

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: github.com/containerd/containerd/api/services/streaming/v1/streaming.proto

package streaming

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StreamingClient is the client API for Streaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Streaming_StreamClient, error)
}

type streamingClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingClient(cc grpc.ClientConnInterface) StreamingClient {
	return &streamingClient{cc}
}

func (c *streamingClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Streaming_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Streaming_ServiceDesc.Streams[0], "/containerd.services.streaming.v1.Streaming/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingStreamClient{stream}
	return x, nil
}

type Streaming_StreamClient interface {
	Send(*anypb.Any) error
	Recv() (*anypb.Any, error)
	grpc.ClientStream
}

type streamingStreamClient struct {
	grpc.ClientStream
}

func (x *streamingStreamClient) Send(m *anypb.Any) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingStreamClient) Recv() (*anypb.Any, error) {
	m := new(anypb.Any)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingServer is the server API for Streaming service.
// All implementations must embed UnimplementedStreamingServer
// for forward compatibility
type StreamingServer interface {
	Stream(Streaming_StreamServer) error
	mustEmbedUnimplementedStreamingServer()
}

// UnimplementedStreamingServer must be embedded to have forward compatible implementations.
type UnimplementedStreamingServer struct {
}

func (UnimplementedStreamingServer) Stream(Streaming_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedStreamingServer) mustEmbedUnimplementedStreamingServer() {}

// UnsafeStreamingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingServer will
// result in compilation errors.
type UnsafeStreamingServer interface {
	mustEmbedUnimplementedStreamingServer()
}

func RegisterStreamingServer(s grpc.ServiceRegistrar, srv StreamingServer) {
	s.RegisterService(&Streaming_ServiceDesc, srv)
}

func _Streaming_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingServer).Stream(&streamingStreamServer{stream})
}

type Streaming_StreamServer interface {
	Send(*anypb.Any) error
	Recv() (*anypb.Any, error)
	grpc.ServerStream
}

type streamingStreamServer struct {
	grpc.ServerStream
}

func (x *streamingStreamServer) Send(m *anypb.Any) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingStreamServer) Recv() (*anypb.Any, error) {
	m := new(anypb.Any)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Streaming_ServiceDesc is the grpc.ServiceDesc for Streaming service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Streaming_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.streaming.v1.Streaming",
	HandlerType: (*StreamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Streaming_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/containerd/containerd/api/services/streaming/v1/streaming.proto",
}

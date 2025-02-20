// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: comm/comm.proto

package comm

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CommService_SaySomething_FullMethodName = "/comm.CommService/SaySomething"
)

// CommServiceClient is the client API for CommService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommServiceClient interface {
	SaySomething(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type commServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommServiceClient(cc grpc.ClientConnInterface) CommServiceClient {
	return &commServiceClient{cc}
}

func (c *commServiceClient) SaySomething(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, CommService_SaySomething_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommServiceServer is the server API for CommService service.
// All implementations must embed UnimplementedCommServiceServer
// for forward compatibility.
type CommServiceServer interface {
	SaySomething(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedCommServiceServer()
}

// UnimplementedCommServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommServiceServer struct{}

func (UnimplementedCommServiceServer) SaySomething(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaySomething not implemented")
}
func (UnimplementedCommServiceServer) mustEmbedUnimplementedCommServiceServer() {}
func (UnimplementedCommServiceServer) testEmbeddedByValue()                     {}

// UnsafeCommServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommServiceServer will
// result in compilation errors.
type UnsafeCommServiceServer interface {
	mustEmbedUnimplementedCommServiceServer()
}

func RegisterCommServiceServer(s grpc.ServiceRegistrar, srv CommServiceServer) {
	// If the following call pancis, it indicates UnimplementedCommServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommService_ServiceDesc, srv)
}

func _CommService_SaySomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommServiceServer).SaySomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommService_SaySomething_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommServiceServer).SaySomething(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// CommService_ServiceDesc is the grpc.ServiceDesc for CommService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comm.CommService",
	HandlerType: (*CommServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaySomething",
			Handler:    _CommService_SaySomething_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comm/comm.proto",
}

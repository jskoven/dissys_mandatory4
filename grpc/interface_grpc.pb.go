// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package exclusion

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

// ExclusionClient is the client API for Exclusion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExclusionClient interface {
	GiveToken(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type exclusionClient struct {
	cc grpc.ClientConnInterface
}

func NewExclusionClient(cc grpc.ClientConnInterface) ExclusionClient {
	return &exclusionClient{cc}
}

func (c *exclusionClient) GiveToken(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/exclusion.Exclusion/giveToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExclusionServer is the server API for Exclusion service.
// All implementations must embed UnimplementedExclusionServer
// for forward compatibility
type ExclusionServer interface {
	GiveToken(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedExclusionServer()
}

// UnimplementedExclusionServer must be embedded to have forward compatible implementations.
type UnimplementedExclusionServer struct {
}

func (UnimplementedExclusionServer) GiveToken(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GiveToken not implemented")
}
func (UnimplementedExclusionServer) mustEmbedUnimplementedExclusionServer() {}

// UnsafeExclusionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExclusionServer will
// result in compilation errors.
type UnsafeExclusionServer interface {
	mustEmbedUnimplementedExclusionServer()
}

func RegisterExclusionServer(s grpc.ServiceRegistrar, srv ExclusionServer) {
	s.RegisterService(&Exclusion_ServiceDesc, srv)
}

func _Exclusion_GiveToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExclusionServer).GiveToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exclusion.Exclusion/giveToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExclusionServer).GiveToken(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Exclusion_ServiceDesc is the grpc.ServiceDesc for Exclusion service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exclusion_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exclusion.Exclusion",
	HandlerType: (*ExclusionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "giveToken",
			Handler:    _Exclusion_GiveToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/interface.proto",
}

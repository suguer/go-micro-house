// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: smsService.proto

package service

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

// SmsServiceClient is the client API for SmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsServiceClient interface {
	Index(ctx context.Context, in *SmsIndexRequest, opts ...grpc.CallOption) (*SmsIndexResponse, error)
	Create(ctx context.Context, in *SmsCreateRequest, opts ...grpc.CallOption) (*SmsResponse, error)
	SetConfig(ctx context.Context, in *SmsConfigRequest, opts ...grpc.CallOption) (*SmsConfigResponse, error)
}

type smsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsServiceClient(cc grpc.ClientConnInterface) SmsServiceClient {
	return &smsServiceClient{cc}
}

func (c *smsServiceClient) Index(ctx context.Context, in *SmsIndexRequest, opts ...grpc.CallOption) (*SmsIndexResponse, error) {
	out := new(SmsIndexResponse)
	err := c.cc.Invoke(ctx, "/pb.SmsService/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsServiceClient) Create(ctx context.Context, in *SmsCreateRequest, opts ...grpc.CallOption) (*SmsResponse, error) {
	out := new(SmsResponse)
	err := c.cc.Invoke(ctx, "/pb.SmsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsServiceClient) SetConfig(ctx context.Context, in *SmsConfigRequest, opts ...grpc.CallOption) (*SmsConfigResponse, error) {
	out := new(SmsConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.SmsService/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServiceServer is the server API for SmsService service.
// All implementations must embed UnimplementedSmsServiceServer
// for forward compatibility
type SmsServiceServer interface {
	Index(context.Context, *SmsIndexRequest) (*SmsIndexResponse, error)
	Create(context.Context, *SmsCreateRequest) (*SmsResponse, error)
	SetConfig(context.Context, *SmsConfigRequest) (*SmsConfigResponse, error)
	mustEmbedUnimplementedSmsServiceServer()
}

// UnimplementedSmsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSmsServiceServer struct {
}

func (UnimplementedSmsServiceServer) Index(context.Context, *SmsIndexRequest) (*SmsIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedSmsServiceServer) Create(context.Context, *SmsCreateRequest) (*SmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSmsServiceServer) SetConfig(context.Context, *SmsConfigRequest) (*SmsConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedSmsServiceServer) mustEmbedUnimplementedSmsServiceServer() {}

// UnsafeSmsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServiceServer will
// result in compilation errors.
type UnsafeSmsServiceServer interface {
	mustEmbedUnimplementedSmsServiceServer()
}

func RegisterSmsServiceServer(s grpc.ServiceRegistrar, srv SmsServiceServer) {
	s.RegisterService(&SmsService_ServiceDesc, srv)
}

func _SmsService_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SmsService/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).Index(ctx, req.(*SmsIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SmsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).Create(ctx, req.(*SmsCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsService_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SmsService/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).SetConfig(ctx, req.(*SmsConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmsService_ServiceDesc is the grpc.ServiceDesc for SmsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SmsService",
	HandlerType: (*SmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _SmsService_Index_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SmsService_Create_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _SmsService_SetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smsService.proto",
}

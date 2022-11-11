// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: grpc/centralog.proto

package centralog_agent

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

// CentralogClient is the client API for Centralog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CentralogClient interface {
	Health(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	CheckAPIKey(ctx context.Context, in *CheckAPIKeyRequest, opts ...grpc.CallOption) (*CheckAPIKeyResponse, error)
	FollowLogs(ctx context.Context, in *FollowLogsRequest, opts ...grpc.CallOption) (Centralog_FollowLogsClient, error)
	GetContainers(ctx context.Context, in *GetContainersRequest, opts ...grpc.CallOption) (*ContainerResponse, error)
}

type centralogClient struct {
	cc grpc.ClientConnInterface
}

func NewCentralogClient(cc grpc.ClientConnInterface) CentralogClient {
	return &centralogClient{cc}
}

func (c *centralogClient) Health(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/Centralog/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centralogClient) CheckAPIKey(ctx context.Context, in *CheckAPIKeyRequest, opts ...grpc.CallOption) (*CheckAPIKeyResponse, error) {
	out := new(CheckAPIKeyResponse)
	err := c.cc.Invoke(ctx, "/Centralog/CheckAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centralogClient) FollowLogs(ctx context.Context, in *FollowLogsRequest, opts ...grpc.CallOption) (Centralog_FollowLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Centralog_ServiceDesc.Streams[0], "/Centralog/FollowLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &centralogFollowLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Centralog_FollowLogsClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type centralogFollowLogsClient struct {
	grpc.ClientStream
}

func (x *centralogFollowLogsClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *centralogClient) GetContainers(ctx context.Context, in *GetContainersRequest, opts ...grpc.CallOption) (*ContainerResponse, error) {
	out := new(ContainerResponse)
	err := c.cc.Invoke(ctx, "/Centralog/GetContainers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CentralogServer is the server API for Centralog service.
// All implementations must embed UnimplementedCentralogServer
// for forward compatibility
type CentralogServer interface {
	Health(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	CheckAPIKey(context.Context, *CheckAPIKeyRequest) (*CheckAPIKeyResponse, error)
	FollowLogs(*FollowLogsRequest, Centralog_FollowLogsServer) error
	GetContainers(context.Context, *GetContainersRequest) (*ContainerResponse, error)
	mustEmbedUnimplementedCentralogServer()
}

// UnimplementedCentralogServer must be embedded to have forward compatible implementations.
type UnimplementedCentralogServer struct {
}

func (UnimplementedCentralogServer) Health(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedCentralogServer) CheckAPIKey(context.Context, *CheckAPIKeyRequest) (*CheckAPIKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAPIKey not implemented")
}
func (UnimplementedCentralogServer) FollowLogs(*FollowLogsRequest, Centralog_FollowLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method FollowLogs not implemented")
}
func (UnimplementedCentralogServer) GetContainers(context.Context, *GetContainersRequest) (*ContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainers not implemented")
}
func (UnimplementedCentralogServer) mustEmbedUnimplementedCentralogServer() {}

// UnsafeCentralogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CentralogServer will
// result in compilation errors.
type UnsafeCentralogServer interface {
	mustEmbedUnimplementedCentralogServer()
}

func RegisterCentralogServer(s grpc.ServiceRegistrar, srv CentralogServer) {
	s.RegisterService(&Centralog_ServiceDesc, srv)
}

func _Centralog_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentralogServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Centralog/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentralogServer).Health(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Centralog_CheckAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentralogServer).CheckAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Centralog/CheckAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentralogServer).CheckAPIKey(ctx, req.(*CheckAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Centralog_FollowLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FollowLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CentralogServer).FollowLogs(m, &centralogFollowLogsServer{stream})
}

type Centralog_FollowLogsServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type centralogFollowLogsServer struct {
	grpc.ServerStream
}

func (x *centralogFollowLogsServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

func _Centralog_GetContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentralogServer).GetContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Centralog/GetContainers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentralogServer).GetContainers(ctx, req.(*GetContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Centralog_ServiceDesc is the grpc.ServiceDesc for Centralog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Centralog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Centralog",
	HandlerType: (*CentralogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Centralog_Health_Handler,
		},
		{
			MethodName: "CheckAPIKey",
			Handler:    _Centralog_CheckAPIKey_Handler,
		},
		{
			MethodName: "GetContainers",
			Handler:    _Centralog_GetContainers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FollowLogs",
			Handler:       _Centralog_FollowLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/centralog.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: Proto/message.proto

package __

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

// IntercambioDataNodeClient is the client API for IntercambioDataNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntercambioDataNodeClient interface {
	GuardarData(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Message, error)
	EnviarData(ctx context.Context, in *DataID, opts ...grpc.CallOption) (*Data, error)
	SolicitudCierre(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type intercambioDataNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewIntercambioDataNodeClient(cc grpc.ClientConnInterface) IntercambioDataNodeClient {
	return &intercambioDataNodeClient{cc}
}

func (c *intercambioDataNodeClient) GuardarData(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/grpc.IntercambioDataNode/GuardarData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *intercambioDataNodeClient) EnviarData(ctx context.Context, in *DataID, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/grpc.IntercambioDataNode/EnviarData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *intercambioDataNodeClient) SolicitudCierre(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/grpc.IntercambioDataNode/SolicitudCierre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntercambioDataNodeServer is the server API for IntercambioDataNode service.
// All implementations must embed UnimplementedIntercambioDataNodeServer
// for forward compatibility
type IntercambioDataNodeServer interface {
	GuardarData(context.Context, *Data) (*Message, error)
	EnviarData(context.Context, *DataID) (*Data, error)
	SolicitudCierre(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedIntercambioDataNodeServer()
}

// UnimplementedIntercambioDataNodeServer must be embedded to have forward compatible implementations.
type UnimplementedIntercambioDataNodeServer struct {
}

func (UnimplementedIntercambioDataNodeServer) GuardarData(context.Context, *Data) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuardarData not implemented")
}
func (UnimplementedIntercambioDataNodeServer) EnviarData(context.Context, *DataID) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarData not implemented")
}
func (UnimplementedIntercambioDataNodeServer) SolicitudCierre(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitudCierre not implemented")
}
func (UnimplementedIntercambioDataNodeServer) mustEmbedUnimplementedIntercambioDataNodeServer() {}

// UnsafeIntercambioDataNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntercambioDataNodeServer will
// result in compilation errors.
type UnsafeIntercambioDataNodeServer interface {
	mustEmbedUnimplementedIntercambioDataNodeServer()
}

func RegisterIntercambioDataNodeServer(s grpc.ServiceRegistrar, srv IntercambioDataNodeServer) {
	s.RegisterService(&IntercambioDataNode_ServiceDesc, srv)
}

func _IntercambioDataNode_GuardarData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntercambioDataNodeServer).GuardarData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.IntercambioDataNode/GuardarData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntercambioDataNodeServer).GuardarData(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntercambioDataNode_EnviarData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntercambioDataNodeServer).EnviarData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.IntercambioDataNode/EnviarData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntercambioDataNodeServer).EnviarData(ctx, req.(*DataID))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntercambioDataNode_SolicitudCierre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntercambioDataNodeServer).SolicitudCierre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.IntercambioDataNode/SolicitudCierre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntercambioDataNodeServer).SolicitudCierre(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// IntercambioDataNode_ServiceDesc is the grpc.ServiceDesc for IntercambioDataNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntercambioDataNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.IntercambioDataNode",
	HandlerType: (*IntercambioDataNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GuardarData",
			Handler:    _IntercambioDataNode_GuardarData_Handler,
		},
		{
			MethodName: "EnviarData",
			Handler:    _IntercambioDataNode_EnviarData_Handler,
		},
		{
			MethodName: "SolicitudCierre",
			Handler:    _IntercambioDataNode_SolicitudCierre_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/message.proto",
}

// IntercambioCombineClient is the client API for IntercambioCombine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntercambioCombineClient interface {
	EnvioData(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Message, error)
}

type intercambioCombineClient struct {
	cc grpc.ClientConnInterface
}

func NewIntercambioCombineClient(cc grpc.ClientConnInterface) IntercambioCombineClient {
	return &intercambioCombineClient{cc}
}

func (c *intercambioCombineClient) EnvioData(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/grpc.IntercambioCombine/EnvioData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntercambioCombineServer is the server API for IntercambioCombine service.
// All implementations must embed UnimplementedIntercambioCombineServer
// for forward compatibility
type IntercambioCombineServer interface {
	EnvioData(context.Context, *Data) (*Message, error)
	mustEmbedUnimplementedIntercambioCombineServer()
}

// UnimplementedIntercambioCombineServer must be embedded to have forward compatible implementations.
type UnimplementedIntercambioCombineServer struct {
}

func (UnimplementedIntercambioCombineServer) EnvioData(context.Context, *Data) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnvioData not implemented")
}
func (UnimplementedIntercambioCombineServer) mustEmbedUnimplementedIntercambioCombineServer() {}

// UnsafeIntercambioCombineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntercambioCombineServer will
// result in compilation errors.
type UnsafeIntercambioCombineServer interface {
	mustEmbedUnimplementedIntercambioCombineServer()
}

func RegisterIntercambioCombineServer(s grpc.ServiceRegistrar, srv IntercambioCombineServer) {
	s.RegisterService(&IntercambioCombine_ServiceDesc, srv)
}

func _IntercambioCombine_EnvioData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntercambioCombineServer).EnvioData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.IntercambioCombine/EnvioData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntercambioCombineServer).EnvioData(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

// IntercambioCombine_ServiceDesc is the grpc.ServiceDesc for IntercambioCombine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntercambioCombine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.IntercambioCombine",
	HandlerType: (*IntercambioCombineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnvioData",
			Handler:    _IntercambioCombine_EnvioData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/message.proto",
}

// IntercambioRebeldeClient is the client API for IntercambioRebelde service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntercambioRebeldeClient interface {
	ConsultarData(ctx context.Context, in *ConsultaTipo, opts ...grpc.CallOption) (IntercambioRebelde_ConsultarDataClient, error)
	SolicitudCierre(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type intercambioRebeldeClient struct {
	cc grpc.ClientConnInterface
}

func NewIntercambioRebeldeClient(cc grpc.ClientConnInterface) IntercambioRebeldeClient {
	return &intercambioRebeldeClient{cc}
}

func (c *intercambioRebeldeClient) ConsultarData(ctx context.Context, in *ConsultaTipo, opts ...grpc.CallOption) (IntercambioRebelde_ConsultarDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &IntercambioRebelde_ServiceDesc.Streams[0], "/grpc.IntercambioRebelde/ConsultarData", opts...)
	if err != nil {
		return nil, err
	}
	x := &intercambioRebeldeConsultarDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IntercambioRebelde_ConsultarDataClient interface {
	Recv() (*DataSolicitada, error)
	grpc.ClientStream
}

type intercambioRebeldeConsultarDataClient struct {
	grpc.ClientStream
}

func (x *intercambioRebeldeConsultarDataClient) Recv() (*DataSolicitada, error) {
	m := new(DataSolicitada)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *intercambioRebeldeClient) SolicitudCierre(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/grpc.IntercambioRebelde/SolicitudCierre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntercambioRebeldeServer is the server API for IntercambioRebelde service.
// All implementations must embed UnimplementedIntercambioRebeldeServer
// for forward compatibility
type IntercambioRebeldeServer interface {
	ConsultarData(*ConsultaTipo, IntercambioRebelde_ConsultarDataServer) error
	SolicitudCierre(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedIntercambioRebeldeServer()
}

// UnimplementedIntercambioRebeldeServer must be embedded to have forward compatible implementations.
type UnimplementedIntercambioRebeldeServer struct {
}

func (UnimplementedIntercambioRebeldeServer) ConsultarData(*ConsultaTipo, IntercambioRebelde_ConsultarDataServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsultarData not implemented")
}
func (UnimplementedIntercambioRebeldeServer) SolicitudCierre(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitudCierre not implemented")
}
func (UnimplementedIntercambioRebeldeServer) mustEmbedUnimplementedIntercambioRebeldeServer() {}

// UnsafeIntercambioRebeldeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntercambioRebeldeServer will
// result in compilation errors.
type UnsafeIntercambioRebeldeServer interface {
	mustEmbedUnimplementedIntercambioRebeldeServer()
}

func RegisterIntercambioRebeldeServer(s grpc.ServiceRegistrar, srv IntercambioRebeldeServer) {
	s.RegisterService(&IntercambioRebelde_ServiceDesc, srv)
}

func _IntercambioRebelde_ConsultarData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsultaTipo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IntercambioRebeldeServer).ConsultarData(m, &intercambioRebeldeConsultarDataServer{stream})
}

type IntercambioRebelde_ConsultarDataServer interface {
	Send(*DataSolicitada) error
	grpc.ServerStream
}

type intercambioRebeldeConsultarDataServer struct {
	grpc.ServerStream
}

func (x *intercambioRebeldeConsultarDataServer) Send(m *DataSolicitada) error {
	return x.ServerStream.SendMsg(m)
}

func _IntercambioRebelde_SolicitudCierre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntercambioRebeldeServer).SolicitudCierre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.IntercambioRebelde/SolicitudCierre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntercambioRebeldeServer).SolicitudCierre(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// IntercambioRebelde_ServiceDesc is the grpc.ServiceDesc for IntercambioRebelde service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntercambioRebelde_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.IntercambioRebelde",
	HandlerType: (*IntercambioRebeldeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SolicitudCierre",
			Handler:    _IntercambioRebelde_SolicitudCierre_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConsultarData",
			Handler:       _IntercambioRebelde_ConsultarData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "Proto/message.proto",
}

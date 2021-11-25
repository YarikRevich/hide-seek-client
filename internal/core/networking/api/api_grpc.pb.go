// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExternalServiceClient is the client API for ExternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalServiceClient interface {
	AddWorld(ctx context.Context, in *World, opts ...grpc.CallOption) (*Status, error)
	AddMap(ctx context.Context, in *Map, opts ...grpc.CallOption) (*Status, error)
	AddPC(ctx context.Context, in *PC, opts ...grpc.CallOption) (*Status, error)
	AddElement(ctx context.Context, in *Element, opts ...grpc.CallOption) (*Status, error)
	AddWeapon(ctx context.Context, in *Weapon, opts ...grpc.CallOption) (*Status, error)
	AddAmmo(ctx context.Context, in *Ammo, opts ...grpc.CallOption) (*Status, error)
	UpdateWorld(ctx context.Context, in *World, opts ...grpc.CallOption) (*Status, error)
	UpdatePC(ctx context.Context, in *PC, opts ...grpc.CallOption) (*Status, error)
	UpdateAmmo(ctx context.Context, in *Ammo, opts ...grpc.CallOption) (*Status, error)
	ChooseSpawns(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Status, error)
	DeleteWorld(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Status, error)
	DeletePC(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Status, error)
	GetWorld(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*GetWorldResponse, error)
	GetWorldProperty(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*GetWorldPropertyResponse, error)
	SetGameStarted(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Status, error)
	IsGameStarted(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*IsGameStartedResponse, error)
}

type externalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalServiceClient(cc grpc.ClientConnInterface) ExternalServiceClient {
	return &externalServiceClient{cc}
}

func (c *externalServiceClient) AddWorld(ctx context.Context, in *World, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/AddWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) AddMap(ctx context.Context, in *Map, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/AddMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) AddPC(ctx context.Context, in *PC, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/AddPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) AddElement(ctx context.Context, in *Element, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/AddElement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) AddWeapon(ctx context.Context, in *Weapon, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/AddWeapon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) AddAmmo(ctx context.Context, in *Ammo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/AddAmmo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) UpdateWorld(ctx context.Context, in *World, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/UpdateWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) UpdatePC(ctx context.Context, in *PC, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/UpdatePC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) UpdateAmmo(ctx context.Context, in *Ammo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/UpdateAmmo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) ChooseSpawns(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/ChooseSpawns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) DeleteWorld(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/DeleteWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) DeletePC(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/DeletePC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) GetWorld(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*GetWorldResponse, error) {
	out := new(GetWorldResponse)
	err := c.cc.Invoke(ctx, "/ExternalService/GetWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) GetWorldProperty(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*GetWorldPropertyResponse, error) {
	out := new(GetWorldPropertyResponse)
	err := c.cc.Invoke(ctx, "/ExternalService/GetWorldProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) SetGameStarted(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ExternalService/SetGameStarted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) IsGameStarted(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*IsGameStartedResponse, error) {
	out := new(IsGameStartedResponse)
	err := c.cc.Invoke(ctx, "/ExternalService/IsGameStarted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalServiceServer is the server API for ExternalService service.
// All implementations must embed UnimplementedExternalServiceServer
// for forward compatibility
type ExternalServiceServer interface {
	AddWorld(context.Context, *World) (*Status, error)
	AddMap(context.Context, *Map) (*Status, error)
	AddPC(context.Context, *PC) (*Status, error)
	AddElement(context.Context, *Element) (*Status, error)
	AddWeapon(context.Context, *Weapon) (*Status, error)
	AddAmmo(context.Context, *Ammo) (*Status, error)
	UpdateWorld(context.Context, *World) (*Status, error)
	UpdatePC(context.Context, *PC) (*Status, error)
	UpdateAmmo(context.Context, *Ammo) (*Status, error)
	ChooseSpawns(context.Context, *wrapperspb.StringValue) (*Status, error)
	DeleteWorld(context.Context, *wrapperspb.StringValue) (*Status, error)
	DeletePC(context.Context, *wrapperspb.StringValue) (*Status, error)
	GetWorld(context.Context, *wrapperspb.StringValue) (*GetWorldResponse, error)
	GetWorldProperty(context.Context, *wrapperspb.StringValue) (*GetWorldPropertyResponse, error)
	SetGameStarted(context.Context, *wrapperspb.StringValue) (*Status, error)
	IsGameStarted(context.Context, *wrapperspb.StringValue) (*IsGameStartedResponse, error)
	mustEmbedUnimplementedExternalServiceServer()
}

// UnimplementedExternalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExternalServiceServer struct {
}

func (UnimplementedExternalServiceServer) AddWorld(context.Context, *World) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWorld not implemented")
}
func (UnimplementedExternalServiceServer) AddMap(context.Context, *Map) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMap not implemented")
}
func (UnimplementedExternalServiceServer) AddPC(context.Context, *PC) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPC not implemented")
}
func (UnimplementedExternalServiceServer) AddElement(context.Context, *Element) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddElement not implemented")
}
func (UnimplementedExternalServiceServer) AddWeapon(context.Context, *Weapon) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWeapon not implemented")
}
func (UnimplementedExternalServiceServer) AddAmmo(context.Context, *Ammo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAmmo not implemented")
}
func (UnimplementedExternalServiceServer) UpdateWorld(context.Context, *World) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorld not implemented")
}
func (UnimplementedExternalServiceServer) UpdatePC(context.Context, *PC) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePC not implemented")
}
func (UnimplementedExternalServiceServer) UpdateAmmo(context.Context, *Ammo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAmmo not implemented")
}
func (UnimplementedExternalServiceServer) ChooseSpawns(context.Context, *wrapperspb.StringValue) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChooseSpawns not implemented")
}
func (UnimplementedExternalServiceServer) DeleteWorld(context.Context, *wrapperspb.StringValue) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorld not implemented")
}
func (UnimplementedExternalServiceServer) DeletePC(context.Context, *wrapperspb.StringValue) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePC not implemented")
}
func (UnimplementedExternalServiceServer) GetWorld(context.Context, *wrapperspb.StringValue) (*GetWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorld not implemented")
}
func (UnimplementedExternalServiceServer) GetWorldProperty(context.Context, *wrapperspb.StringValue) (*GetWorldPropertyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorldProperty not implemented")
}
func (UnimplementedExternalServiceServer) SetGameStarted(context.Context, *wrapperspb.StringValue) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGameStarted not implemented")
}
func (UnimplementedExternalServiceServer) IsGameStarted(context.Context, *wrapperspb.StringValue) (*IsGameStartedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsGameStarted not implemented")
}
func (UnimplementedExternalServiceServer) mustEmbedUnimplementedExternalServiceServer() {}

// UnsafeExternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalServiceServer will
// result in compilation errors.
type UnsafeExternalServiceServer interface {
	mustEmbedUnimplementedExternalServiceServer()
}

func RegisterExternalServiceServer(s grpc.ServiceRegistrar, srv ExternalServiceServer) {
	s.RegisterService(&ExternalService_ServiceDesc, srv)
}

func _ExternalService_AddWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(World)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).AddWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/AddWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).AddWorld(ctx, req.(*World))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_AddMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Map)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).AddMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/AddMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).AddMap(ctx, req.(*Map))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_AddPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).AddPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/AddPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).AddPC(ctx, req.(*PC))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_AddElement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Element)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).AddElement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/AddElement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).AddElement(ctx, req.(*Element))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_AddWeapon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Weapon)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).AddWeapon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/AddWeapon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).AddWeapon(ctx, req.(*Weapon))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_AddAmmo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ammo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).AddAmmo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/AddAmmo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).AddAmmo(ctx, req.(*Ammo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_UpdateWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(World)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).UpdateWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/UpdateWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).UpdateWorld(ctx, req.(*World))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_UpdatePC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).UpdatePC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/UpdatePC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).UpdatePC(ctx, req.(*PC))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_UpdateAmmo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ammo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).UpdateAmmo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/UpdateAmmo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).UpdateAmmo(ctx, req.(*Ammo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_ChooseSpawns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).ChooseSpawns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/ChooseSpawns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).ChooseSpawns(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_DeleteWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).DeleteWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/DeleteWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).DeleteWorld(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_DeletePC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).DeletePC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/DeletePC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).DeletePC(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_GetWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).GetWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/GetWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).GetWorld(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_GetWorldProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).GetWorldProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/GetWorldProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).GetWorldProperty(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_SetGameStarted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).SetGameStarted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/SetGameStarted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).SetGameStarted(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_IsGameStarted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).IsGameStarted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternalService/IsGameStarted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).IsGameStarted(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

// ExternalService_ServiceDesc is the grpc.ServiceDesc for ExternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ExternalService",
	HandlerType: (*ExternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddWorld",
			Handler:    _ExternalService_AddWorld_Handler,
		},
		{
			MethodName: "AddMap",
			Handler:    _ExternalService_AddMap_Handler,
		},
		{
			MethodName: "AddPC",
			Handler:    _ExternalService_AddPC_Handler,
		},
		{
			MethodName: "AddElement",
			Handler:    _ExternalService_AddElement_Handler,
		},
		{
			MethodName: "AddWeapon",
			Handler:    _ExternalService_AddWeapon_Handler,
		},
		{
			MethodName: "AddAmmo",
			Handler:    _ExternalService_AddAmmo_Handler,
		},
		{
			MethodName: "UpdateWorld",
			Handler:    _ExternalService_UpdateWorld_Handler,
		},
		{
			MethodName: "UpdatePC",
			Handler:    _ExternalService_UpdatePC_Handler,
		},
		{
			MethodName: "UpdateAmmo",
			Handler:    _ExternalService_UpdateAmmo_Handler,
		},
		{
			MethodName: "ChooseSpawns",
			Handler:    _ExternalService_ChooseSpawns_Handler,
		},
		{
			MethodName: "DeleteWorld",
			Handler:    _ExternalService_DeleteWorld_Handler,
		},
		{
			MethodName: "DeletePC",
			Handler:    _ExternalService_DeletePC_Handler,
		},
		{
			MethodName: "GetWorld",
			Handler:    _ExternalService_GetWorld_Handler,
		},
		{
			MethodName: "GetWorldProperty",
			Handler:    _ExternalService_GetWorldProperty_Handler,
		},
		{
			MethodName: "SetGameStarted",
			Handler:    _ExternalService_SetGameStarted_Handler,
		},
		{
			MethodName: "IsGameStarted",
			Handler:    _ExternalService_IsGameStarted_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

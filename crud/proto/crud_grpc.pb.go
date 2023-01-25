// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: crud.proto

package proto

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

// CrudClient is the client API for Crud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrudClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserInfo, error)
	GetUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserInfo, error)
	GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Users, error)
	UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Status, error)
	DeleteUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error)
}

type crudClient struct {
	cc grpc.ClientConnInterface
}

func NewCrudClient(cc grpc.ClientConnInterface) CrudClient {
	return &crudClient{cc}
}

func (c *crudClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/crud.Crud/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) GetUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/crud.Crud/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/crud.Crud/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/crud.Crud/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) DeleteUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/crud.Crud/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrudServer is the server API for Crud service.
// All implementations must embed UnimplementedCrudServer
// for forward compatibility
type CrudServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*UserInfo, error)
	GetUser(context.Context, *Id) (*UserInfo, error)
	GetUsers(context.Context, *Empty) (*Users, error)
	UpdateUser(context.Context, *UserInfo) (*Status, error)
	DeleteUser(context.Context, *Id) (*Status, error)
	mustEmbedUnimplementedCrudServer()
}

// UnimplementedCrudServer must be embedded to have forward compatible implementations.
type UnimplementedCrudServer struct {
}

func (UnimplementedCrudServer) CreateUser(context.Context, *CreateUserRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedCrudServer) GetUser(context.Context, *Id) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedCrudServer) GetUsers(context.Context, *Empty) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedCrudServer) UpdateUser(context.Context, *UserInfo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedCrudServer) DeleteUser(context.Context, *Id) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedCrudServer) mustEmbedUnimplementedCrudServer() {}

// UnsafeCrudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrudServer will
// result in compilation errors.
type UnsafeCrudServer interface {
	mustEmbedUnimplementedCrudServer()
}

func RegisterCrudServer(s grpc.ServiceRegistrar, srv CrudServer) {
	s.RegisterService(&Crud_ServiceDesc, srv)
}

func _Crud_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud.Crud/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud.Crud/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).GetUser(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud.Crud/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).GetUsers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud.Crud/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).UpdateUser(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud.Crud/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).DeleteUser(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// Crud_ServiceDesc is the grpc.ServiceDesc for Crud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Crud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crud.Crud",
	HandlerType: (*CrudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Crud_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Crud_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Crud_GetUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Crud_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Crud_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crud.proto",
}

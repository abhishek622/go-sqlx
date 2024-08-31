// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: api.proto

package pb

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
	Ecomm_CreateProduct_FullMethodName           = "/pb.ecomm/CreateProduct"
	Ecomm_GetProduct_FullMethodName              = "/pb.ecomm/GetProduct"
	Ecomm_ListProducts_FullMethodName            = "/pb.ecomm/ListProducts"
	Ecomm_UpdateProduct_FullMethodName           = "/pb.ecomm/UpdateProduct"
	Ecomm_DeleteProduct_FullMethodName           = "/pb.ecomm/DeleteProduct"
	Ecomm_CreateOrder_FullMethodName             = "/pb.ecomm/CreateOrder"
	Ecomm_GetOrder_FullMethodName                = "/pb.ecomm/GetOrder"
	Ecomm_ListOrders_FullMethodName              = "/pb.ecomm/ListOrders"
	Ecomm_UpdateOrderStatus_FullMethodName       = "/pb.ecomm/UpdateOrderStatus"
	Ecomm_DeleteOrder_FullMethodName             = "/pb.ecomm/DeleteOrder"
	Ecomm_CreateUser_FullMethodName              = "/pb.ecomm/CreateUser"
	Ecomm_GetUser_FullMethodName                 = "/pb.ecomm/GetUser"
	Ecomm_ListUsers_FullMethodName               = "/pb.ecomm/ListUsers"
	Ecomm_UpdateUser_FullMethodName              = "/pb.ecomm/UpdateUser"
	Ecomm_DeleteUser_FullMethodName              = "/pb.ecomm/DeleteUser"
	Ecomm_CreateSession_FullMethodName           = "/pb.ecomm/CreateSession"
	Ecomm_GetSession_FullMethodName              = "/pb.ecomm/GetSession"
	Ecomm_RevokeSession_FullMethodName           = "/pb.ecomm/RevokeSession"
	Ecomm_DeleteSession_FullMethodName           = "/pb.ecomm/DeleteSession"
	Ecomm_ListNotificationEvents_FullMethodName  = "/pb.ecomm/ListNotificationEvents"
	Ecomm_UpdateNotificationEvent_FullMethodName = "/pb.ecomm/UpdateNotificationEvent"
)

// EcommClient is the client API for Ecomm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EcommClient interface {
	CreateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	GetProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	ListProducts(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ListProductRes, error)
	UpdateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	DeleteProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	CreateOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	GetOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	ListOrders(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*ListOrderRes, error)
	UpdateOrderStatus(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	DeleteOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	CreateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	ListUsers(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ListUserRes, error)
	UpdateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	DeleteUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	CreateSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
	GetSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
	RevokeSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
	DeleteSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
	ListNotificationEvents(ctx context.Context, in *ListNotificationEventsReq, opts ...grpc.CallOption) (*ListNotificationEventsRes, error)
	UpdateNotificationEvent(ctx context.Context, in *UpdateNotificationEventReq, opts ...grpc.CallOption) (*UpdateNotificationEventRes, error)
}

type ecommClient struct {
	cc grpc.ClientConnInterface
}

func NewEcommClient(cc grpc.ClientConnInterface) EcommClient {
	return &ecommClient{cc}
}

func (c *ecommClient) CreateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, Ecomm_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) GetProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, Ecomm_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) ListProducts(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ListProductRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductRes)
	err := c.cc.Invoke(ctx, Ecomm_ListProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) UpdateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, Ecomm_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) DeleteProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, Ecomm_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) CreateOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, Ecomm_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) GetOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, Ecomm_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) ListOrders(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*ListOrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOrderRes)
	err := c.cc.Invoke(ctx, Ecomm_ListOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) UpdateOrderStatus(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, Ecomm_UpdateOrderStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) DeleteOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, Ecomm_DeleteOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) CreateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRes)
	err := c.cc.Invoke(ctx, Ecomm_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRes)
	err := c.cc.Invoke(ctx, Ecomm_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) ListUsers(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ListUserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserRes)
	err := c.cc.Invoke(ctx, Ecomm_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) UpdateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRes)
	err := c.cc.Invoke(ctx, Ecomm_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) DeleteUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRes)
	err := c.cc.Invoke(ctx, Ecomm_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) CreateSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, Ecomm_CreateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) GetSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, Ecomm_GetSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) RevokeSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, Ecomm_RevokeSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) DeleteSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, Ecomm_DeleteSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) ListNotificationEvents(ctx context.Context, in *ListNotificationEventsReq, opts ...grpc.CallOption) (*ListNotificationEventsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNotificationEventsRes)
	err := c.cc.Invoke(ctx, Ecomm_ListNotificationEvents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) UpdateNotificationEvent(ctx context.Context, in *UpdateNotificationEventReq, opts ...grpc.CallOption) (*UpdateNotificationEventRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNotificationEventRes)
	err := c.cc.Invoke(ctx, Ecomm_UpdateNotificationEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EcommServer is the server API for Ecomm service.
// All implementations must embed UnimplementedEcommServer
// for forward compatibility.
type EcommServer interface {
	CreateProduct(context.Context, *ProductReq) (*ProductRes, error)
	GetProduct(context.Context, *ProductReq) (*ProductRes, error)
	ListProducts(context.Context, *ProductReq) (*ListProductRes, error)
	UpdateProduct(context.Context, *ProductReq) (*ProductRes, error)
	DeleteProduct(context.Context, *ProductReq) (*ProductRes, error)
	CreateOrder(context.Context, *OrderReq) (*OrderRes, error)
	GetOrder(context.Context, *OrderReq) (*OrderRes, error)
	ListOrders(context.Context, *OrderReq) (*ListOrderRes, error)
	UpdateOrderStatus(context.Context, *OrderReq) (*OrderRes, error)
	DeleteOrder(context.Context, *OrderReq) (*OrderRes, error)
	CreateUser(context.Context, *UserReq) (*UserRes, error)
	GetUser(context.Context, *UserReq) (*UserRes, error)
	ListUsers(context.Context, *UserReq) (*ListUserRes, error)
	UpdateUser(context.Context, *UserReq) (*UserRes, error)
	DeleteUser(context.Context, *UserReq) (*UserRes, error)
	CreateSession(context.Context, *SessionReq) (*SessionRes, error)
	GetSession(context.Context, *SessionReq) (*SessionRes, error)
	RevokeSession(context.Context, *SessionReq) (*SessionRes, error)
	DeleteSession(context.Context, *SessionReq) (*SessionRes, error)
	ListNotificationEvents(context.Context, *ListNotificationEventsReq) (*ListNotificationEventsRes, error)
	UpdateNotificationEvent(context.Context, *UpdateNotificationEventReq) (*UpdateNotificationEventRes, error)
	mustEmbedUnimplementedEcommServer()
}

// UnimplementedEcommServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEcommServer struct{}

func (UnimplementedEcommServer) CreateProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedEcommServer) GetProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedEcommServer) ListProducts(context.Context, *ProductReq) (*ListProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedEcommServer) UpdateProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedEcommServer) DeleteProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedEcommServer) CreateOrder(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedEcommServer) GetOrder(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedEcommServer) ListOrders(context.Context, *OrderReq) (*ListOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedEcommServer) UpdateOrderStatus(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedEcommServer) DeleteOrder(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedEcommServer) CreateUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedEcommServer) GetUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedEcommServer) ListUsers(context.Context, *UserReq) (*ListUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedEcommServer) UpdateUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedEcommServer) DeleteUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedEcommServer) CreateSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedEcommServer) GetSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedEcommServer) RevokeSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeSession not implemented")
}
func (UnimplementedEcommServer) DeleteSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedEcommServer) ListNotificationEvents(context.Context, *ListNotificationEventsReq) (*ListNotificationEventsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotificationEvents not implemented")
}
func (UnimplementedEcommServer) UpdateNotificationEvent(context.Context, *UpdateNotificationEventReq) (*UpdateNotificationEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotificationEvent not implemented")
}
func (UnimplementedEcommServer) mustEmbedUnimplementedEcommServer() {}
func (UnimplementedEcommServer) testEmbeddedByValue()               {}

// UnsafeEcommServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EcommServer will
// result in compilation errors.
type UnsafeEcommServer interface {
	mustEmbedUnimplementedEcommServer()
}

func RegisterEcommServer(s grpc.ServiceRegistrar, srv EcommServer) {
	// If the following call pancis, it indicates UnimplementedEcommServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Ecomm_ServiceDesc, srv)
}

func _Ecomm_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).CreateProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).GetProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).ListProducts(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).UpdateProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).DeleteProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).CreateOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).GetOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_ListOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).ListOrders(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_UpdateOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).UpdateOrderStatus(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).DeleteOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).CreateUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).GetUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).ListUsers(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).UpdateUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).DeleteUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).CreateSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_GetSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).GetSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_RevokeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).RevokeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_RevokeSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).RevokeSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_DeleteSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).DeleteSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_ListNotificationEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotificationEventsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).ListNotificationEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_ListNotificationEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).ListNotificationEvents(ctx, req.(*ListNotificationEventsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_UpdateNotificationEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotificationEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).UpdateNotificationEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecomm_UpdateNotificationEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).UpdateNotificationEvent(ctx, req.(*UpdateNotificationEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Ecomm_ServiceDesc is the grpc.ServiceDesc for Ecomm service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ecomm_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ecomm",
	HandlerType: (*EcommServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _Ecomm_CreateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Ecomm_GetProduct_Handler,
		},
		{
			MethodName: "ListProducts",
			Handler:    _Ecomm_ListProducts_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Ecomm_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Ecomm_DeleteProduct_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Ecomm_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Ecomm_GetOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _Ecomm_ListOrders_Handler,
		},
		{
			MethodName: "UpdateOrderStatus",
			Handler:    _Ecomm_UpdateOrderStatus_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _Ecomm_DeleteOrder_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Ecomm_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Ecomm_GetUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Ecomm_ListUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Ecomm_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Ecomm_DeleteUser_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _Ecomm_CreateSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _Ecomm_GetSession_Handler,
		},
		{
			MethodName: "RevokeSession",
			Handler:    _Ecomm_RevokeSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _Ecomm_DeleteSession_Handler,
		},
		{
			MethodName: "ListNotificationEvents",
			Handler:    _Ecomm_ListNotificationEvents_Handler,
		},
		{
			MethodName: "UpdateNotificationEvent",
			Handler:    _Ecomm_UpdateNotificationEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

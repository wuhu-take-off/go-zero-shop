// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: shop.proto

package shop

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

const (
	UserServe_UserLogin_FullMethodName = "/shop.UserServe/UserLogin"
)

// UserServeClient is the client API for UserServe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServeClient interface {
	UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
}

type userServeClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServeClient(cc grpc.ClientConnInterface) UserServeClient {
	return &userServeClient{cc}
}

func (c *userServeClient) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	out := new(UserLoginResp)
	err := c.cc.Invoke(ctx, UserServe_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServeServer is the server API for UserServe service.
// All implementations must embed UnimplementedUserServeServer
// for forward compatibility
type UserServeServer interface {
	UserLogin(context.Context, *UserLoginReq) (*UserLoginResp, error)
	mustEmbedUnimplementedUserServeServer()
}

// UnimplementedUserServeServer must be embedded to have forward compatible implementations.
type UnimplementedUserServeServer struct {
}

func (UnimplementedUserServeServer) UserLogin(context.Context, *UserLoginReq) (*UserLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServeServer) mustEmbedUnimplementedUserServeServer() {}

// UnsafeUserServeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServeServer will
// result in compilation errors.
type UnsafeUserServeServer interface {
	mustEmbedUnimplementedUserServeServer()
}

func RegisterUserServeServer(s grpc.ServiceRegistrar, srv UserServeServer) {
	s.RegisterService(&UserServe_ServiceDesc, srv)
}

func _UserServe_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServeServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServe_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServeServer).UserLogin(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServe_ServiceDesc is the grpc.ServiceDesc for UserServe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.UserServe",
	HandlerType: (*UserServeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserServe_UserLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

const (
	GoodsServe_GoodsList_FullMethodName   = "/shop.GoodsServe/GoodsList"
	GoodsServe_GoodsAdd_FullMethodName    = "/shop.GoodsServe/GoodsAdd"
	GoodsServe_GoodsUpdate_FullMethodName = "/shop.GoodsServe/GoodsUpdate"
	GoodsServe_GoodsDel_FullMethodName    = "/shop.GoodsServe/GoodsDel"
)

// GoodsServeClient is the client API for GoodsServe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsServeClient interface {
	// 获取商品列表
	GoodsList(ctx context.Context, in *GoodsListReq, opts ...grpc.CallOption) (*GoodsListResp, error)
	// 添加商品
	GoodsAdd(ctx context.Context, in *GoodsAddReq, opts ...grpc.CallOption) (*GoodsAddResp, error)
	// 更新商品
	GoodsUpdate(ctx context.Context, in *GoodsUpdateReq, opts ...grpc.CallOption) (*GoodsUpdateResp, error)
	// 删除商品
	GoodsDel(ctx context.Context, in *GoodsDelReq, opts ...grpc.CallOption) (*GoodsDelResp, error)
}

type goodsServeClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsServeClient(cc grpc.ClientConnInterface) GoodsServeClient {
	return &goodsServeClient{cc}
}

func (c *goodsServeClient) GoodsList(ctx context.Context, in *GoodsListReq, opts ...grpc.CallOption) (*GoodsListResp, error) {
	out := new(GoodsListResp)
	err := c.cc.Invoke(ctx, GoodsServe_GoodsList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServeClient) GoodsAdd(ctx context.Context, in *GoodsAddReq, opts ...grpc.CallOption) (*GoodsAddResp, error) {
	out := new(GoodsAddResp)
	err := c.cc.Invoke(ctx, GoodsServe_GoodsAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServeClient) GoodsUpdate(ctx context.Context, in *GoodsUpdateReq, opts ...grpc.CallOption) (*GoodsUpdateResp, error) {
	out := new(GoodsUpdateResp)
	err := c.cc.Invoke(ctx, GoodsServe_GoodsUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServeClient) GoodsDel(ctx context.Context, in *GoodsDelReq, opts ...grpc.CallOption) (*GoodsDelResp, error) {
	out := new(GoodsDelResp)
	err := c.cc.Invoke(ctx, GoodsServe_GoodsDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServeServer is the server API for GoodsServe service.
// All implementations must embed UnimplementedGoodsServeServer
// for forward compatibility
type GoodsServeServer interface {
	// 获取商品列表
	GoodsList(context.Context, *GoodsListReq) (*GoodsListResp, error)
	// 添加商品
	GoodsAdd(context.Context, *GoodsAddReq) (*GoodsAddResp, error)
	// 更新商品
	GoodsUpdate(context.Context, *GoodsUpdateReq) (*GoodsUpdateResp, error)
	// 删除商品
	GoodsDel(context.Context, *GoodsDelReq) (*GoodsDelResp, error)
	mustEmbedUnimplementedGoodsServeServer()
}

// UnimplementedGoodsServeServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServeServer struct {
}

func (UnimplementedGoodsServeServer) GoodsList(context.Context, *GoodsListReq) (*GoodsListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoodsList not implemented")
}
func (UnimplementedGoodsServeServer) GoodsAdd(context.Context, *GoodsAddReq) (*GoodsAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoodsAdd not implemented")
}
func (UnimplementedGoodsServeServer) GoodsUpdate(context.Context, *GoodsUpdateReq) (*GoodsUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoodsUpdate not implemented")
}
func (UnimplementedGoodsServeServer) GoodsDel(context.Context, *GoodsDelReq) (*GoodsDelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoodsDel not implemented")
}
func (UnimplementedGoodsServeServer) mustEmbedUnimplementedGoodsServeServer() {}

// UnsafeGoodsServeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServeServer will
// result in compilation errors.
type UnsafeGoodsServeServer interface {
	mustEmbedUnimplementedGoodsServeServer()
}

func RegisterGoodsServeServer(s grpc.ServiceRegistrar, srv GoodsServeServer) {
	s.RegisterService(&GoodsServe_ServiceDesc, srv)
}

func _GoodsServe_GoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServeServer).GoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoodsServe_GoodsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServeServer).GoodsList(ctx, req.(*GoodsListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsServe_GoodsAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServeServer).GoodsAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoodsServe_GoodsAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServeServer).GoodsAdd(ctx, req.(*GoodsAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsServe_GoodsUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServeServer).GoodsUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoodsServe_GoodsUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServeServer).GoodsUpdate(ctx, req.(*GoodsUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsServe_GoodsDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServeServer).GoodsDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoodsServe_GoodsDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServeServer).GoodsDel(ctx, req.(*GoodsDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GoodsServe_ServiceDesc is the grpc.ServiceDesc for GoodsServe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodsServe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.GoodsServe",
	HandlerType: (*GoodsServeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GoodsList",
			Handler:    _GoodsServe_GoodsList_Handler,
		},
		{
			MethodName: "GoodsAdd",
			Handler:    _GoodsServe_GoodsAdd_Handler,
		},
		{
			MethodName: "GoodsUpdate",
			Handler:    _GoodsServe_GoodsUpdate_Handler,
		},
		{
			MethodName: "GoodsDel",
			Handler:    _GoodsServe_GoodsDel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

const (
	FunctionServe_SendEmailAuth_FullMethodName = "/shop.FunctionServe/SendEmailAuth"
	FunctionServe_UploadImg_FullMethodName     = "/shop.FunctionServe/UploadImg"
	FunctionServe_LoadImg_FullMethodName       = "/shop.FunctionServe/LoadImg"
)

// FunctionServeClient is the client API for FunctionServe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FunctionServeClient interface {
	// 发送邮箱验证码
	SendEmailAuth(ctx context.Context, in *SendEmailAuthReq, opts ...grpc.CallOption) (*SendEmailAuthResp, error)
	// 上传如图片
	UploadImg(ctx context.Context, in *UploadImgReq, opts ...grpc.CallOption) (*UploadImgResp, error)
	// 加载图片
	LoadImg(ctx context.Context, in *LoadImgReq, opts ...grpc.CallOption) (*LoadImgResp, error)
}

type functionServeClient struct {
	cc grpc.ClientConnInterface
}

func NewFunctionServeClient(cc grpc.ClientConnInterface) FunctionServeClient {
	return &functionServeClient{cc}
}

func (c *functionServeClient) SendEmailAuth(ctx context.Context, in *SendEmailAuthReq, opts ...grpc.CallOption) (*SendEmailAuthResp, error) {
	out := new(SendEmailAuthResp)
	err := c.cc.Invoke(ctx, FunctionServe_SendEmailAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionServeClient) UploadImg(ctx context.Context, in *UploadImgReq, opts ...grpc.CallOption) (*UploadImgResp, error) {
	out := new(UploadImgResp)
	err := c.cc.Invoke(ctx, FunctionServe_UploadImg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionServeClient) LoadImg(ctx context.Context, in *LoadImgReq, opts ...grpc.CallOption) (*LoadImgResp, error) {
	out := new(LoadImgResp)
	err := c.cc.Invoke(ctx, FunctionServe_LoadImg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FunctionServeServer is the server API for FunctionServe service.
// All implementations must embed UnimplementedFunctionServeServer
// for forward compatibility
type FunctionServeServer interface {
	// 发送邮箱验证码
	SendEmailAuth(context.Context, *SendEmailAuthReq) (*SendEmailAuthResp, error)
	// 上传如图片
	UploadImg(context.Context, *UploadImgReq) (*UploadImgResp, error)
	// 加载图片
	LoadImg(context.Context, *LoadImgReq) (*LoadImgResp, error)
	mustEmbedUnimplementedFunctionServeServer()
}

// UnimplementedFunctionServeServer must be embedded to have forward compatible implementations.
type UnimplementedFunctionServeServer struct {
}

func (UnimplementedFunctionServeServer) SendEmailAuth(context.Context, *SendEmailAuthReq) (*SendEmailAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailAuth not implemented")
}
func (UnimplementedFunctionServeServer) UploadImg(context.Context, *UploadImgReq) (*UploadImgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImg not implemented")
}
func (UnimplementedFunctionServeServer) LoadImg(context.Context, *LoadImgReq) (*LoadImgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadImg not implemented")
}
func (UnimplementedFunctionServeServer) mustEmbedUnimplementedFunctionServeServer() {}

// UnsafeFunctionServeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FunctionServeServer will
// result in compilation errors.
type UnsafeFunctionServeServer interface {
	mustEmbedUnimplementedFunctionServeServer()
}

func RegisterFunctionServeServer(s grpc.ServiceRegistrar, srv FunctionServeServer) {
	s.RegisterService(&FunctionServe_ServiceDesc, srv)
}

func _FunctionServe_SendEmailAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionServeServer).SendEmailAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FunctionServe_SendEmailAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionServeServer).SendEmailAuth(ctx, req.(*SendEmailAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionServe_UploadImg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadImgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionServeServer).UploadImg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FunctionServe_UploadImg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionServeServer).UploadImg(ctx, req.(*UploadImgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionServe_LoadImg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadImgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionServeServer).LoadImg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FunctionServe_LoadImg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionServeServer).LoadImg(ctx, req.(*LoadImgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FunctionServe_ServiceDesc is the grpc.ServiceDesc for FunctionServe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FunctionServe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.FunctionServe",
	HandlerType: (*FunctionServeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmailAuth",
			Handler:    _FunctionServe_SendEmailAuth_Handler,
		},
		{
			MethodName: "UploadImg",
			Handler:    _FunctionServe_UploadImg_Handler,
		},
		{
			MethodName: "LoadImg",
			Handler:    _FunctionServe_LoadImg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

const (
	TaskServe_TaskList_FullMethodName   = "/shop.TaskServe/TaskList"
	TaskServe_UpdateTask_FullMethodName = "/shop.TaskServe/UpdateTask"
)

// TaskServeClient is the client API for TaskServe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServeClient interface {
	// 获取任务列表
	TaskList(ctx context.Context, in *TaskListReq, opts ...grpc.CallOption) (*TaskListResp, error)
	// 更新任务信息
	UpdateTask(ctx context.Context, in *UpdateTaskReq, opts ...grpc.CallOption) (*UpdateTaskResp, error)
}

type taskServeClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServeClient(cc grpc.ClientConnInterface) TaskServeClient {
	return &taskServeClient{cc}
}

func (c *taskServeClient) TaskList(ctx context.Context, in *TaskListReq, opts ...grpc.CallOption) (*TaskListResp, error) {
	out := new(TaskListResp)
	err := c.cc.Invoke(ctx, TaskServe_TaskList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServeClient) UpdateTask(ctx context.Context, in *UpdateTaskReq, opts ...grpc.CallOption) (*UpdateTaskResp, error) {
	out := new(UpdateTaskResp)
	err := c.cc.Invoke(ctx, TaskServe_UpdateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServeServer is the server API for TaskServe service.
// All implementations must embed UnimplementedTaskServeServer
// for forward compatibility
type TaskServeServer interface {
	// 获取任务列表
	TaskList(context.Context, *TaskListReq) (*TaskListResp, error)
	// 更新任务信息
	UpdateTask(context.Context, *UpdateTaskReq) (*UpdateTaskResp, error)
	mustEmbedUnimplementedTaskServeServer()
}

// UnimplementedTaskServeServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServeServer struct {
}

func (UnimplementedTaskServeServer) TaskList(context.Context, *TaskListReq) (*TaskListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskList not implemented")
}
func (UnimplementedTaskServeServer) UpdateTask(context.Context, *UpdateTaskReq) (*UpdateTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTaskServeServer) mustEmbedUnimplementedTaskServeServer() {}

// UnsafeTaskServeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServeServer will
// result in compilation errors.
type UnsafeTaskServeServer interface {
	mustEmbedUnimplementedTaskServeServer()
}

func RegisterTaskServeServer(s grpc.ServiceRegistrar, srv TaskServeServer) {
	s.RegisterService(&TaskServe_ServiceDesc, srv)
}

func _TaskServe_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServeServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskServe_TaskList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServeServer).TaskList(ctx, req.(*TaskListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskServe_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServeServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskServe_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServeServer).UpdateTask(ctx, req.(*UpdateTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskServe_ServiceDesc is the grpc.ServiceDesc for TaskServe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskServe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.TaskServe",
	HandlerType: (*TaskServeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskList",
			Handler:    _TaskServe_TaskList_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskServe_UpdateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

const (
	OrderServe_OrderList_FullMethodName           = "/shop.OrderServe/OrderList"
	OrderServe_UpdateCourierNumber_FullMethodName = "/shop.OrderServe/UpdateCourierNumber"
)

// OrderServeClient is the client API for OrderServe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServeClient interface {
	// 获取任务列表
	OrderList(ctx context.Context, in *OrderListReq, opts ...grpc.CallOption) (*OrderListResp, error)
	// 更新任务信息
	UpdateCourierNumber(ctx context.Context, in *UpdateCourierNumberReq, opts ...grpc.CallOption) (*UpdateCourierNumberResp, error)
}

type orderServeClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServeClient(cc grpc.ClientConnInterface) OrderServeClient {
	return &orderServeClient{cc}
}

func (c *orderServeClient) OrderList(ctx context.Context, in *OrderListReq, opts ...grpc.CallOption) (*OrderListResp, error) {
	out := new(OrderListResp)
	err := c.cc.Invoke(ctx, OrderServe_OrderList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServeClient) UpdateCourierNumber(ctx context.Context, in *UpdateCourierNumberReq, opts ...grpc.CallOption) (*UpdateCourierNumberResp, error) {
	out := new(UpdateCourierNumberResp)
	err := c.cc.Invoke(ctx, OrderServe_UpdateCourierNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServeServer is the server API for OrderServe service.
// All implementations must embed UnimplementedOrderServeServer
// for forward compatibility
type OrderServeServer interface {
	// 获取任务列表
	OrderList(context.Context, *OrderListReq) (*OrderListResp, error)
	// 更新任务信息
	UpdateCourierNumber(context.Context, *UpdateCourierNumberReq) (*UpdateCourierNumberResp, error)
	mustEmbedUnimplementedOrderServeServer()
}

// UnimplementedOrderServeServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServeServer struct {
}

func (UnimplementedOrderServeServer) OrderList(context.Context, *OrderListReq) (*OrderListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (UnimplementedOrderServeServer) UpdateCourierNumber(context.Context, *UpdateCourierNumberReq) (*UpdateCourierNumberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourierNumber not implemented")
}
func (UnimplementedOrderServeServer) mustEmbedUnimplementedOrderServeServer() {}

// UnsafeOrderServeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServeServer will
// result in compilation errors.
type UnsafeOrderServeServer interface {
	mustEmbedUnimplementedOrderServeServer()
}

func RegisterOrderServeServer(s grpc.ServiceRegistrar, srv OrderServeServer) {
	s.RegisterService(&OrderServe_ServiceDesc, srv)
}

func _OrderServe_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServeServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServe_OrderList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServeServer).OrderList(ctx, req.(*OrderListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderServe_UpdateCourierNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourierNumberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServeServer).UpdateCourierNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServe_UpdateCourierNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServeServer).UpdateCourierNumber(ctx, req.(*UpdateCourierNumberReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderServe_ServiceDesc is the grpc.ServiceDesc for OrderServe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderServe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.OrderServe",
	HandlerType: (*OrderServeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderList",
			Handler:    _OrderServe_OrderList_Handler,
		},
		{
			MethodName: "UpdateCourierNumber",
			Handler:    _OrderServe_UpdateCourierNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

const (
	MerchantServe_MerchantList_FullMethodName   = "/shop.MerchantServe/MerchantList"
	MerchantServe_UpdateMerchant_FullMethodName = "/shop.MerchantServe/UpdateMerchant"
	MerchantServe_DelMerchant_FullMethodName    = "/shop.MerchantServe/DelMerchant"
)

// MerchantServeClient is the client API for MerchantServe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MerchantServeClient interface {
	MerchantList(ctx context.Context, in *MerchantListReq, opts ...grpc.CallOption) (*MerchantListResp, error)
	UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, opts ...grpc.CallOption) (*UpdateMerchantResp, error)
	DelMerchant(ctx context.Context, in *DelMerchantReq, opts ...grpc.CallOption) (*DelMerchantResp, error)
}

type merchantServeClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantServeClient(cc grpc.ClientConnInterface) MerchantServeClient {
	return &merchantServeClient{cc}
}

func (c *merchantServeClient) MerchantList(ctx context.Context, in *MerchantListReq, opts ...grpc.CallOption) (*MerchantListResp, error) {
	out := new(MerchantListResp)
	err := c.cc.Invoke(ctx, MerchantServe_MerchantList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServeClient) UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, opts ...grpc.CallOption) (*UpdateMerchantResp, error) {
	out := new(UpdateMerchantResp)
	err := c.cc.Invoke(ctx, MerchantServe_UpdateMerchant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServeClient) DelMerchant(ctx context.Context, in *DelMerchantReq, opts ...grpc.CallOption) (*DelMerchantResp, error) {
	out := new(DelMerchantResp)
	err := c.cc.Invoke(ctx, MerchantServe_DelMerchant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantServeServer is the server API for MerchantServe service.
// All implementations must embed UnimplementedMerchantServeServer
// for forward compatibility
type MerchantServeServer interface {
	MerchantList(context.Context, *MerchantListReq) (*MerchantListResp, error)
	UpdateMerchant(context.Context, *UpdateMerchantReq) (*UpdateMerchantResp, error)
	DelMerchant(context.Context, *DelMerchantReq) (*DelMerchantResp, error)
	mustEmbedUnimplementedMerchantServeServer()
}

// UnimplementedMerchantServeServer must be embedded to have forward compatible implementations.
type UnimplementedMerchantServeServer struct {
}

func (UnimplementedMerchantServeServer) MerchantList(context.Context, *MerchantListReq) (*MerchantListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MerchantList not implemented")
}
func (UnimplementedMerchantServeServer) UpdateMerchant(context.Context, *UpdateMerchantReq) (*UpdateMerchantResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMerchant not implemented")
}
func (UnimplementedMerchantServeServer) DelMerchant(context.Context, *DelMerchantReq) (*DelMerchantResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMerchant not implemented")
}
func (UnimplementedMerchantServeServer) mustEmbedUnimplementedMerchantServeServer() {}

// UnsafeMerchantServeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MerchantServeServer will
// result in compilation errors.
type UnsafeMerchantServeServer interface {
	mustEmbedUnimplementedMerchantServeServer()
}

func RegisterMerchantServeServer(s grpc.ServiceRegistrar, srv MerchantServeServer) {
	s.RegisterService(&MerchantServe_ServiceDesc, srv)
}

func _MerchantServe_MerchantList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServeServer).MerchantList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantServe_MerchantList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServeServer).MerchantList(ctx, req.(*MerchantListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantServe_UpdateMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMerchantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServeServer).UpdateMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantServe_UpdateMerchant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServeServer).UpdateMerchant(ctx, req.(*UpdateMerchantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantServe_DelMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelMerchantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServeServer).DelMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantServe_DelMerchant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServeServer).DelMerchant(ctx, req.(*DelMerchantReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MerchantServe_ServiceDesc is the grpc.ServiceDesc for MerchantServe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MerchantServe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.MerchantServe",
	HandlerType: (*MerchantServeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MerchantList",
			Handler:    _MerchantServe_MerchantList_Handler,
		},
		{
			MethodName: "UpdateMerchant",
			Handler:    _MerchantServe_UpdateMerchant_Handler,
		},
		{
			MethodName: "DelMerchant",
			Handler:    _MerchantServe_DelMerchant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

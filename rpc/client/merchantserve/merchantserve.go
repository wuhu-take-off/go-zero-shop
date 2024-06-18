// Code generated by goctl. DO NOT EDIT.
// Source: shop.proto

package merchantserve

import (
	"context"

	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddMerchantReq          = shop.AddMerchantReq
	AddMerchantResp         = shop.AddMerchantResp
	DelMerchantReq          = shop.DelMerchantReq
	DelMerchantResp         = shop.DelMerchantResp
	GoodsAddReq             = shop.GoodsAddReq
	GoodsAddResp            = shop.GoodsAddResp
	GoodsDelReq             = shop.GoodsDelReq
	GoodsDelResp            = shop.GoodsDelResp
	GoodsList               = shop.GoodsList
	GoodsListReq            = shop.GoodsListReq
	GoodsListResp           = shop.GoodsListResp
	GoodsTypeList           = shop.GoodsTypeList
	GoodsTypeListReq        = shop.GoodsTypeListReq
	GoodsTypeListResp       = shop.GoodsTypeListResp
	GoodsUpdateReq          = shop.GoodsUpdateReq
	GoodsUpdateResp         = shop.GoodsUpdateResp
	Img                     = shop.Img
	LoadImgReq              = shop.LoadImgReq
	LoadImgResp             = shop.LoadImgResp
	MerchantList            = shop.MerchantList
	MerchantListReq         = shop.MerchantListReq
	MerchantListResp        = shop.MerchantListResp
	OrderList               = shop.OrderList
	OrderListReq            = shop.OrderListReq
	OrderListResp           = shop.OrderListResp
	SendEmailAuthReq        = shop.SendEmailAuthReq
	SendEmailAuthResp       = shop.SendEmailAuthResp
	SpecificationReq        = shop.SpecificationReq
	SpecificationResp       = shop.SpecificationResp
	TaskList                = shop.TaskList
	TaskListReq             = shop.TaskListReq
	TaskListResp            = shop.TaskListResp
	TaskTypeList            = shop.TaskTypeList
	TaskTypeListReq         = shop.TaskTypeListReq
	TaskTypeListResp        = shop.TaskTypeListResp
	UpdateCourierNumberReq  = shop.UpdateCourierNumberReq
	UpdateCourierNumberResp = shop.UpdateCourierNumberResp
	UpdateMerchantReq       = shop.UpdateMerchantReq
	UpdateMerchantResp      = shop.UpdateMerchantResp
	UpdateTaskReq           = shop.UpdateTaskReq
	UpdateTaskResp          = shop.UpdateTaskResp
	UploadImgReq            = shop.UploadImgReq
	UploadImgResp           = shop.UploadImgResp
	UserLoginReq            = shop.UserLoginReq
	UserLoginResp           = shop.UserLoginResp

	MerchantServe interface {
		MerchantList(ctx context.Context, in *MerchantListReq, opts ...grpc.CallOption) (*MerchantListResp, error)
		UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, opts ...grpc.CallOption) (*UpdateMerchantResp, error)
		DelMerchant(ctx context.Context, in *DelMerchantReq, opts ...grpc.CallOption) (*DelMerchantResp, error)
		AddMerchant(ctx context.Context, in *AddMerchantReq, opts ...grpc.CallOption) (*AddMerchantResp, error)
	}

	defaultMerchantServe struct {
		cli zrpc.Client
	}
)

func NewMerchantServe(cli zrpc.Client) MerchantServe {
	return &defaultMerchantServe{
		cli: cli,
	}
}

func (m *defaultMerchantServe) MerchantList(ctx context.Context, in *MerchantListReq, opts ...grpc.CallOption) (*MerchantListResp, error) {
	client := shop.NewMerchantServeClient(m.cli.Conn())
	return client.MerchantList(ctx, in, opts...)
}

func (m *defaultMerchantServe) UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, opts ...grpc.CallOption) (*UpdateMerchantResp, error) {
	client := shop.NewMerchantServeClient(m.cli.Conn())
	return client.UpdateMerchant(ctx, in, opts...)
}

func (m *defaultMerchantServe) DelMerchant(ctx context.Context, in *DelMerchantReq, opts ...grpc.CallOption) (*DelMerchantResp, error) {
	client := shop.NewMerchantServeClient(m.cli.Conn())
	return client.DelMerchant(ctx, in, opts...)
}

func (m *defaultMerchantServe) AddMerchant(ctx context.Context, in *AddMerchantReq, opts ...grpc.CallOption) (*AddMerchantResp, error) {
	client := shop.NewMerchantServeClient(m.cli.Conn())
	return client.AddMerchant(ctx, in, opts...)
}

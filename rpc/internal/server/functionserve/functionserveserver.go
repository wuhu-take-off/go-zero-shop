// Code generated by goctl. DO NOT EDIT.
// Source: shop.proto

package server

import (
	"context"

	"TongChi_shop/rpc/internal/logic/functionserve"
	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"
)

type FunctionServeServer struct {
	svcCtx *svc.ServiceContext
	shop.UnimplementedFunctionServeServer
}

func NewFunctionServeServer(svcCtx *svc.ServiceContext) *FunctionServeServer {
	return &FunctionServeServer{
		svcCtx: svcCtx,
	}
}

func (s *FunctionServeServer) SendEmailAuth(ctx context.Context, in *shop.SendEmailAuthReq) (*shop.SendEmailAuthResp, error) {
	l := functionservelogic.NewSendEmailAuthLogic(ctx, s.svcCtx)
	return l.SendEmailAuth(in)
}
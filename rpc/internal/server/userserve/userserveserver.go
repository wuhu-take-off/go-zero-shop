// Code generated by goctl. DO NOT EDIT.
// Source: shop.proto

package server

import (
	"context"

	"TongChi_shop/rpc/internal/logic/userserve"
	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"
)

type UserServeServer struct {
	svcCtx *svc.ServiceContext
	shop.UnimplementedUserServeServer
}

func NewUserServeServer(svcCtx *svc.ServiceContext) *UserServeServer {
	return &UserServeServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServeServer) UserLogin(ctx context.Context, in *shop.UserLoginReq) (*shop.UserLoginResp, error) {
	l := userservelogic.NewUserLoginLogic(ctx, s.svcCtx)
	return l.UserLogin(in)
}

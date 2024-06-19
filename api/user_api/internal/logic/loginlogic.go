package logic

import (
	"TongChi_shop/rpc/shop"
	"TongChi_shop/tool/jwt_authentication"
	"context"
	"errors"
	"time"

	"TongChi_shop/api/user_api/internal/svc"
	"TongChi_shop/api/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	login, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &shop.UserLoginReq{
		Email:     req.Email,
		EmailAuth: req.Auth,
	})
	if err != nil || login.UserId == 0 {
		return nil, err
	}

	//签发token
	jwt, err := jwt_authentication.GenerateJWT(login.UserId, []byte(l.svcCtx.Config.Auth.AccessSecret), time.Duration(l.svcCtx.Config.Auth.AccessExpire)*time.Second)
	if err != nil {
		return nil, errors.New("签发失败")
	}
	return &types.UserLoginResp{
		Token:          jwt,
		ExpirationTime: l.svcCtx.Config.Auth.AccessExpire,
		RefreshTime:    l.svcCtx.Config.Auth.AccessExpire / 2,
	}, nil
}

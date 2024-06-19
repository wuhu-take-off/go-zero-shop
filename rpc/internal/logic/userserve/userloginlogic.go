package userservelogic

import (
	"TongChi_shop/model/user_model"
	"TongChi_shop/tool/db_init"
	"context"
	"errors"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *shop.UserLoginReq) (*shop.UserLoginResp, error) {
	redisConn := db_init.NewRedisConn(l.svcCtx.Config.CacheRedis.Addr, l.svcCtx.Config.CacheRedis.Password)
	if redisConn.CheckEmailAuth(context.Background(), in.Email, in.EmailAuth) {
		return &shop.UserLoginResp{
			UserId: user_model.NewUserInfosModel(l.svcCtx.DB).SelectUserIdWhereEmail(in.Email),
		}, nil
	}
	return nil, errors.New("账号或密码错误")
}

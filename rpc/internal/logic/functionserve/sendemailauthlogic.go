package functionservelogic

import (
	"TongChi_shop/model/user_model"
	"TongChi_shop/tool/db_init"
	"context"
	"errors"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailAuthLogic {
	return &SendEmailAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailAuthLogic) SendEmailAuth(in *shop.SendEmailAuthReq) (*shop.SendEmailAuthResp, error) {
	// todo: add your logic here and delete this line

	//判断邮箱是否存在
	userModel := user_model.NewUserInfosModel(l.svcCtx.DB)
	if !userModel.EmailExist(in.Email) {
		return nil, errors.New("邮箱不存在")
	}

	//生成验证码
	redisConn := db_init.NewRedisConn(l.svcCtx.Config.CacheRedis.Addr, l.svcCtx.Config.CacheRedis.Password)
	err := redisConn.AddEmailAuth(context.Background(), in.Email)
	if err != nil {
		return nil, err
	}

	return &shop.SendEmailAuthResp{
		OK: 1,
	}, nil
}

package logic

import (
	"TongChi_shop/tool/jwt_authentication"
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"TongChi_shop/api/function_api/internal/svc"
	"TongChi_shop/api/function_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewTokenLogic {
	return &NewTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewTokenLogic) NewToken(req *types.NewTokenReq) (resp *types.NewTokenResp, err error) {
	value := l.ctx.Value("UserID")
	var userId int32
	switch v := value.(type) {
	case json.Number:
		if intValue, err := strconv.ParseInt(string(v), 10, 32); err == nil {
			userId = int32(intValue)
		} else {
			//fmt.Println("无法将 json.Number 转换为 int32:", err)
			return nil, errors.New("非法访问")
		}
	default:
		//fmt.Println("anyValue 不是 json.Number 类型")
		return nil, errors.New("非法访问")
	}
	////验证token
	//token, ok := l.ctx.Value("token").(string)
	//if !ok {
	//	return nil, errors.New("非法访问")
	//}
	//res, err := jwt_authentication.ParseJWT(token, []byte(l.svcCtx.Config.Auth.AccessSecret))

	if err != nil {
		return nil, errors.New("非法访问")
	}
	token, err := jwt_authentication.GenerateJWT(int32(userId), []byte(l.svcCtx.Config.Auth.AccessSecret), time.Duration(l.svcCtx.Config.Auth.AccessExpire)*time.Second)
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &types.NewTokenResp{
		//新token
		Token: token,
		//过期时间
		ExpirationTime: l.svcCtx.Config.Auth.AccessExpire,
		//期望刷新时间
		RefreshTime: l.svcCtx.Config.Auth.AccessExpire / 2,
	}, nil
}

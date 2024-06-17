package functionservelogic

import (
	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImgLogic {
	return &UploadImgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadImgLogic) UploadImg(in *shop.UploadImgReq) (*shop.UploadImgResp, error) {
	return &shop.UploadImgResp{
		Url: "url",
	}, nil
}

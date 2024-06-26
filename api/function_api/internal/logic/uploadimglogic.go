package logic

import (
	"TongChi_shop/api/function_api/internal/svc"
	"TongChi_shop/api/function_api/internal/types"
	"TongChi_shop/rpc/shop"
	"context"
	"errors"
	"io"
	"mime/multipart"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImgLogic {
	return &UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//func (l *UploadImgLogic) UploadImg(req *types.UploadImgReq) (resp *types.UploadImgResp, err error) {
//
//	fmt.Println(req.GoodsId)
//	return
//}

func (l *UploadImgLogic) UploadImg(fileDate multipart.File, filename string, req *types.UploadImgReq) (resp *types.UploadImgResp, err error) {
	date, err := io.ReadAll(fileDate)
	if err != nil {
		return nil, errors.New("文件提取失败")
	}
	res, err := l.svcCtx.FunctionRpc.UploadImg(l.ctx, &shop.UploadImgReq{
		Img:     date,
		ImgName: filename,
		GoodsId: req.GoodsId,
	})
	if err != nil {
		return nil, err
	}
	return &types.UploadImgResp{
		URL: res.Url,
	}, nil
}

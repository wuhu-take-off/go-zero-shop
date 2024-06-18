package functionservelogic

import (
	imgs_modelr "TongChi_shop/model/imgs_model"
	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"
	"TongChi_shop/tool/file_load"
	"context"
	"errors"
	"strconv"
	"strings"

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
	load, err := file_load.NewFileLoad(l.svcCtx.Config.SFTP)
	if err != nil {
		return nil, errors.New("上传失败")
	}
	defer load.Close()
	imgModel := imgs_modelr.NewImgsModel(l.svcCtx.DB)
	imgId, err := imgModel.Add(in.GoodsId, in.ImgName)
	if err != nil {
		return nil, errors.New("添加记录失败")
	}
	split := strings.Split(in.ImgName, ".")
	endName := "." + split[len(split)-1]

	err = load.UpFile(in.Img, strconv.Itoa(int(imgId))+endName)
	if err != nil {
		return nil, errors.New("上传失败")
	}
	return &shop.UploadImgResp{
		Url: "url",
	}, nil
}

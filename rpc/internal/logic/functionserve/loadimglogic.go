package functionservelogic

import (
	imgs_modelr "TongChi_shop/model/imgs_model"
	"TongChi_shop/tool/file_load"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoadImgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoadImgLogic {
	return &LoadImgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoadImgLogic) LoadImg(in *shop.LoadImgReq) (*shop.LoadImgResp, error) {
	imgsModel := imgs_modelr.NewImgsModel(l.svcCtx.DB)
	load, err := file_load.NewFileLoad(l.svcCtx.Config.SFTP)
	if err != nil {
		return nil, errors.New("服务连接失败")
	}
	imgs := imgsModel.FindImgs(in.GoodsId)
	var tmp []*shop.Img
	for i := 0; i < len(imgs); i++ {
		split := strings.Split(imgs[i].ImgName, ".")
		fileName := strconv.Itoa(int(imgs[i].ImgId)) + "." + split[len(split)-1]
		file, err := load.DownloadFile(fileName)
		if err != nil {
			fmt.Println("文件获取失败:", fileName)
			continue
		}
		tmp = append(tmp, &shop.Img{Img: file})
	}
	return &shop.LoadImgResp{
		Imgs: tmp,
	}, nil
}

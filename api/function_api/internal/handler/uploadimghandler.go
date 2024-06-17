package handler

import (
	"TongChi_shop/api/function_api/internal/logic"
	"TongChi_shop/api/function_api/internal/svc"
	"TongChi_shop/api/function_api/internal/types"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"mime/multipart"
	"net/http"
	"strconv"
)

func UploadImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadImgReq
		file, handler, err2 := r.FormFile("file")
		if err2 != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("获取文件失败"))
		}
		atoi, err2 := strconv.Atoi(r.FormValue("goodsId"))
		if err2 != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("获取信息失败"))
		}
		req.GoodsId = int32(atoi)

		// 关闭文件数据
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		// 获取文件名
		filename := handler.Filename
		l := logic.NewUploadImgLogic(r.Context(), svcCtx)

		resp, err := l.UploadImg(file, filename, &req)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

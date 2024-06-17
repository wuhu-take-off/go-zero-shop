package handler

import (
	"net/http"

	"TongChi_shop/api/function_api/internal/logic"
	"TongChi_shop/api/function_api/internal/svc"
	"TongChi_shop/api/function_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadImgReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadImgLogic(r.Context(), svcCtx)
		resp, err := l.UploadImg(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

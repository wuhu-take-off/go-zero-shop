package handler

import (
	"net/http"

	"TongChi_shop/api/task_api/internal/logic"
	"TongChi_shop/api/task_api/internal/svc"
	"TongChi_shop/api/task_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TaskListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTaskListLogic(r.Context(), svcCtx)
		resp, err := l.TaskList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

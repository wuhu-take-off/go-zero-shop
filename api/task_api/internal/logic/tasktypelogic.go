package logic

import (
	"context"

	"TongChi_shop/api/task_api/internal/svc"
	"TongChi_shop/api/task_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskTypeLogic {
	return &TaskTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskTypeLogic) TaskType(req *types.TaskTypeListReq) (resp *types.TaskTypeListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

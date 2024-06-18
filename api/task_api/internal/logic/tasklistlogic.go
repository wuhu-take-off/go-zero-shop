package logic

import (
	"context"

	"TongChi_shop/api/task_api/internal/svc"
	"TongChi_shop/api/task_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskListLogic {
	return &TaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskListLogic) TaskList(req *types.TaskListReq) (resp *types.TaskListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

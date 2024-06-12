package taskservelogic

import (
	"context"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskListLogic {
	return &TaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取任务列表
func (l *TaskListLogic) TaskList(in *shop.TaskListReq) (*shop.TaskListResp, error) {
	// todo: add your logic here and delete this line

	return &shop.TaskListResp{}, nil
}

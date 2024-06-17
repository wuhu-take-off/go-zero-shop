package logic

import (
	"context"

	"TongChi_shop/api/task_api/internal/svc"
	"TongChi_shop/api/task_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskUpdateLogic {
	return &TaskUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskUpdateLogic) TaskUpdate(req *types.UpDateTaskReq) (resp *types.UpDateTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}

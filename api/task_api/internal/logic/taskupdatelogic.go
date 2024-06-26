package logic

import (
	"TongChi_shop/rpc/shop"
	json_number_transition "TongChi_shop/tool/json.number_transition"
	"context"
	"errors"

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
	value := l.ctx.Value("UserID")
	userId, err := json_number_transition.Transition(value)
	if err != nil {
		return nil, errors.New("非法访问")
	}

	task, err := l.svcCtx.TaskRpc.UpdateTask(l.ctx, &shop.UpdateTaskReq{
		UserId: int32(userId),
		TaskId: req.TaskId,
		Score:  req.Score,
	})
	if err != nil || !task.Ok {
		return nil, err
	}

	return
}

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

	value := l.ctx.Value("UserID")
	userId, err := json_number_transition.Transition(value)
	if err != nil {
		return nil, errors.New("非法访问")
	}
	taskList, err := l.svcCtx.TaskRpc.TaskList(l.ctx, &shop.TaskListReq{
		UserId:     int32(userId),
		Status:     req.Status,
		TaskTypeId: req.TaskTypeId,
		TaskName:   req.TaskName,
		TaskId:     req.TaskId,
		Page:       req.Page,
		Limit:      req.Limit,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.TaskList
	for _, t := range taskList.TaskList {
		list = append(list, &types.TaskList{
			TaskId:      t.TaskId,
			TaskName:    t.TaskName,
			TaskTypeId:  t.TaskTypeId,
			Status:      t.Status,
			Description: t.Description,
			Score:       t.Score,
		})
	}
	return &types.TaskListResp{
		Count:  taskList.Count,
		Limit:  taskList.Limit,
		Number: taskList.Number,
		Page:   taskList.Page,
		List:   list,
	}, nil
}

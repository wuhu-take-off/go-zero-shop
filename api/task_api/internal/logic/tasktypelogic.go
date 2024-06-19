package logic

import (
	"TongChi_shop/rpc/shop"
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

	tasktypeList, err := l.svcCtx.TaskRpc.TaskTypeList(l.ctx, &shop.TaskTypeListReq{})
	if err != nil {
		return nil, err
	}
	var list []*types.TaskTypeList
	for _, t := range tasktypeList.TaskTypeList {
		list = append(list, &types.TaskTypeList{
			TaskTypeId:   t.TaskTypeId,
			TaskTypeName: t.TaskTypeName,
		})
	}
	return &types.TaskTypeListResp{
		Count: tasktypeList.Count,
		List:  list,
	}, nil
}

package taskservelogic

import (
	"TongChi_shop/model/task_type_model"
	"context"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskTypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskTypeListLogic {
	return &TaskTypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取任务类型列表
func (l *TaskTypeListLogic) TaskTypeList(in *shop.TaskTypeListReq) (*shop.TaskTypeListResp, error) {

	goodsTypes := task_type_model.NewTaskTypeModel(l.svcCtx.DB).FindTaskTypes()
	var list []*shop.TaskTypeList
	for i := 0; i < len(goodsTypes); i++ {
		list = append(list, goodsTypes[i].ToTaskTypesResp())
	}
	return &shop.TaskTypeListResp{
		TaskTypeList: list,
		Count:        int32(len(goodsTypes)),
	}, nil
}

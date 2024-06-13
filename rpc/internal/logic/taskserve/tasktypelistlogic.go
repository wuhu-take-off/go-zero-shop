package taskservelogic

import (
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
	// todo: add your logic here and delete this line

	return &shop.TaskTypeListResp{}, nil
}

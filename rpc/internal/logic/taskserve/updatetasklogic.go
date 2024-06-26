package taskservelogic

import (
	"TongChi_shop/model/task_model"
	"TongChi_shop/model/user_model"
	"context"
	"errors"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskLogic {
	return &UpdateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新任务信息
func (l *UpdateTaskLogic) UpdateTask(in *shop.UpdateTaskReq) (*shop.UpdateTaskResp, error) {

	if !user_model.NewUserInfosModel(l.svcCtx.DB).CheckUsersTask(in.UserId, in.TaskId) {
		return nil, errors.New("非法访问")
	}
	if err := task_model.NewTaskInfosModel(l.svcCtx.DB).UpdateTaskScore(in.TaskId, in.Score); err != nil {
		return nil, errors.New("修改失败")
	}
	return &shop.UpdateTaskResp{
		Ok: true,
	}, nil
}

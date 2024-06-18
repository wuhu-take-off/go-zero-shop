package taskservelogic

import (
	"TongChi_shop/model/user_model"
	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"
	"TongChi_shop/tool/raw_field"
	"context"
	"strings"

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
	sql, val := raw_field.UpdateFieldMap(in, user_model.UserTaskSqlFileMap)
	userInfosModel := user_model.NewUserInfosModel(l.svcCtx.DB)
	count, tasks := userInfosModel.FindUserTask(in.UserId, in.Page, in.Limit, strings.Join(sql, " AND "), val...)
	var taskList []*shop.TaskList
	for i := 0; i < len(tasks); i++ {
		taskList = append(taskList, tasks[i].ToTaskListResp())
	}

	return &shop.TaskListResp{
		TaskList: taskList,
		Count:    count,
		Limit:    in.Limit,
		Number:   (count + in.Limit + 1) / in.Limit,
		Page:     in.Page,
	}, nil
}

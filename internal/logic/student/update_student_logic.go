package student

import (
	"context"

	"school/internal/svc"
	"school/internal/types"
	"school/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentLogic {
	return &UpdateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStudentLogic) UpdateStudent(req *types.StudentInfo) (*types.BaseMsgResp, error) {
	err := l.svcCtx.DB.Student.UpdateOneID(req.Id).
		SetNotEmptySort(req.Sort).
		SetNotEmptyStatus(req.Status).
		SetNotEmptyName(req.Name).
		SetNotEmptyIDCard(req.IdCard).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}

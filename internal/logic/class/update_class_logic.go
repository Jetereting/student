package class

import (
	"context"

	"school/internal/svc"
	"school/internal/types"
	"school/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClassLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClassLogic {
	return &UpdateClassLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateClassLogic) UpdateClass(req *types.ClassInfo) (*types.BaseMsgResp, error) {
	err := l.svcCtx.DB.Class.UpdateOneID(req.Id).
		SetNotEmptySort(req.Sort).
		SetNotEmptyStatus(req.Status).
		SetNotEmptyName(req.Name).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}

package class

import (
	"context"

	"school/internal/svc"
	"school/internal/types"
	"school/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClassByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassByIdLogic {
	return &GetClassByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetClassByIdLogic) GetClassById(req *types.IDReq) (*types.ClassInfoResp, error) {
	data, err := l.svcCtx.DB.Class.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.ClassInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.ClassInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.ID,
				CreatedAt: data.CreatedAt.UnixMilli(),
				UpdatedAt: data.UpdatedAt.UnixMilli(),
			},
			Sort:   data.Sort,
			Status: data.Status,
			Name:   data.Name,
		},
	}, nil
}

package student

import (
	"context"

	"school/internal/svc"
	"school/internal/types"
	"school/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentByIdLogic {
	return &GetStudentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentByIdLogic) GetStudentById(req *types.IDReq) (*types.StudentInfoResp, error) {
	data, err := l.svcCtx.DB.Student.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.StudentInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.StudentInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.ID,
				CreatedAt: data.CreatedAt.UnixMilli(),
				UpdatedAt: data.UpdatedAt.UnixMilli(),
			},
			Sort:   data.Sort,
			Status: data.Status,
			Name:   data.Name,
			IdCard: data.IDCard,
		},
	}, nil
}

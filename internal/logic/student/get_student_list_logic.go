package student

import (
	"context"

	"school/ent/predicate"
	"school/ent/student"
	"school/internal/svc"
	"school/internal/types"
	"school/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentListLogic {
	return &GetStudentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentListLogic) GetStudentList(req *types.StudentListReq) (*types.StudentListResp, error) {
	var predicates []predicate.Student
	data, err := l.svcCtx.DB.Student.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.StudentListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.StudentInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.ID,
					CreatedAt: v.CreatedAt.UnixMilli(),
					UpdatedAt: v.UpdatedAt.UnixMilli(),
				},
				Sort:   v.Sort,
				Status: v.Status,
				Name:   v.Name,
				IdCard: v.IDCard,
			})
	}

	return resp, nil
}

package class

import (
	"context"

	"school/ent/class"
	"school/ent/predicate"
	"school/internal/svc"
	"school/internal/types"
	"school/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClassListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassListLogic {
	return &GetClassListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetClassListLogic) GetClassList(req *types.ClassListReq) (*types.ClassListResp, error) {
	var predicates []predicate.Class
	if req.Name != "" {
		predicates = append(predicates, class.NameContains(req.Name))
	}
	data, err := l.svcCtx.DB.Class.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.ClassListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.ClassInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.ID,
					CreatedAt: v.CreatedAt.UnixMilli(),
					UpdatedAt: v.UpdatedAt.UnixMilli(),
				},
				Sort:   v.Sort,
				Status: v.Status,
				Name:   v.Name,
			})
	}

	return resp, nil
}

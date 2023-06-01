package base

import (
	"context"
	"school/ent/migrate"
	"school/internal/svc"
	"school/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	l.ctx = context.Background()
	l.svcCtx.DB.Schema.Create(l.ctx, migrate.WithForeignKeys(false),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	)
	return
}

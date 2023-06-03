package base

import (
	"context"
	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"school/ent/migrate"
	"school/internal/svc"
	"school/internal/types"
	"strings"
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
	err = l.svcCtx.DB.Schema.Create(l.ctx,
		migrate.WithForeignKeys(false),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	)
	if err != nil {
		logx.Error(err)
		return
	}
	// 解析swagger 到 sys_apis 表
	err = l.addApis()
	if err != nil {
		logx.Error(err)
		return
	}
	// 给管理员添加所有权限
	_, err = l.svcCtx.DB.QueryContext(l.ctx, `
	insert into casbin_rules (ptype, v0, v1, v2)
select 'p', '001', path, method from sys_apis s
left join casbin_rules c on s.path=c.v1 and s.method=c.v2 and c.ptype='p' and c.v0='001'
where c.id is null;`)
	if err != nil {
		logx.Error(err)
		return
	}
	return
}

type api struct {
	Path        string // 路径
	Method      string // 方法
	ApiGroup    string // api组
	Description string // 描述
}

func getApis() (apis []api, err error) {
	swagger, err := os.ReadFile("school.json")
	gjson.GetBytes(swagger, "paths").ForEach(func(path, value gjson.Result) bool {
		value.ForEach(func(method, value gjson.Result) bool {
			apiGroup := ""
			tags := value.Get("tags").Array()
			if len(tags) > 0 {
				apiGroup = tags[0].Str
			}
			description := value.Get("description").String()
			if len(strings.Split(description, " | ")) > 1 {
				description = strings.Split(description, " | ")[1]
			}
			apis = append(apis, api{
				Path:        path.String(),
				Method:      strings.ToUpper(method.String()),
				ApiGroup:    apiGroup,
				Description: description,
			})
			return true
		})
		return true
	})
	return
}
func (l *InitDatabaseLogic) addApis() (err error) {
	apis, err := getApis()
	if err != nil {
		logx.Error(err)
		return
	}
	hasAdd := false
	sqlAdd := `insert ignore into sys_apis (created_at,updated_at,path,method,api_group,description)values`
	for _, v := range apis {
		hasAdd = true
		sqlAdd += `(now(),now(),'` + v.Path + `','` + v.Method + `','` + v.ApiGroup + `','` + v.Description + `'),`
	}
	if hasAdd {
		_, err = l.svcCtx.DB.QueryContext(l.ctx, sqlAdd[:len(sqlAdd)-1])
		if err != nil {
			logx.Error(err)
			return
		}
	}
	return
}

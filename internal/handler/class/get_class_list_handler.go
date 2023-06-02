package class

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"school/internal/logic/class"
	"school/internal/svc"
	"school/internal/types"
)

// swagger:route post /class/list class GetClassList
//
// Get class list | 获取Class列表
//
// Get class list | 获取Class列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ClassListReq
//
// Responses:
//  200: ClassListResp

func GetClassListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClassListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := class.NewGetClassListLogic(r.Context(), svcCtx)
		resp, err := l.GetClassList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

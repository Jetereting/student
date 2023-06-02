package class

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"school/internal/logic/class"
	"school/internal/svc"
	"school/internal/types"
)

// swagger:route post /class/create class CreateClass
//
// Create class information | 创建Class
//
// Create class information | 创建Class
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ClassInfo
//
// Responses:
//  200: BaseMsgResp

func CreateClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClassInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := class.NewCreateClassLogic(r.Context(), svcCtx)
		resp, err := l.CreateClass(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package class

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"school/internal/logic/class"
	"school/internal/svc"
	"school/internal/types"
)

// swagger:route post /class class GetClassById
//
// Get class by ID | 通过ID获取Class
//
// Get class by ID | 通过ID获取Class
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: ClassInfoResp

func GetClassByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := class.NewGetClassByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetClassById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package class

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"school/internal/logic/class"
	"school/internal/svc"
	"school/internal/types"
)

// swagger:route post /class/delete class DeleteClass
//
// Delete class information | 删除Class信息
//
// Delete class information | 删除Class信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := class.NewDeleteClassLogic(r.Context(), svcCtx)
		resp, err := l.DeleteClass(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

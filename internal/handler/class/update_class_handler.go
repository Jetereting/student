package class

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"school/internal/logic/class"
	"school/internal/svc"
	"school/internal/types"
)

// swagger:route post /class/update class UpdateClass
//
// Update class information | 更新Class
//
// Update class information | 更新Class
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ClassInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClassInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := class.NewUpdateClassLogic(r.Context(), svcCtx)
		resp, err := l.UpdateClass(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package student

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"school/internal/logic/student"
	"school/internal/svc"
	"school/internal/types"
)

// swagger:route post /student/delete student DeleteStudent
//
// Delete student information | 删除Student信息
//
// Delete student information | 删除Student信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteStudentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := student.NewDeleteStudentLogic(r.Context(), svcCtx)
		resp, err := l.DeleteStudent(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package interviews

import (
	"net/http"

	"github.com/minibear2333/programmer-go/api/internal/logic/interviews"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllInterviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqInterviews
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := interviews.NewGetAllInterviewLogic(r.Context(), svcCtx)
		resp, err := l.GetAllInterview(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

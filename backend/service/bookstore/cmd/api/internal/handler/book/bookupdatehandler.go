package handler

import (
	"backend/common/errorx"
	logic "backend/service/bookstore/cmd/api/internal/logic/book"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

// 图书管理 update
func BookUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BookPostReq

		fmt.Printf("%v\n", r.Form)

		for name := range r.Form {
			formValue := r.Form.Get(name)
			if len(formValue) > 0 {
				fmt.Printf("%v  %v\n", name, formValue)
			}
		}
		if r.PostForm == nil {
			fmt.Println("r.PostForm==nil")
		}
		if r.Form == nil {
			fmt.Println("r.Form==nil", len(r.PostForm))
			fmt.Println("")
		}
		fmt.Println("title", r.PostFormValue("title"))
		fmt.Println("description", r.PostFormValue("description"))
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}
		fmt.Println("---------------req------------------")
		fmt.Println(req)
		fmt.Println("---------------req------------------")
		req.Title = r.PostFormValue("title")
		req.Description = r.PostFormValue("description")
		id, _ := strconv.Atoi(r.PostFormValue("id"))
		req.Id = int64(id)
		req.Cover = r.PostFormValue("cover")
		sort, _ := strconv.Atoi(r.PostFormValue("sort"))
		req.Sort = int64(sort)
		req.CreateBy = r.PostFormValue("create_by")
		req.UpdateBy = r.PostFormValue("update_by")

		l := logic.NewBookUpdateLogic(r.Context(), ctx)
		_, header, _ := r.FormFile("file")
		resp, err := l.BookUpdate(req, header)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

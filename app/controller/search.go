package controller

import (
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/projectforge/app/controller/cutil"

	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/views/vsearch"
)

func Search(ctx *fasthttp.RequestCtx) {
	act("search", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := string(ctx.URI().QueryArgs().Peek("q"))
		q = strings.TrimSpace(q)
		results := []string{"a", "b", "c"}
		ps.Title = "Search Results"
		ps.Data = results
		return render(ctx, as, &vsearch.Results{Q: q, Results: results}, ps, "Search")
	})
}

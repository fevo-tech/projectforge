// Code generated by Project Forge, see https://projectforge.dev for details.
package controller

import (
	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/controller/cutil"
	"github.com/kyleu/projectforge/app/util"
	"github.com/kyleu/projectforge/views"
	"github.com/valyala/fasthttp"
)

func About(ctx *fasthttp.RequestCtx) {
	act("about", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = util.AppName + " v" + as.BuildInfo.Version
		return render(ctx, as, &views.About{}, ps)
	})
}


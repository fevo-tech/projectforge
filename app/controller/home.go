// Package controller $PF_IGNORE$
package controller

import (
	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/controller/cutil"
	"github.com/kyleu/projectforge/views"
	"github.com/valyala/fasthttp"
)

func Home(rc *fasthttp.RequestCtx) {
	act("home", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prjs := as.Services.Projects.Projects()
		mods := as.Services.Modules.Modules()
		ps.Data = map[string]interface{}{"projects": prjs, "modules": mods}
		return render(rc, as, &views.Home{Projects: prjs, Modules: mods}, ps)
	})
}

package controller

import (
	"github.com/valyala/fasthttp"
	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/action"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/doctor"
	"projectforge.dev/projectforge/app/doctor/checks"
	"projectforge.dev/projectforge/app/lib/menu"
	"projectforge.dev/projectforge/views/vdoctor"
)

func DoctorMenu(i string, r string) *menu.Item {
	return &menu.Item{Key: action.TypeDoctor.Key, Title: action.TypeDoctor.Title, Description: action.TypeDoctor.Description, Icon: i, Route: r}
}

func Doctor(rc *fasthttp.RequestCtx) {
	act("doctor", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prjs := as.Services.Projects.Projects()
		ret := checks.ForModules(prjs.AllModules())
		ps.Data = ret
		return render(rc, as, &vdoctor.List{Checks: ret}, ps, "doctor")
	})
}

func DoctorRunAll(rc *fasthttp.RequestCtx) {
	act("doctor.run.all", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prjs := as.Services.Projects.Projects()
		ret := checks.CheckAll(ps.Context, prjs.AllModules(), as.Logger)
		ps.Data = ret
		return render(rc, as, &vdoctor.Results{Results: ret}, ps, "doctor", "All")
	})
}

func DoctorRun(rc *fasthttp.RequestCtx) {
	act("doctor.run", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		c := checks.GetCheck(key)
		ret := c.Check(ps.Context, as.Logger)
		ps.Data = ret
		return render(rc, as, &vdoctor.Results{Results: doctor.Results{ret}}, ps, "doctor", c.Title)
	})
}

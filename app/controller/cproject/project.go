package cproject

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/views/vproject"
)

func ProjectDetail(rc *fasthttp.RequestCtx) {
	controller.Act("project.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, err := getProject(rc, as)
		if err != nil {
			return "", err
		}
		prj.ExportArgs, _ = prj.ModuleArgExport(as.Services.Projects, ps.Logger)
		mods := as.Services.Modules.Modules()
		gitStatus, _ := as.Services.Git.Status(ps.Context, prj, ps.Logger)
		execs := as.Services.Exec.Execs.GetByKey(prj.Key)
		ps.Title = fmt.Sprintf("%s (project %s)", prj.Title(), prj.Key)
		ps.Data = prj
		return controller.Render(rc, as, &vproject.Detail{Project: prj, Modules: mods, GitResult: gitStatus, Execs: execs}, ps, "projects", prj.Key)
	})
}

func ProjectForm(rc *fasthttp.RequestCtx) {
	controller.Act("project.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj := project.NewProject("", ".")
		ps.Title = "New Project"
		ps.Data = prj
		return controller.Render(rc, as, &vproject.Edit{Project: prj}, ps, "projects", "New")
	})
}

func ProjectCreate(rc *fasthttp.RequestCtx) {
	controller.Act("project.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		prj := project.NewProject("", "")
		err = projectFromForm(frm, prj)
		if err != nil {
			return "", err
		}
		key := frm.GetStringOpt("key")
		prj.Key = key
		err = as.Services.Projects.Save(prj, ps.Logger)
		if err != nil {
			return controller.ERsp("unable to save project: %+v", err)
		}
		return controller.FlashAndRedir(true, "Created project ["+prj.Title()+"]", "/p/"+prj.Key, rc, ps)
	})
}

func ProjectEdit(rc *fasthttp.RequestCtx) {
	controller.Act("project.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, err := getProject(rc, as)
		if err != nil {
			return "", err
		}
		ps.Title = fmt.Sprintf("Edit %s (project %s)", prj.Title(), prj.Key)
		ps.Data = prj
		return controller.Render(rc, as, &vproject.Edit{Project: prj}, ps, "projects", prj.Key)
	})
}

func ProjectSave(rc *fasthttp.RequestCtx) {
	controller.Act("project.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		prj, err := getProject(rc, as)
		if err != nil {
			return "", err
		}
		err = projectFromForm(frm, prj)
		if err != nil {
			return "", err
		}
		err = as.Services.Projects.Save(prj, ps.Logger)
		if err != nil {
			return controller.ERsp("unable to save project: %+v", err)
		}

		return controller.FlashAndRedir(true, "Saved changes", "/p/"+prj.Key, rc, ps)
	})
}

package checks

import (
	"fmt"

	"github.com/kyleu/projectforge/app/doctor"
	"github.com/kyleu/projectforge/app/project"
	"github.com/kyleu/projectforge/app/util"
	"go.uber.org/zap"
)

var prj = &doctor.Check{
	Key:     "project",
	Section: "app",
	Title:   "Project",
	Summary: "Verifies the Project Forge project in the working directory",
	URL:     util.AppURL,
	UsedBy:  util.AppName,
	Fn:      checkProject,
	Solve:   solveProject,
}

func checkProject(r *doctor.Result, logger *zap.SugaredLogger) *doctor.Result {
	p, r := loadProject(r, logger)
	if len(r.Errors) > 0 {
		return r
	}
	if p.Port == 0 {
		r = r.WithError(doctor.NewError("config", "port must be a non-zero integer"))
	}
	if p.Info == nil {
		return r.WithError(doctor.NewError("config", "project [%s] has no project info", p.Key))
	}
	if len(p.Modules) == 0 {
		r = r.WithError(doctor.NewError("config", "No modules enabled?!"))
	}

	if p.Build == nil {
		p.Build = &project.Build{}
	}

	r = checkMods(p, r)

	if p.Build.Notarize && p.Info.SigningIdentity == "" {
		r = r.WithError(doctor.NewError("config", "Notarizing is enabled, but [Signing Idenitity] isn't set"))
	}
	if p.Info.Homepage == "" {
		r = r.WithError(doctor.NewError("config", "No homepage set"))
	}
	if p.Info.License == "" {
		r = r.WithError(doctor.NewError("config", "No license set"))
	}
	if p.Info.AuthorName == "" {
		r = r.WithError(doctor.NewError("config", "No author name set"))
	}
	if p.Info.AuthorEmail == "" {
		r = r.WithError(doctor.NewError("config", "No author email set"))
	}
	return r
}

func solveProject(r *doctor.Result, logger *zap.SugaredLogger) *doctor.Result {
	if r.Errors.Find("missing") != nil {
		r.AddSolution("run [projectforge create] in this directory")
	}
	if r.Errors.Find("invalid") != nil {
		r.AddSolution("the project file isn't valid JSON, not sure what you can do")
	}
	if r.Errors.Find("config") != nil {
		r.AddSolution(fmt.Sprintf("use the %s UI to configure your project", util.AppName))
	}
	return r
}
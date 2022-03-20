package action

import (
	"context"
	"path/filepath"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"projectforge.dev/projectforge/app/lib/telemetry"

	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/app/util"
)

type Build struct {
	Key         string `json:"key"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	Run func(pm *PrjAndMods, ret *Result) *Result `json:"-"`
}

func simpleProc(cmd string, path string, ret *Result) *Result {
	exitCode, out, err := util.RunProcessSimple(cmd, path)
	if err != nil {
		return ret.WithError(err)
	}
	ret.AddLog("build output for [" + cmd + "]:\n" + out)
	if exitCode != 0 {
		ret.WithError(errors.Errorf("build failed with exit code [%d]", exitCode))
	}
	return ret
}

func simpleBuild(key string, title string, cmd string) *Build {
	return &Build{Key: key, Title: title, Description: "Runs [" + cmd + "]", Run: func(pm *PrjAndMods, ret *Result) *Result {
		return simpleProc(cmd, pm.Prj.Path, ret)
	}}
}

const ciDesc = "Installs dependencies for the TypeScript client"

type Builds []*Build

func (b Builds) Get(key string) *Build {
	for _, x := range b {
		if x.Key == key {
			return x
		}
	}
	return nil
}

var AllBuilds = Builds{
	simpleBuild("build", "Build", "make build"),
	simpleBuild("clean", "Clean", "make clean"),
	simpleBuild("tidy", "Tidy", "go mod tidy"),
	simpleBuild("format", "Format", "bin/format.sh"),
	simpleBuild("lint", "Lint", "bin/check.sh"),
	{Key: "deps", Title: "Dependencies", Description: "Manages Go dependencies", Run: onDeps},
	{Key: "clientInstall", Title: "Client Install", Description: ciDesc, Run: func(pm *PrjAndMods, ret *Result) *Result {
		return simpleProc("npm install", filepath.Join(pm.Prj.Path, "client"), ret)
	}},
	simpleBuild("clientBuild", "Client Build", "bin/build/client.sh"),
	{Key: "test", Title: "Test", Description: "Does a test", Run: func(pm *PrjAndMods, ret *Result) *Result {
		return simpleProc("ls", pm.Prj.Path, ret)
	}},
}

func onBuild(ctx context.Context, pm *PrjAndMods) *Result {
	phaseStr, _ := pm.Cfg.GetString("phase", true)
	if phaseStr == "" {
		phaseStr = "build"
	}

	ctx, span, logger := telemetry.StartSpan(ctx, "build:"+phaseStr, pm.Logger)
	defer span.Complete()
	pm.Logger = logger

	ret := newResult(pm.Cfg, logger)
	ret.AddLog("building project [%s] in [%s] with phase [%s]", pm.Prj.Key, pm.Prj.Path, phaseStr)
	var phase *Build
	for _, x := range AllBuilds {
		if x.Key == phaseStr {
			phase = x
			break
		}
	}
	if phase == nil {
		return ret.WithError(errors.Errorf("invalid phase [%s]", phaseStr))
	}
	timer := util.TimerStart()
	ret = phase.Run(pm, ret)
	ret.Duration = timer.End()
	return ret
}

func fullBuild(prj *project.Project, r *Result, logger *zap.SugaredLogger) *Result {
	r.AddLog("building project [%s] in [%s]", prj.Key, prj.Path)

	exitCode, out, err := util.RunProcessSimple("bin/templates.sh", prj.Path)
	if err != nil {
		return r.WithError(err)
	}
	r.AddLog("templates.sh output for [" + prj.Key + "]:\n" + out)
	if exitCode != 0 {
		return r.WithError(errors.Errorf("templates.sh failed with exit code [%d]", exitCode))
	}

	exitCode, out, err = util.RunProcessSimple("go mod tidy", prj.Path)
	if err != nil {
		return r.WithError(err)
	}
	r.AddLog("go mod output for [" + prj.Key + "]:\n" + out)
	if exitCode != 0 {
		return r.WithError(errors.Errorf("\"go mod tidy\" failed with exit code [%d]", exitCode))
	}

	exitCode, out, err = util.RunProcessSimple("npm install", filepath.Join(prj.Path, "client"))
	if err != nil {
		return r.WithError(err)
	}
	r.AddLog("npm output for [" + prj.Key + "]:\n" + out)
	if exitCode != 0 {
		r.WithError(errors.Errorf("npm install failed with exit code [%d]", exitCode))
	}

	exitCode, out, err = util.RunProcessSimple("bin/build/client.sh", prj.Path)
	if err != nil {
		return r.WithError(err)
	}
	r.AddLog("client build output for [" + prj.Key + "]:\n" + out)
	if exitCode != 0 {
		r.WithError(errors.Errorf("client build failed with exit code [%d]", exitCode))
	}

	exitCode, out, err = util.RunProcessSimple("make build", prj.Path)
	if err != nil {
		return r.WithError(err)
	}
	r.AddLog("build output for [" + prj.Key + "]:\n" + out)
	if exitCode != 0 {
		r.WithError(errors.Errorf("build failed with exit code [%d]", exitCode))
	}

	return r
}

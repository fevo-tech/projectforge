package cmd

import (
	"context"
	"strings"

	"projectforge.dev/projectforge/app/lib/exec"
	"projectforge.dev/projectforge/app/lib/filesystem"
	"projectforge.dev/projectforge/app/module"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/app/project/action"
	"projectforge.dev/projectforge/app/project/export"
	"projectforge.dev/projectforge/app/util"
)

func runToCompletion(ctx context.Context, projectKey string, t action.Type, cfg util.ValueMap) *action.Result {
	fs := filesystem.NewFileSystem(_flags.ConfigDir)
	mSvc := module.NewService(ctx, fs, _logger)
	pSvc := project.NewService()
	eSvc := export.NewService()
	xSvc := exec.NewService()
	logger := _logger.With("service", "runner")
	p := &action.Params{ProjectKey: projectKey, T: t, Cfg: cfg, MSvc: mSvc, PSvc: pSvc, SSvc: nil, XSvc: xSvc, ESvc: eSvc, CLI: true, Logger: logger}
	return action.Apply(ctx, p)
}

func extractConfig(args []string) util.ValueMap {
	var retArgs []string
	retMap := util.ValueMap{}
	for _, arg := range args {
		l, r := util.StringSplit(arg, '=', true)
		l = strings.TrimSpace(l)
		r = strings.TrimSpace(r)
		if r == "" {
			retArgs = append(retArgs, l)
		} else {
			retMap[l] = r
		}
	}
	if len(retArgs) > 0 {
		retMap["cmds"] = retArgs
	}
	return retMap
}

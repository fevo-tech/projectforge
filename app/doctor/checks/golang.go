package checks

import (
	"github.com/kyleu/projectforge/app/doctor"
	"go.uber.org/zap"
)

var golang = &doctor.Check{
	Key:     "golang",
	Section: "build",
	Title:   "Go",
	Summary: "The main programming language",
	URL:     "https://golang.org",
	UsedBy:  "All builds",
	Fn:      simpleOut(".", "go", []string{"version"}, noop),
	Solve:   solveGo,
}

func solveGo(r *doctor.Result, logger *zap.SugaredLogger) *doctor.Result {
	if r.Errors.Find("missing") != nil || r.Errors.Find("exitcode") != nil {
		r.AddSolution("Download Go for your platform")
	}
	return r
}
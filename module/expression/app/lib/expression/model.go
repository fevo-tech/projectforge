package expression

import (
	"fmt"
	"time"

	"github.com/google/cel-go/cel"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"{{{ .Package }}}/app/util"
)

type Expression struct {
	Key         string
	Description string
	Pattern     string
	Program     cel.Program
}

func NewExpression(key string, desc string, pattern string) *Expression {
	return &Expression{Key: key, Description: desc, Pattern: pattern}
}

func (e *Expression) Compile(eng *Engine) error {
	ast, issues := eng.env.Compile(e.Pattern)
	if issues != nil && issues.Err() != nil {
		return errors.Wrapf(issues.Err(), "compile error for pattern [%s]", e.Pattern)
	}
	prg, err := eng.env.Program(ast)
	if err != nil {
		return errors.Wrap(err, "program construction error")
	}

	e.Program = prg
	return nil
}

func (e *Expression) Run(params map[string]interface{}) (interface{}, error, int64) {
	if e.Program == nil {
		return nil, errors.New("no program"), 0
	}

	startNanos := util.TimerStart()
	out, _, err := e.Program.Eval(params)
	if err != nil {
		return nil, errors.Wrap(err, "cannot run program"), 0
	}
	duration := (time.Now().UnixNano() - startNanos) / int64(time.Microsecond)

	return out.Value(), nil, duration
}

func CheckResult(x interface{}, logger *zap.SugaredLogger) bool {
	switch t := x.(type) {
	case bool:
		return t
	default:
		logger.Info(fmt.Sprintf("invalid result type [%T]", x))
		return false
	}
}
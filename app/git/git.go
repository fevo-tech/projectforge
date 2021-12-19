package git

import (
	"github.com/kyleu/projectforge/app/util"
	"github.com/pkg/errors"
)

var noRepo = errors.New("not a git repository")

func gitCmd(args string, path string) (string, error) {
	exit, out, err := util.RunProcessSimple("git "+args, path)
	if err != nil {
		return "", errors.Wrap(err, "can't read git status for path ["+path+"]")
	}
	if exit == 128 {
		return "", errors.Wrapf(noRepo, "path [%s] is not a git repo", path)
	}
	if exit != 0 {
		return "", errors.Errorf("git status returned exit code [%d] for path [%s]", exit, path)
	}
	return out, nil
}

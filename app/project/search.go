package project

import (
	"github.com/kyleu/projectforge/app/lib/search/result"
)

func (s *Service) Search(q string) (result.Results, error) {
	ret := result.Results{}
	for _, p := range s.Projects() {
		if res := result.NewResult("project", p.Key, p.WebPath(), p.Title(), p.IconSafe(), p, q); len(res.Matches) > 0 {
			ret = append(ret, res)
		}
	}
	return ret, nil
}
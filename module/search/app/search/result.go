package search

import (
	"sort"
)

type Result struct {
	ID    string      `json:"id,omitempty"`
	Type  string      `json:"type,omitempty"`
	Title string      `json:"title,omitempty"`
	Icon  string      `json:"icon,omitempty"`
	URL   string      `json:"url,omitempty"`
	Match string      `json:"match,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	HTML  string      `json:"-"`
}

type Results []*Result

func (rs Results) Sort() {
	sort.Slice(rs, func(i, j int) bool {
		l, r := rs[i], rs[j]
		if l.Type == r.Type {
			return l.Title < r.Title
		}
		return l.Type < r.Type
	})
}

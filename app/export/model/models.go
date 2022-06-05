package model

import (
	"golang.org/x/exp/slices"
)

type Models []*Model

func (m Models) Get(n string) *Model {
	for _, x := range m {
		if x.Name == n {
			return x
		}
	}
	return nil
}

func (m Models) ReverseRelations(t string) Relations {
	var rels Relations
	for _, x := range m {
		for _, rel := range x.Relations {
			if rel.Table == t {
				rels = append(rels, rel.Reverse(x.Name))
			}
		}
	}
	return rels
}

func (m Models) Sort() {
	slices.SortFunc(m, func(l *Model, r *Model) bool {
		return l.Offset < r.Offset
	})
}

func (m Models) Sorted() Models {
	ret := make(Models, 0, len(m))
	for _, mdl := range m {
		if curr := m.Get(mdl.Name); curr != nil {
			ret = append(ret, m.withDeps(mdl)...)
		}
	}
	return ret
}

func (m Models) withDeps(mdl *Model) Models {
	var deps Models
	for _, rel := range m.ReverseRelations(mdl.Name) {
		if rel.Table == mdl.Name {
			deps = append(deps, m.withDeps(m.Get(rel.Table))...)
		}
	}
	return append(Models{mdl}, deps...)
}

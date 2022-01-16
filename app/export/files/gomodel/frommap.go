package gomodel

import (
	"github.com/kyleu/projectforge/app/export/golang"
	"github.com/kyleu/projectforge/app/export/model"
	"github.com/kyleu/projectforge/app/util"
)

func modelFromMap(m *model.Model) *golang.Block {
	ret := golang.NewBlock(m.Package+"FromForm", "func")
	ret.W("func FromMap(m util.ValueMap, setPK bool) (*%s, error) {", m.Proper())
	ret.W("\tret := &%s{}", m.Proper())
	ret.W("\tvar err error")
	ret.W("\tif setPK {")
	cols := m.Columns.WithoutTag("created").WithoutTag("updated").WithoutTag(model.RevisionType)
	forCols(ret, 2, cols.PKs()...)
	ret.W("\t\t// $PF_SECTION_START(pkchecks)$")
	ret.W("\t\t// $PF_SECTION_END(pkchecks)$")
	ret.W("\t}")
	forCols(ret, 1, cols.NonPKs()...)
	ret.W("\t// $PF_SECTION_START(extrachecks)$")
	ret.W("\t// $PF_SECTION_END(extrachecks)$")
	ret.W("\treturn ret, nil")
	ret.W("}")

	return ret
}

func forCols(ret *golang.Block, indent int, cols ...*model.Column) {
	ind := util.StringRepeat("\t", indent)
	catchErr := func() {
		ret.W(ind + "if err != nil {")
		ret.W(ind + "\treturn nil, err")
		ret.W(ind + "}")
	}
	for _, col := range cols {
		if col.Nullable || col.Type.IsScalar() {
			ret.W(ind+"ret.%s, err = m.Parse%s(%q, true, true)", col.Proper(), col.ToGoMapParse(), col.Camel())
			catchErr()
		} else {
			ret.W(ind+"ret%s, err := m.Parse%s(%q, true, true)", col.Proper(), col.ToGoMapParse(), col.Camel())
			catchErr()
			ret.W(ind+"ret.%s = *ret%s", col.Proper(), col.Proper())
		}
	}
}
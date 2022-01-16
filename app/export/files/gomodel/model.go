package gomodel

import (
	"fmt"
	"strings"

	"github.com/kyleu/projectforge/app/export/files/helper"
	"github.com/kyleu/projectforge/app/export/golang"
	"github.com/kyleu/projectforge/app/export/model"
	"github.com/kyleu/projectforge/app/file"
	"github.com/kyleu/projectforge/app/util"
)

func Model(m *model.Model, args *model.Args) (*file.File, error) {
	g := golang.NewFile(m.Package, []string{"app", m.Package}, m.Camel())
	for _, imp := range helper.ImportsForTypes("go", m.Columns.Types()...) {
		g.AddImport(imp)
	}
	for _, imp := range helper.ImportsForTypes("string", m.PKs().Types()...) {
		g.AddImport(imp)
	}
	if len(m.PKs()) > 1 {
		g.AddImport(helper.ImpFmt)
	}
	g.AddImport(helper.ImpAppUtil)
	g.AddBlocks(modelStruct(m), modelConstructor(m), modelRandom(m), modelFromMap(m))
	g.AddBlocks(modelString(m), modelWebPath(m), modelToData(m, m.Columns, ""))
	if m.IsRevision() {
		hc := m.HistoryColumns(false)
		g.AddBlocks(modelToData(m, hc.Const, "Core"), modelToData(m, hc.Var, hc.Col.Proper()))
	}
	g.AddBlocks(modelArray(m))
	return g.Render()
}

func modelStruct(m *model.Model) *golang.Block {
	ret := golang.NewBlock(m.Proper(), "struct")
	ret.W("type %s struct {", m.Proper())
	maxColLength := util.StringArrayMaxLength(m.Columns.CamelNames())
	maxTypeLength := m.Columns.MaxGoKeyLength()
	for _, c := range m.Columns {
		suffix := ""
		if c.Nullable {
			suffix = ",omitempty"
		}
		ret.W("\t%s %s `json:%q`", util.StringPad(c.Proper(), maxColLength), util.StringPad(c.ToGoType(), maxTypeLength), c.Camel()+suffix)
	}
	ret.W("}")
	return ret
}

func modelConstructor(m *model.Model) *golang.Block {
	ret := golang.NewBlock("New"+m.Proper(), "func")
	ret.W("func New(%s) *%s {", m.PKs().Args(), m.Proper())
	ret.W("\treturn &%s{%s}", m.Proper(), m.PKs().Refs())
	ret.W("}")
	return ret
}

func modelToData(m *model.Model, cols model.Columns, suffix string) *golang.Block {
	ret := golang.NewBlock(m.Proper(), "func")
	ret.W("func (%s *%s) ToData%s() []interface{} {", m.FirstLetter(), m.Proper(), suffix)
	refs := make([]string, 0, len(cols))
	for _, c := range cols {
		refs = append(refs, fmt.Sprintf("%s.%s", m.FirstLetter(), c.Proper()))
	}
	ret.W("\treturn []interface{}{%s}", strings.Join(refs, ", "))
	ret.W("}")
	return ret
}

func modelArray(m *model.Model) *golang.Block {
	ret := golang.NewBlock(m.Proper()+"Array", "type")
	ret.W("type %s []*%s", m.ProperPlural(), m.Proper())
	return ret
}
package gomodel

import (
	"fmt"
	"strings"

	"projectforge.dev/projectforge/app/lib/types"
	"projectforge.dev/projectforge/app/project/export/files/helper"
	"projectforge.dev/projectforge/app/project/export/golang"
	"projectforge.dev/projectforge/app/project/export/model"
	"projectforge.dev/projectforge/app/util"
)

func modelClone(m *model.Model) *golang.Block {
	ret := golang.NewBlock("Clone", "func")
	ret.W("func (%s *%s) Clone() *%s {", m.FirstLetter(), m.Proper(), m.Proper())
	ret.W("\treturn &%s{", m.Proper())
	max := m.Columns.MaxCamelLength() + 1
	for _, col := range m.Columns {
		decl := col.Proper()
		switch col.Type.Key() {
		case types.KeyMap, types.KeyValueMap, types.KeyReference:
			decl += ".Clone()"
		}
		ret.W("\t\t%s %s.%s,", util.StringPad(col.Proper()+":", max), m.FirstLetter(), decl)
	}
	ret.W("\t}")
	ret.W("}")
	return ret
}

func modelString(g *golang.File, m *model.Model) *golang.Block {
	ret := golang.NewBlock("String", "func")
	ret.W("func (%s *%s) String() string {", m.FirstLetter(), m.Proper())
	if pks := m.PKs(); len(pks) == 1 {
		ret.W("\treturn %s", model.ToGoString(pks[0].Type, fmt.Sprintf("%s.%s", m.FirstLetter(), pks[0].Proper()), false))
	} else {
		g.AddImport(helper.ImpFmt)
		s := "\treturn fmt.Sprintf(\""
		for idx := range m.PKs() {
			if idx > 0 {
				s += "::"
			}
			s += "%%s"
		}
		s += "\""
		for _, c := range m.PKs() {
			s += ", " + c.ToGoString(m.FirstLetter()+".")
		}
		ret.W(s + ")")
	}
	ret.W("}")
	return ret
}

func modelTitle(m *model.Model) *golang.Block {
	ret := golang.NewBlock("Title", "func")
	ret.W("func (%s *%s) TitleString() string {", m.FirstLetter(), m.Proper())
	if titles := m.Columns.WithTag("title"); len(titles) > 0 {
		var toStrings []string
		for _, title := range titles {
			toStrings = append(toStrings, model.ToGoString(title.Type, fmt.Sprintf("%s.%s", m.FirstLetter(), title.Proper()), true))
		}
		ret.W("\treturn %s", strings.Join(toStrings, " + \" / \" + "))
	} else {
		ret.W("\treturn %s.String()", m.FirstLetter())
	}
	ret.W("}")
	return ret
}

func modelWebPath(g *golang.File, m *model.Model) *golang.Block {
	ret := golang.NewBlock("WebPath", "type")
	ret.W("func (%s *%s) WebPath() string {", m.FirstLetter(), m.Proper())
	p := "\"/" + m.Route() + "\""
	for _, pk := range m.PKs() {
		if strings.HasSuffix(p, "\"") {
			p = p[:len(p)-1] + "/" + "\" + "
		} else {
			p += " + \"/\" + "
		}
		if types.IsStringList(pk.Type) {
			g.AddImport(helper.ImpStrings)
			p += fmt.Sprintf(`strings.Join(%s, ",")`, pk.ToGoString(m.FirstLetter()+"."))
		} else {
			p += pk.ToGoString(m.FirstLetter() + ".")
		}
	}
	ret.W("\treturn " + p)
	ret.W("}")
	return ret
}

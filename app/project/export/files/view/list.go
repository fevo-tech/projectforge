package view

import (
	"projectforge.dev/projectforge/app/file"
	"projectforge.dev/projectforge/app/project/export/files/helper"
	"projectforge.dev/projectforge/app/project/export/golang"
	"projectforge.dev/projectforge/app/project/export/model"
)

func list(m *model.Model, args *model.Args, addHeader bool) (*file.File, error) {
	g := golang.NewGoTemplate([]string{"views", m.PackageWithGroup("v")}, "List.html")
	g.AddImport(helper.ImpApp, helper.ImpComponents, helper.ImpCutil, helper.ImpFilter, helper.ImpLayout)
	g.AddImport(helper.AppImport("app/" + m.PackageWithGroup("")))
	g.AddBlocks(exportViewListClass(m, args.Models, g), exportViewListBody(m, args.Models))
	return g.Render(addHeader)
}

func exportViewListClass(m *model.Model, models model.Models, g *golang.Template) *golang.Block {
	ret := golang.NewBlock("List", "struct")
	ret.W("{%% code type List struct {")
	ret.W("  layout.Basic")
	ret.W("  Models %s.%s", m.Package, m.ProperPlural())
	for _, rel := range m.Relations {
		relModel := models.Get(rel.Table)
		g.AddImport(helper.AppImport("app/" + relModel.PackageWithGroup("")))
		ret.W(commonLine, relModel.ProperPlural(), relModel.Package, relModel.ProperPlural())
	}
	ret.W("  Params filter.ParamSet")
	if len(m.AllSearches()) > 0 {
		ret.W("  SearchQuery string")
	}
	ret.W("} %%}")
	return ret
}

func exportViewListBody(m *model.Model, models model.Models) *golang.Block {
	ret := golang.NewBlock("ListBody", "func")

	suffix := ""
	for _, rel := range m.Relations {
		if relModel := models.Get(rel.Table); relModel.CanTraverseRelation() {
			suffix += ", p." + relModel.ProperPlural()
		}
	}

	ret.W("{%% func (p *List) Body(as *app.State, ps *cutil.PageState) %%}")
	ret.W("  <div class=\"card\">")
	if len(m.AllSearches()) == 0 {
		ret.W("    <div class=\"right\"><a href=\"/%s/new\"><button>New</button></a></div>", m.Route())
		ret.W("    <h3>{%%= components.SVGRefIcon(`" + m.Icon + "`, ps) %%}{%%s ps.Title %%}</h3>")
	} else {
		ret.W(`    <div class="right">{%%%%= components.SearchForm("", "q", "Search %s", p.SearchQuery, ps) %%%%}</div>`, m.Plural())
		ret.W("    <h3>{%%%%= components.SVGRefIcon(`"+m.Icon+"`, ps) %%%%}{%%%%s ps.Title %%%%} <a href=\"/%s/new\"><button>New</button></a></h3>", m.Route())
		ret.W("    <div class=\"clear\"></div>")
		ret.W("    {%%- if p.SearchQuery != \"\" -%%}")
		ret.W("    <em>Search results for [{%%s p.SearchQuery %%}]</em>")
		ret.W("    {%%- endif -%%}")
	}
	ret.W("    {%%- if len(p.Models) == 0 -%%}")
	ret.W("    <div class=\"mt\"><em>No %s available</em></div>", m.TitlePluralLower())
	ret.W("    {%%- else -%%}")
	ret.W("    <div class=\"overflow clear\">")
	ret.W("      {%%= Table(p.Models" + suffix + ", p.Params, as, ps) %%}")
	ret.W("    </div>")
	ret.W("    {%%- endif -%%}")
	ret.W("  </div>")
	ret.W("{%% endfunc %%}")
	return ret
}

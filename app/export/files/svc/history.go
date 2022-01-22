package svc

import (
	"fmt"
	"strings"

	"github.com/kyleu/projectforge/app/export/files/helper"
	"github.com/kyleu/projectforge/app/export/golang"
	"github.com/kyleu/projectforge/app/export/model"
	"github.com/kyleu/projectforge/app/file"
	"github.com/kyleu/projectforge/app/util"
)

func ServiceHistory(m *model.Model, args *model.Args, addHeader bool) (*file.File, error) {
	g := golang.NewFile(m.Package, []string{"app", m.Package}, "servicehistory")
	g.AddImport(helper.ImpContext, helper.ImpUUID, helper.ImpErrors, helper.ImpFmt, helper.ImpTime, helper.ImpStrings)
	g.AddImport(helper.ImpSQLx, helper.ImpAppUtil, helper.ImpDatabase)
	g.AddBlocks(serviceHistoryVars(m), serviceHistoryGetHistory(m), serviceHistoryGetHistories(m), serviceHistorySaveHistory(m))
	return g.Render(addHeader)
}

func serviceHistoryVars(m *model.Model) *golang.Block {
	ret := golang.NewBlock("HistoryVars", "func")
	ret.W("var (")
	ret.W("\thistoryColumns       = " + `[]string{"id", "history_id", "o", "n", "c", "created"}`)
	ret.W("\thistoryColumnsQuoted = util.StringArrayQuoted(historyColumns)")
	ret.W("\thistoryColumnsString = strings.Join(historyColumnsQuoted, \", \")")
	ret.W("")
	ret.W("\thistoryTable       = table + \"_history\"")
	ret.W("\thistoryTableQuoted = fmt.Sprintf(\"%%q\", historyTable)")
	ret.W(")")
	return ret
}

func serviceHistoryGetHistory(m *model.Model) *golang.Block {
	ret := golang.NewBlock("GetHistory", "func")
	ret.W("func (s *Service) GetHistory(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (*HistoryHistory, error) {")
	ret.W("\tq := database.SQLSelectSimple(historyColumnsString, historyTableQuoted, \"id = $1\")")
	ret.W("\tret := historyDTO{}")
	ret.W("\terr := s.db.Get(ctx, &ret, q, tx)")
	ret.W("\tif err != nil {")
	ret.W("\t\treturn nil, errors.Wrapf(err, \"unable to get %s history [%%%%s]\", id.String())", m.TitleLower())
	ret.W("\t}")
	ret.W("\treturn ret.ToHistory(), nil")
	ret.W("}")
	return ret
}

func serviceHistoryGetHistories(m *model.Model) *golang.Block {
	ret := golang.NewBlock("GetHistories", "func")
	ret.W("func (s *Service) GetHistories(ctx context.Context, tx *sqlx.Tx, id string) (HistoryHistories, error) {")
	pks := m.PKs()
	joins := make([]string, 0, len(pks))
	logs := make([]string, 0, len(pks))
	for idx, pk := range pks {
		joins = append(joins, fmt.Sprintf("%s_%s = $%d", m.Name, pk.Name, idx+1))
		logs = append(logs, pk.Camel()+" [%%v]")
	}
	ret.W("\tq := database.SQLSelectSimple(historyColumnsString, historyTableQuoted, %q)", strings.Join(joins, " and "))
	ret.W("\tret := historyDTOs{}")
	ret.W("\terr := s.db.Select(ctx, &ret, q, tx, %s)", strings.Join(pks.CamelNames(), ", "))
	ret.W("\tif err != nil {")
	msg := "\t\treturn nil, errors.Wrapf(err, \"unable to get %s by %s\", %s)"
	ret.W(msg, m.TitlePluralLower(), strings.Join(logs, ", "), strings.Join(pks.CamelNames(), ", "))
	ret.W("\t}")
	ret.W("\treturn ret.ToHistories(), nil")
	ret.W("}")
	return ret
}

func serviceHistorySaveHistory(m *model.Model) *golang.Block {
	ret := golang.NewBlock("SaveHistory", "func")
	decl := "func (s *Service) SaveHistory(ctx context.Context, tx *sqlx.Tx, o *%s, n *%s) (*%sHistory, error) {"
	ret.W(decl, m.Proper(), m.Proper(), m.Proper())
	ret.W("\tq := database.SQLInsert(historyTableQuoted, historyColumns, 1, \"\")")
	ret.W("\th := &historyDTO{")
	max := m.PKs().MaxCamelLength() + len(m.Proper()) + 1
	ret.W("\t\t%s util.UUID(),", util.StringPad("ID:", max))
	for _, pk := range m.PKs() {
		ret.W("\t\t%s o.%s,", util.StringPad(m.Proper()+pk.Proper()+":", max), pk.Proper())
	}
	ret.W("\t\t%s util.ToJSONBytes(o, true),", util.StringPad("Old:", max))
	ret.W("\t\t%s util.ToJSONBytes(n, true),", util.StringPad("New:", max))
	ret.W("\t\t%s util.ToJSONBytes(o.Diff(n), true),", util.StringPad("Changes:", max))
	ret.W("\t\t%s time.Now(),", util.StringPad("Created:", max))
	ret.W("\t}")
	ret.W("\thist := h.ToHistory()")
	ret.W("\terr := s.db.Insert(ctx, q, tx, hist.ToData()...)")
	ret.W("\tif err != nil {")
	ret.W("\t\treturn nil, errors.Wrap(err, \"unable to insert %s\")", m.TitleLower())
	ret.W("\t}")
	ret.W("\treturn hist, nil")
	ret.W("}")
	return ret
}

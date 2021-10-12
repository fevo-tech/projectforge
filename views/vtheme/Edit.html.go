// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Code generated by Project Forge, see https://projectforge.dev for details. -->

//line views/vtheme/Edit.html:2
package vtheme

//line views/vtheme/Edit.html:2
import (
	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/controller/cutil"
	"github.com/kyleu/projectforge/app/theme"
	"github.com/kyleu/projectforge/app/util"
	"github.com/kyleu/projectforge/views/components"
	"github.com/kyleu/projectforge/views/layout"
)

//line views/vtheme/Edit.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/Edit.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/Edit.html:11
type Edit struct {
	layout.Basic
	Theme *theme.Theme
}

//line views/vtheme/Edit.html:16
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/Edit.html:16
	qw422016.N().S(`
  <form action="" method="post">
    <div class="card">
      <div class="right"><a href="#modal-theme"><button type="button">JSON</button></a></div>
      <h3>`)
//line views/vtheme/Edit.html:20
	if p.Theme.Key == theme.KeyNew {
//line views/vtheme/Edit.html:20
		qw422016.N().S(`New Theme`)
//line views/vtheme/Edit.html:20
	} else {
//line views/vtheme/Edit.html:20
		qw422016.N().S(`Theme Edit`)
//line views/vtheme/Edit.html:20
	}
//line views/vtheme/Edit.html:20
	qw422016.N().S(`</h3>
      <table class="mt">
        <tbody>
          `)
//line views/vtheme/Edit.html:23
	components.StreamTableInput(qw422016, "key", "Key", p.Theme.Key, 5)
//line views/vtheme/Edit.html:23
	qw422016.N().S(`
        </tbody>
      </table>
    </div>
    `)
//line views/vtheme/Edit.html:27
	StreamEditor(qw422016, util.AppName, p.Theme, as, ps)
//line views/vtheme/Edit.html:27
	qw422016.N().S(`
    <div class="card">
      <button type="submit">Save All Changes</button>
      <a href="/theme/`)
//line views/vtheme/Edit.html:30
	qw422016.E().S(p.Theme.Key)
//line views/vtheme/Edit.html:30
	qw422016.N().S(`"><button type="button">Reset</button></a>
    </div>
  </form>
  `)
//line views/vtheme/Edit.html:33
	components.StreamJSONModal(qw422016, "theme", "Theme JSON", p.Theme, 1)
//line views/vtheme/Edit.html:33
	qw422016.N().S(`
`)
//line views/vtheme/Edit.html:34
}

//line views/vtheme/Edit.html:34
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/Edit.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Edit.html:34
	p.StreamBody(qw422016, as, ps)
//line views/vtheme/Edit.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Edit.html:34
}

//line views/vtheme/Edit.html:34
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vtheme/Edit.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Edit.html:34
	p.WriteBody(qb422016, as, ps)
//line views/vtheme/Edit.html:34
	qs422016 := string(qb422016.B)
//line views/vtheme/Edit.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Edit.html:34
	return qs422016
//line views/vtheme/Edit.html:34
}

// Code generated by qtc from "Debug.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Code generated by Project Forge, see https://projectforge.dev for details. -->

//line views/Debug.html:2
package views

//line views/Debug.html:2
import (
	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/controller/cutil"
	"github.com/kyleu/projectforge/app/util"
	"github.com/kyleu/projectforge/views/layout"
)

//line views/Debug.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Debug.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Debug.html:9
type Debug struct{ layout.Basic }

//line views/Debug.html:11
func (p *Debug) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Debug.html:11
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/Debug.html:13
	qw422016.E().S(ps.Title)
//line views/Debug.html:13
	qw422016.N().S(`</h3>
    <pre>`)
//line views/Debug.html:14
	qw422016.E().S(util.ToJSON(ps.Data))
//line views/Debug.html:14
	qw422016.N().S(`</pre>
  </div>
`)
//line views/Debug.html:16
}

//line views/Debug.html:16
func (p *Debug) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Debug.html:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Debug.html:16
	p.StreamBody(qw422016, as, ps)
//line views/Debug.html:16
	qt422016.ReleaseWriter(qw422016)
//line views/Debug.html:16
}

//line views/Debug.html:16
func (p *Debug) Body(as *app.State, ps *cutil.PageState) string {
//line views/Debug.html:16
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Debug.html:16
	p.WriteBody(qb422016, as, ps)
//line views/Debug.html:16
	qs422016 := string(qb422016.B)
//line views/Debug.html:16
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Debug.html:16
	return qs422016
//line views/Debug.html:16
}

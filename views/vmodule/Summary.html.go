// Code generated by qtc from "Summary.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vmodule/Summary.html:1
package vmodule

//line views/vmodule/Summary.html:1
import (
	"github.com/kyleu/projectforge/app/controller/cutil"
	"github.com/kyleu/projectforge/app/module"
	"github.com/kyleu/projectforge/views/components"
	"strings"
)

//line views/vmodule/Summary.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vmodule/Summary.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vmodule/Summary.html:8
func StreamSummary(qw422016 *qt422016.Writer, mod *module.Module, ps *cutil.PageState, path ...string) {
//line views/vmodule/Summary.html:8
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-module"><button type="button">JSON</button></a></div>
    <h3>`)
//line views/vmodule/Summary.html:11
	components.StreamSVGRefIcon(qw422016, mod.SafeIcon(), ps)
//line views/vmodule/Summary.html:11
	qw422016.N().S(` `)
//line views/vmodule/Summary.html:11
	qw422016.E().S(mod.Title())
//line views/vmodule/Summary.html:11
	qw422016.N().S(`</h3>
    <div class="mt">
`)
//line views/vmodule/Summary.html:13
	if len(path) == 0 {
//line views/vmodule/Summary.html:13
		qw422016.N().S(`      <a href="/m/`)
//line views/vmodule/Summary.html:14
		qw422016.E().S(mod.Key)
//line views/vmodule/Summary.html:14
		qw422016.N().S(`/fs"><button>Filesystem</button></a>
`)
//line views/vmodule/Summary.html:15
	} else {
//line views/vmodule/Summary.html:16
		var ctx []string

//line views/vmodule/Summary.html:17
		for _, pth := range path {
//line views/vmodule/Summary.html:18
			ctx = append(ctx, pth)

//line views/vmodule/Summary.html:18
			qw422016.N().S(`      <a href="/m/`)
//line views/vmodule/Summary.html:19
			qw422016.E().S(mod.Key)
//line views/vmodule/Summary.html:19
			qw422016.N().S(`/fs/`)
//line views/vmodule/Summary.html:19
			qw422016.E().S(strings.Join(ctx, `/`))
//line views/vmodule/Summary.html:19
			qw422016.N().S(`"><button>`)
//line views/vmodule/Summary.html:19
			qw422016.E().S(pth)
//line views/vmodule/Summary.html:19
			qw422016.N().S(`}</button></a>
`)
//line views/vmodule/Summary.html:20
		}
//line views/vmodule/Summary.html:21
	}
//line views/vmodule/Summary.html:21
	qw422016.N().S(`    </div>
  </div>
  `)
//line views/vmodule/Summary.html:24
	components.StreamJSONModal(qw422016, "module", "Module JSON", mod, 1)
//line views/vmodule/Summary.html:24
	qw422016.N().S(`
`)
//line views/vmodule/Summary.html:25
}

//line views/vmodule/Summary.html:25
func WriteSummary(qq422016 qtio422016.Writer, mod *module.Module, ps *cutil.PageState, path ...string) {
//line views/vmodule/Summary.html:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vmodule/Summary.html:25
	StreamSummary(qw422016, mod, ps, path...)
//line views/vmodule/Summary.html:25
	qt422016.ReleaseWriter(qw422016)
//line views/vmodule/Summary.html:25
}

//line views/vmodule/Summary.html:25
func Summary(mod *module.Module, ps *cutil.PageState, path ...string) string {
//line views/vmodule/Summary.html:25
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vmodule/Summary.html:25
	WriteSummary(qb422016, mod, ps, path...)
//line views/vmodule/Summary.html:25
	qs422016 := string(qb422016.B)
//line views/vmodule/Summary.html:25
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vmodule/Summary.html:25
	return qs422016
//line views/vmodule/Summary.html:25
}

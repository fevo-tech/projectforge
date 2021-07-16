// Code generated by qtc from "Result.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Result.html:1
package vaction

//line views/vaction/Result.html:1
import (
	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/action"
	"github.com/kyleu/projectforge/app/controller/cutil"
	"github.com/kyleu/projectforge/app/util"
	"github.com/kyleu/projectforge/views/components"
	"github.com/kyleu/projectforge/views/layout"
)

//line views/vaction/Result.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Result.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Result.html:10
type Result struct {
	layout.Basic
	Cfg    util.ValueMap
	Result *action.Result
}

//line views/vaction/Result.html:16
func (p *Result) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Result.html:16
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-cfg"><button type="button">Config</button></a>
      <a href="#modal-result"><button type="button">Result</button></a>
    </div>
    <h3>Result</h3>
    <p>`)
//line views/vaction/Result.html:23
	qw422016.E().S(p.Result.Status)
//line views/vaction/Result.html:23
	qw422016.N().S(`</p>
  </div>

  <div class="card">
    <h3>Config</h3>
    <table>
      <tbody>
`)
//line views/vaction/Result.html:30
	for _, k := range p.Cfg.Keys() {
//line views/vaction/Result.html:30
		qw422016.N().S(`        <tr>
          <td>`)
//line views/vaction/Result.html:32
		qw422016.E().S(k)
//line views/vaction/Result.html:32
		qw422016.N().S(`</td>
          <td>`)
//line views/vaction/Result.html:33
		qw422016.E().V(p.Cfg[k])
//line views/vaction/Result.html:33
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vaction/Result.html:35
	}
//line views/vaction/Result.html:35
	qw422016.N().S(`      </tbody>
    </table>
  </div>

  <div class="card">
    <h3>Errors</h3>
`)
//line views/vaction/Result.html:42
	if len(p.Result.Errors) == 0 {
//line views/vaction/Result.html:42
		qw422016.N().S(`    <p>no errors</p>
`)
//line views/vaction/Result.html:44
	} else {
//line views/vaction/Result.html:44
		qw422016.N().S(`    <table>
      <tbody>
`)
//line views/vaction/Result.html:47
		for _, e := range p.Result.Errors {
//line views/vaction/Result.html:47
			qw422016.N().S(`        <tr>
          <td>`)
//line views/vaction/Result.html:49
			qw422016.E().S(e)
//line views/vaction/Result.html:49
			qw422016.N().S(`</td>
        </tr>
`)
//line views/vaction/Result.html:51
		}
//line views/vaction/Result.html:51
		qw422016.N().S(`      </tbody>
    </table>
`)
//line views/vaction/Result.html:54
	}
//line views/vaction/Result.html:54
	qw422016.N().S(`  </div>

  <div class="card">
    <h3>Logs</h3>
`)
//line views/vaction/Result.html:59
	if len(p.Result.Logs) == 0 {
//line views/vaction/Result.html:59
		qw422016.N().S(`    <p>no logs</p>
`)
//line views/vaction/Result.html:61
	} else {
//line views/vaction/Result.html:61
		qw422016.N().S(`    <table>
      <tbody>
`)
//line views/vaction/Result.html:64
		for _, l := range p.Result.Logs {
//line views/vaction/Result.html:64
			qw422016.N().S(`        <tr>
          <td><pre>`)
//line views/vaction/Result.html:66
			qw422016.E().S(l)
//line views/vaction/Result.html:66
			qw422016.N().S(`</pre></td>
        </tr>
`)
//line views/vaction/Result.html:68
		}
//line views/vaction/Result.html:68
		qw422016.N().S(`      </tbody>
    </table>
`)
//line views/vaction/Result.html:71
	}
//line views/vaction/Result.html:71
	qw422016.N().S(`  </div>

`)
//line views/vaction/Result.html:74
	for _, mr := range p.Result.Modules {
//line views/vaction/Result.html:74
		qw422016.N().S(`    <div class="card">
      <div class="right">`)
//line views/vaction/Result.html:76
		qw422016.E().S(mr.Status)
//line views/vaction/Result.html:76
		qw422016.N().S(`</div>
      <h3>Module [`)
//line views/vaction/Result.html:77
		qw422016.E().S(mr.Key)
//line views/vaction/Result.html:77
		qw422016.N().S(`]</h3>
      <h4>Diffs</h4>
      <table>
        <thead>
          <tr>
            <th>Path</th>
            <th>Status</th>
            <th>Patch</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vaction/Result.html:88
		for _, diff := range mr.Diffs {
//line views/vaction/Result.html:88
			qw422016.N().S(`          <tr>
            <td>`)
//line views/vaction/Result.html:90
			qw422016.E().S(diff.Path)
//line views/vaction/Result.html:90
			qw422016.N().S(`</td>
            <td>`)
//line views/vaction/Result.html:91
			qw422016.E().S(diff.Status)
//line views/vaction/Result.html:91
			qw422016.N().S(`</td>
            <td><pre>`)
//line views/vaction/Result.html:92
			qw422016.N().S(diff.Patch)
//line views/vaction/Result.html:92
			qw422016.N().S(`</pre></td>
          </tr>
`)
//line views/vaction/Result.html:94
		}
//line views/vaction/Result.html:94
		qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/vaction/Result.html:98
	}
//line views/vaction/Result.html:98
	qw422016.N().S(`
  `)
//line views/vaction/Result.html:100
	components.StreamJSONModal(qw422016, "cfg", "Config JSON", p.Cfg, 1)
//line views/vaction/Result.html:100
	qw422016.N().S(`
  `)
//line views/vaction/Result.html:101
	components.StreamJSONModal(qw422016, "result", "Result JSON", p.Result, 1)
//line views/vaction/Result.html:101
	qw422016.N().S(`
`)
//line views/vaction/Result.html:102
}

//line views/vaction/Result.html:102
func (p *Result) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Result.html:102
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Result.html:102
	p.StreamBody(qw422016, as, ps)
//line views/vaction/Result.html:102
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Result.html:102
}

//line views/vaction/Result.html:102
func (p *Result) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/Result.html:102
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Result.html:102
	p.WriteBody(qb422016, as, ps)
//line views/vaction/Result.html:102
	qs422016 := string(qb422016.B)
//line views/vaction/Result.html:102
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Result.html:102
	return qs422016
//line views/vaction/Result.html:102
}

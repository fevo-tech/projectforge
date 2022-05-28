// Code generated by qtc from "SearchAll.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vproject/SearchAll.html:1
package vproject

//line views/vproject/SearchAll.html:1
import (
	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/lib/search"
	"projectforge.dev/projectforge/app/lib/search/result"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
)

//line views/vproject/SearchAll.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vproject/SearchAll.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vproject/SearchAll.html:11
type SearchAll struct {
	layout.Basic
	Params   *search.Params
	Projects project.Projects
	Results  map[string]result.Results
	Tags     []string
}

//line views/vproject/SearchAll.html:19
func (p *SearchAll) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/SearchAll.html:19
	qw422016.N().S(`
  <div class="card">
    `)
//line views/vproject/SearchAll.html:21
	StreamAvailActions(qw422016, "Search", p.Tags, p.Projects.Tags(), ps)
//line views/vproject/SearchAll.html:21
	qw422016.N().S(`
  </div>
  <div class="card">
    <ul class="accordion">
`)
//line views/vproject/SearchAll.html:25
	for _, prj := range p.Projects {
//line views/vproject/SearchAll.html:26
		res := p.Results[prj.Key]

//line views/vproject/SearchAll.html:26
		qw422016.N().S(`      <li>
        <input id="accordion-`)
//line views/vproject/SearchAll.html:28
		qw422016.E().S(prj.Key)
//line views/vproject/SearchAll.html:28
		qw422016.N().S(`" type="checkbox" hidden />
        <label for="accordion-`)
//line views/vproject/SearchAll.html:29
		qw422016.E().S(prj.Key)
//line views/vproject/SearchAll.html:29
		qw422016.N().S(`">
          <div class="right">
`)
//line views/vproject/SearchAll.html:31
		if len(res) == 0 {
//line views/vproject/SearchAll.html:31
			qw422016.N().S(`            <em>0 matches</em>
`)
//line views/vproject/SearchAll.html:33
		} else {
//line views/vproject/SearchAll.html:33
			qw422016.N().S(`            `)
//line views/vproject/SearchAll.html:34
			qw422016.N().D(len(res))
//line views/vproject/SearchAll.html:34
			qw422016.N().S(` matches
`)
//line views/vproject/SearchAll.html:35
		}
//line views/vproject/SearchAll.html:35
		qw422016.N().S(`          </div>
          `)
//line views/vproject/SearchAll.html:37
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vproject/SearchAll.html:37
		qw422016.N().S(` `)
//line views/vproject/SearchAll.html:37
		components.StreamSVGRef(qw422016, prj.IconSafe(), 16, 16, "icon", ps)
//line views/vproject/SearchAll.html:37
		qw422016.N().S(` `)
//line views/vproject/SearchAll.html:37
		qw422016.E().S(prj.Title())
//line views/vproject/SearchAll.html:37
		qw422016.N().S(`
        </label>
        <div class="bd">
          `)
//line views/vproject/SearchAll.html:40
		StreamSummary(qw422016, prj, nil, nil, nil, ps)
//line views/vproject/SearchAll.html:40
		qw422016.N().S(`
          `)
//line views/vproject/SearchAll.html:41
		streamsearchResults(qw422016, prj, p.Params, res, nil, as, ps)
//line views/vproject/SearchAll.html:41
		qw422016.N().S(`
        </div>
      </li>
`)
//line views/vproject/SearchAll.html:44
	}
//line views/vproject/SearchAll.html:44
	qw422016.N().S(`    </ul>
  </div>
</div>
`)
//line views/vproject/SearchAll.html:48
}

//line views/vproject/SearchAll.html:48
func (p *SearchAll) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/SearchAll.html:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vproject/SearchAll.html:48
	p.StreamBody(qw422016, as, ps)
//line views/vproject/SearchAll.html:48
	qt422016.ReleaseWriter(qw422016)
//line views/vproject/SearchAll.html:48
}

//line views/vproject/SearchAll.html:48
func (p *SearchAll) Body(as *app.State, ps *cutil.PageState) string {
//line views/vproject/SearchAll.html:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vproject/SearchAll.html:48
	p.WriteBody(qb422016, as, ps)
//line views/vproject/SearchAll.html:48
	qs422016 := string(qb422016.B)
//line views/vproject/SearchAll.html:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vproject/SearchAll.html:48
	return qs422016
//line views/vproject/SearchAll.html:48
}

// Code generated by qtc from "Results.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vgit/Results.html:1
package vgit

//line views/vgit/Results.html:1
import (
	"strings"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/git"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vproject"
)

//line views/vgit/Results.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vgit/Results.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vgit/Results.html:13
type Results struct {
	layout.Basic
	Action   *git.Action
	Results  git.Results
	Projects project.Projects
	Tags     []string
}

//line views/vgit/Results.html:21
func (p *Results) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgit/Results.html:21
	qw422016.N().S(`
`)
//line views/vgit/Results.html:23
	var tags string
	if len(p.Tags) > 0 {
		tags = "?tags=" + strings.Join(p.Tags, ",")
	}

//line views/vgit/Results.html:27
	qw422016.N().S(`  <div class="card">
    `)
//line views/vgit/Results.html:29
	vproject.StreamAvailActions(qw422016, "Git "+p.Action.Title, p.Tags, p.Projects.Tags(), ps)
//line views/vgit/Results.html:29
	qw422016.N().S(`
  </div>
  <div class="card">
    <table class="min-200">
      <thead>
        <tr>
          <th class="shrink">Git Action</th>
          <th>Description</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><a href="/git`)
//line views/vgit/Results.html:41
	qw422016.E().S(tags)
//line views/vgit/Results.html:41
	qw422016.N().S(`"><button>Status</button></a></td>
          <td>Build status for all projects</td>
        </tr>
        <tr>
          <td><a href="/git/all/fetch`)
//line views/vgit/Results.html:45
	qw422016.E().S(tags)
//line views/vgit/Results.html:45
	qw422016.N().S(`"><button>Fetch</button></a></td>
          <td>Fetches updates for all projects</td>
        </tr>
        <tr>
          <td><a href="/git/all/magic`)
//line views/vgit/Results.html:49
	qw422016.E().S(tags)
//line views/vgit/Results.html:49
	qw422016.N().S(`"><button>Magic</button></a></td>
          <td>Runs the magic for all projects</td>
        </tr>
      </tbody>
    </table>
  </div>
  <div class="card">
    <ul class="accordion">
`)
//line views/vgit/Results.html:57
	for _, x := range p.Results {
//line views/vgit/Results.html:57
		qw422016.N().S(`      <li>
        <input id="accordion-`)
//line views/vgit/Results.html:59
		qw422016.E().S(x.Project.Key)
//line views/vgit/Results.html:59
		qw422016.N().S(`" type="checkbox" hidden />
        <label for="accordion-`)
//line views/vgit/Results.html:60
		qw422016.E().S(x.Project.Key)
//line views/vgit/Results.html:60
		qw422016.N().S(`">
          <em class="right">`)
//line views/vgit/Results.html:61
		qw422016.E().S(x.Status)
//line views/vgit/Results.html:61
		if x.Error != "" {
//line views/vgit/Results.html:61
			qw422016.N().S(` (error)`)
//line views/vgit/Results.html:61
		}
//line views/vgit/Results.html:61
		qw422016.N().S(`</em>
          `)
//line views/vgit/Results.html:62
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vgit/Results.html:62
		qw422016.N().S(` `)
//line views/vgit/Results.html:62
		components.StreamSVGRef(qw422016, x.Project.IconSafe(), 16, 16, "icon", ps)
//line views/vgit/Results.html:62
		qw422016.N().S(`
          `)
//line views/vgit/Results.html:63
		qw422016.E().S(x.Project.Title())
//line views/vgit/Results.html:63
		qw422016.N().S(`
`)
//line views/vgit/Results.html:64
		if x.DataString("branch") != "" {
//line views/vgit/Results.html:64
			qw422016.N().S(`          <em>(`)
//line views/vgit/Results.html:65
			qw422016.E().S(x.DataString("branch"))
//line views/vgit/Results.html:65
			qw422016.N().S(`)</em>
`)
//line views/vgit/Results.html:66
		}
//line views/vgit/Results.html:66
		qw422016.N().S(`        </label>
        <div class="bd">
          `)
//line views/vgit/Results.html:69
		streamstatusActions(qw422016, x)
//line views/vgit/Results.html:69
		qw422016.N().S(`
          `)
//line views/vgit/Results.html:70
		streamstatusDetail(qw422016, x)
//line views/vgit/Results.html:70
		qw422016.N().S(`
        </div>
      </li>
`)
//line views/vgit/Results.html:73
	}
//line views/vgit/Results.html:73
	qw422016.N().S(`    </ul>
  </div>
`)
//line views/vgit/Results.html:76
}

//line views/vgit/Results.html:76
func (p *Results) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgit/Results.html:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgit/Results.html:76
	p.StreamBody(qw422016, as, ps)
//line views/vgit/Results.html:76
	qt422016.ReleaseWriter(qw422016)
//line views/vgit/Results.html:76
}

//line views/vgit/Results.html:76
func (p *Results) Body(as *app.State, ps *cutil.PageState) string {
//line views/vgit/Results.html:76
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgit/Results.html:76
	p.WriteBody(qb422016, as, ps)
//line views/vgit/Results.html:76
	qs422016 := string(qb422016.B)
//line views/vgit/Results.html:76
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgit/Results.html:76
	return qs422016
//line views/vgit/Results.html:76
}

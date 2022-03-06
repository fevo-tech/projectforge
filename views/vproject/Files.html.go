// Code generated by qtc from "Files.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vproject/Files.html:1
package vproject

//line views/vproject/Files.html:1
import (
	"path/filepath"
	"projectforge.dev/app"
	"projectforge.dev/app/controller/cutil"
	"projectforge.dev/app/project"
	"projectforge.dev/views/layout"
	"projectforge.dev/views/vfile"
)

//line views/vproject/Files.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vproject/Files.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vproject/Files.html:10
type Files struct {
	layout.Basic
	Project *project.Project
	Path    []string
}

//line views/vproject/Files.html:16
func (p *Files) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/Files.html:16
	qw422016.N().S(`
`)
//line views/vproject/Files.html:18
	prj := p.Project
	fs := as.Services.Projects.GetFilesystem(prj)
	u := "/p/" + prj.Key + "/fs"

//line views/vproject/Files.html:21
	qw422016.N().S(`  `)
//line views/vproject/Files.html:22
	StreamSummary(qw422016, prj, nil, nil, ps)
//line views/vproject/Files.html:22
	qw422016.N().S(`
`)
//line views/vproject/Files.html:23
	if fs.IsDir(filepath.Join(p.Path...)) {
//line views/vproject/Files.html:24
		files := fs.ListFiles(filepath.Join(p.Path...), nil)

//line views/vproject/Files.html:24
		qw422016.N().S(`  `)
//line views/vproject/Files.html:25
		vfile.StreamList(qw422016, p.Path, files, fs, u, as, ps)
//line views/vproject/Files.html:25
		qw422016.N().S(`
`)
//line views/vproject/Files.html:26
	} else {
//line views/vproject/Files.html:28
		b, err := fs.ReadFile(filepath.Join(p.Path...))
		if err != nil {
			panic(err)
		}

//line views/vproject/Files.html:32
		qw422016.N().S(`  `)
//line views/vproject/Files.html:33
		vfile.StreamDetail(qw422016, p.Path, b, u, as, ps)
//line views/vproject/Files.html:33
		qw422016.N().S(`
`)
//line views/vproject/Files.html:34
	}
//line views/vproject/Files.html:35
}

//line views/vproject/Files.html:35
func (p *Files) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/Files.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vproject/Files.html:35
	p.StreamBody(qw422016, as, ps)
//line views/vproject/Files.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/vproject/Files.html:35
}

//line views/vproject/Files.html:35
func (p *Files) Body(as *app.State, ps *cutil.PageState) string {
//line views/vproject/Files.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vproject/Files.html:35
	p.WriteBody(qb422016, as, ps)
//line views/vproject/Files.html:35
	qs422016 := string(qb422016.B)
//line views/vproject/Files.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vproject/Files.html:35
	return qs422016
//line views/vproject/Files.html:35
}
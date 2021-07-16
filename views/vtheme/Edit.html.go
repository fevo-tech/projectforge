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
	"github.com/kyleu/projectforge/views/components"
	"github.com/kyleu/projectforge/views/layout"
)

//line views/vtheme/Edit.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/Edit.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/Edit.html:10
type Edit struct {
	layout.Basic
	Theme *theme.Theme
}

//line views/vtheme/Edit.html:15
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/Edit.html:15
	qw422016.N().S(`
  <form action="" method="post">
    <div class="card">
      <div class="right"><a href="#modal-theme"><button type="button">JSON</button></a></div>
      <h3>`)
//line views/vtheme/Edit.html:19
	if p.Theme.Key == theme.KeyNew {
//line views/vtheme/Edit.html:19
		qw422016.N().S(`New Theme`)
//line views/vtheme/Edit.html:19
	} else {
//line views/vtheme/Edit.html:19
		qw422016.N().S(`Theme Edit`)
//line views/vtheme/Edit.html:19
	}
//line views/vtheme/Edit.html:19
	qw422016.N().S(`</h3>
      <table class="mt">
        <tbody>
          `)
//line views/vtheme/Edit.html:22
	components.StreamTableInput(qw422016, "key", "Key", p.Theme.Key, 5)
//line views/vtheme/Edit.html:22
	qw422016.N().S(`
        </tbody>
      </table>
    </div>
    <div class="card">
      <table class="centered">
        <thead>
        <tr>
          <th></th>
          <th class="border-left" colspan="2">
            <div>Light</div>
            <div id="mockup-light">`)
//line views/vtheme/Edit.html:33
	StreamMockupColors(qw422016, "", p.Theme.Light, 5, ps)
//line views/vtheme/Edit.html:33
	qw422016.N().S(`</div>
          </th>
          <th class="border-left" colspan="2">
            <div>Dark</div>
            <div id="mockup-dark">`)
//line views/vtheme/Edit.html:37
	StreamMockupColors(qw422016, "", p.Theme.Dark, 5, ps)
//line views/vtheme/Edit.html:37
	qw422016.N().S(`</div>
          </th>
        </tr>
        <tr>
          <th></th>
          <th class="border-left">Background</th>
          <th>Foreground</th>
          <th class="border-left">Background</th>
          <th>Foreground</th>
        </tr>
        </thead>
`)
//line views/vtheme/Edit.html:49
	t := p.Theme
	lp := "light"
	dp := "dark"

//line views/vtheme/Edit.html:52
	qw422016.N().S(`        <tbody>
          <tr>
            <th>Main Content</th>
            <td class="border-left">`)
//line views/vtheme/Edit.html:56
	streamcinput(qw422016, lp, "background", t.Light.Background)
//line views/vtheme/Edit.html:56
	qw422016.N().S(`</td>
            <td>`)
//line views/vtheme/Edit.html:57
	streamcinput(qw422016, lp, "foreground", t.Light.Foreground)
//line views/vtheme/Edit.html:57
	qw422016.N().S(`</td>
            <td class="border-left">`)
//line views/vtheme/Edit.html:58
	streamcinput(qw422016, dp, "background", t.Dark.Background)
//line views/vtheme/Edit.html:58
	qw422016.N().S(`</td>
            <td>`)
//line views/vtheme/Edit.html:59
	streamcinput(qw422016, dp, "foreground", t.Dark.Foreground)
//line views/vtheme/Edit.html:59
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Muted</th>
            <td class="border-left">`)
//line views/vtheme/Edit.html:63
	streamcinput(qw422016, lp, "background-muted", t.Light.BackgroundMuted)
//line views/vtheme/Edit.html:63
	qw422016.N().S(`</td>
            <td>`)
//line views/vtheme/Edit.html:64
	streamcinput(qw422016, lp, "foreground-muted", t.Light.ForegroundMuted)
//line views/vtheme/Edit.html:64
	qw422016.N().S(`</td>
            <td class="border-left">`)
//line views/vtheme/Edit.html:65
	streamcinput(qw422016, dp, "background-muted", t.Dark.BackgroundMuted)
//line views/vtheme/Edit.html:65
	qw422016.N().S(`</td>
            <td>`)
//line views/vtheme/Edit.html:66
	streamcinput(qw422016, dp, "foreground-muted", t.Dark.ForegroundMuted)
//line views/vtheme/Edit.html:66
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Link</th>
            <td class="border-left"></td>
            <td>`)
//line views/vtheme/Edit.html:71
	streamcinput(qw422016, lp, "link-foreground", t.Light.LinkForeground)
//line views/vtheme/Edit.html:71
	qw422016.N().S(`</td>
            <td class="border-left"></td>
            <td>`)
//line views/vtheme/Edit.html:73
	streamcinput(qw422016, dp, "link-foreground", t.Dark.LinkForeground)
//line views/vtheme/Edit.html:73
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Visited Link</th>
            <td class="border-left"></td>
            <td>`)
//line views/vtheme/Edit.html:78
	streamcinput(qw422016, lp, "link-visited-foreground", t.Light.LinkVisitedForeground)
//line views/vtheme/Edit.html:78
	qw422016.N().S(`</td>
            <td class="border-left"></td>
            <td>`)
//line views/vtheme/Edit.html:80
	streamcinput(qw422016, dp, "link-visited-foreground", t.Dark.LinkVisitedForeground)
//line views/vtheme/Edit.html:80
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Navigation</th>
            <td class="border-left">`)
//line views/vtheme/Edit.html:84
	streamcinput(qw422016, lp, "nav-background", t.Light.NavBackground)
//line views/vtheme/Edit.html:84
	qw422016.N().S(`</td>
            <td>`)
//line views/vtheme/Edit.html:85
	streamcinput(qw422016, lp, "nav-foreground", t.Light.NavForeground)
//line views/vtheme/Edit.html:85
	qw422016.N().S(`</td>
            <td class="border-left">`)
//line views/vtheme/Edit.html:86
	streamcinput(qw422016, dp, "nav-background", t.Dark.NavBackground)
//line views/vtheme/Edit.html:86
	qw422016.N().S(`</td>
            <td>`)
//line views/vtheme/Edit.html:87
	streamcinput(qw422016, dp, "nav-foreground", t.Dark.NavForeground)
//line views/vtheme/Edit.html:87
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Menu</th>
            <td class="border-left">`)
//line views/vtheme/Edit.html:91
	streamcinput(qw422016, lp, "menu-background", t.Light.MenuBackground)
//line views/vtheme/Edit.html:91
	qw422016.N().S(`</td>
            <td>`)
//line views/vtheme/Edit.html:92
	streamcinput(qw422016, lp, "menu-foreground", t.Light.MenuForeground)
//line views/vtheme/Edit.html:92
	qw422016.N().S(`</td>
            <td class="border-left">`)
//line views/vtheme/Edit.html:93
	streamcinput(qw422016, dp, "menu-background", t.Dark.MenuBackground)
//line views/vtheme/Edit.html:93
	qw422016.N().S(`</td>
            <td>`)
//line views/vtheme/Edit.html:94
	streamcinput(qw422016, dp, "menu-foreground", t.Dark.MenuForeground)
//line views/vtheme/Edit.html:94
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Selected Menu</th>
            <td class="border-left">`)
//line views/vtheme/Edit.html:98
	streamcinput(qw422016, lp, "menu-selected-background", t.Light.MenuSelectedBackground)
//line views/vtheme/Edit.html:98
	qw422016.N().S(`</td>
            <td>`)
//line views/vtheme/Edit.html:99
	streamcinput(qw422016, lp, "menu-selected-foreground", t.Light.MenuSelectedForeground)
//line views/vtheme/Edit.html:99
	qw422016.N().S(`</td>
            <td class="border-left">`)
//line views/vtheme/Edit.html:100
	streamcinput(qw422016, dp, "menu-selected-background", t.Dark.MenuSelectedBackground)
//line views/vtheme/Edit.html:100
	qw422016.N().S(`</td>
            <td>`)
//line views/vtheme/Edit.html:101
	streamcinput(qw422016, dp, "menu-selected-foreground", t.Dark.MenuSelectedForeground)
//line views/vtheme/Edit.html:101
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="card">
      <button type="submit">Save All Changes</button>
      <a href="/theme/`)
//line views/vtheme/Edit.html:108
	qw422016.E().S(t.Key)
//line views/vtheme/Edit.html:108
	qw422016.N().S(`"><button type="button">Reset</button></a>
    </div>
  </form>
  `)
//line views/vtheme/Edit.html:111
	components.StreamJSONModal(qw422016, "theme", "Theme JSON", p.Theme, 1)
//line views/vtheme/Edit.html:111
	qw422016.N().S(`
`)
//line views/vtheme/Edit.html:112
}

//line views/vtheme/Edit.html:112
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/Edit.html:112
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Edit.html:112
	p.StreamBody(qw422016, as, ps)
//line views/vtheme/Edit.html:112
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Edit.html:112
}

//line views/vtheme/Edit.html:112
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vtheme/Edit.html:112
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Edit.html:112
	p.WriteBody(qb422016, as, ps)
//line views/vtheme/Edit.html:112
	qs422016 := string(qb422016.B)
//line views/vtheme/Edit.html:112
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Edit.html:112
	return qs422016
//line views/vtheme/Edit.html:112
}

//line views/vtheme/Edit.html:114
func streamcinput(qw422016 *qt422016.Writer, mode string, k string, v string) {
//line views/vtheme/Edit.html:114
	qw422016.N().S(`<input class="color-var" data-mode="`)
//line views/vtheme/Edit.html:115
	qw422016.E().S(mode)
//line views/vtheme/Edit.html:115
	qw422016.N().S(`" data-var="color-`)
//line views/vtheme/Edit.html:115
	qw422016.E().S(k)
//line views/vtheme/Edit.html:115
	qw422016.N().S(`" type="color" id="`)
//line views/vtheme/Edit.html:115
	qw422016.E().S(mode)
//line views/vtheme/Edit.html:115
	qw422016.N().S(`-`)
//line views/vtheme/Edit.html:115
	qw422016.E().S(k)
//line views/vtheme/Edit.html:115
	qw422016.N().S(`" name="`)
//line views/vtheme/Edit.html:115
	qw422016.E().S(mode)
//line views/vtheme/Edit.html:115
	qw422016.N().S(`-`)
//line views/vtheme/Edit.html:115
	qw422016.E().S(k)
//line views/vtheme/Edit.html:115
	qw422016.N().S(`" value="`)
//line views/vtheme/Edit.html:115
	qw422016.E().S(v)
//line views/vtheme/Edit.html:115
	qw422016.N().S(`" autocomplete="off" />`)
//line views/vtheme/Edit.html:116
}

//line views/vtheme/Edit.html:116
func writecinput(qq422016 qtio422016.Writer, mode string, k string, v string) {
//line views/vtheme/Edit.html:116
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Edit.html:116
	streamcinput(qw422016, mode, k, v)
//line views/vtheme/Edit.html:116
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Edit.html:116
}

//line views/vtheme/Edit.html:116
func cinput(mode string, k string, v string) string {
//line views/vtheme/Edit.html:116
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Edit.html:116
	writecinput(qb422016, mode, k, v)
//line views/vtheme/Edit.html:116
	qs422016 := string(qb422016.B)
//line views/vtheme/Edit.html:116
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Edit.html:116
	return qs422016
//line views/vtheme/Edit.html:116
}

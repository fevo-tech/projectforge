// Code generated by qtc from "Mockup.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vtheme/Mockup.html:1
package vtheme

//line views/vtheme/Mockup.html:1
import (
	"fmt"
	"github.com/kyleu/projectforge/app/controller/cutil"
	"github.com/kyleu/projectforge/app/theme"
	"github.com/kyleu/projectforge/app/util"
	"github.com/kyleu/projectforge/views/components"
	"github.com/kyleu/projectforge/views/vutil"
)

//line views/vtheme/Mockup.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/Mockup.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/Mockup.html:10
func StreamMockupTheme(qw422016 *qt422016.Writer, t *theme.Theme, indent int, ps *cutil.PageState) {
//line views/vtheme/Mockup.html:11
	StreamMockupColors(qw422016, "light", t.Light, indent, ps)
//line views/vtheme/Mockup.html:12
	StreamMockupColors(qw422016, "dark", t.Dark, indent, ps)
//line views/vtheme/Mockup.html:12
	qw422016.N().S(`<div class="title">`)
//line views/vtheme/Mockup.html:13
	qw422016.E().S(t.Key)
//line views/vtheme/Mockup.html:13
	qw422016.N().S(`</div>`)
//line views/vtheme/Mockup.html:14
}

//line views/vtheme/Mockup.html:14
func WriteMockupTheme(qq422016 qtio422016.Writer, t *theme.Theme, indent int, ps *cutil.PageState) {
//line views/vtheme/Mockup.html:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Mockup.html:14
	StreamMockupTheme(qw422016, t, indent, ps)
//line views/vtheme/Mockup.html:14
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Mockup.html:14
}

//line views/vtheme/Mockup.html:14
func MockupTheme(t *theme.Theme, indent int, ps *cutil.PageState) string {
//line views/vtheme/Mockup.html:14
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Mockup.html:14
	WriteMockupTheme(qb422016, t, indent, ps)
//line views/vtheme/Mockup.html:14
	qs422016 := string(qb422016.B)
//line views/vtheme/Mockup.html:14
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Mockup.html:14
	return qs422016
//line views/vtheme/Mockup.html:14
}

//line views/vtheme/Mockup.html:16
func StreamMockupColors(qw422016 *qt422016.Writer, mode string, c *theme.Colors, indent int, ps *cutil.PageState) {
//line views/vtheme/Mockup.html:18
	navStyle := fmt.Sprintf("color: %s; background-color: %s;", c.NavForeground, c.NavBackground)

	menuStyle := fmt.Sprintf("color: %s; background-color: %s;", c.MenuForeground, c.MenuBackground)
	menuLinkStyle := fmt.Sprintf("color: %s;", c.MenuForeground)
	menuLinkSelectedStyle := fmt.Sprintf("color: %s; background-color: %s;", c.MenuSelectedForeground, c.MenuSelectedBackground)

	mainStyle := fmt.Sprintf("color: %s; background-color: %s;", c.Foreground, c.Background)
	mutedStyle := fmt.Sprintf("color: %s; background-color: %s;", c.ForegroundMuted, c.BackgroundMuted)
	linkStyle := fmt.Sprintf("color: %s;", c.LinkForeground)
	linkVisitedStyle := fmt.Sprintf("color: %s;", c.LinkVisitedForeground)

	cls := "mockup"
	if mode != "" {
		cls += " only-" + mode
	}

//line views/vtheme/Mockup.html:34
	vutil.StreamIndent(qw422016, true, indent)
//line views/vtheme/Mockup.html:34
	qw422016.N().S(`<div class="`)
//line views/vtheme/Mockup.html:35
	qw422016.E().S(cls)
//line views/vtheme/Mockup.html:35
	qw422016.N().S(`"><div class="mock-nav" style="`)
//line views/vtheme/Mockup.html:36
	qw422016.E().S(navStyle)
//line views/vtheme/Mockup.html:36
	qw422016.N().S(`">`)
//line views/vtheme/Mockup.html:36
	components.StreamSVGRef(qw422016, `app`, 12, 12, `icon`, ps)
//line views/vtheme/Mockup.html:36
	qw422016.E().S(util.AppName)
//line views/vtheme/Mockup.html:36
	qw422016.N().S(`</div><div class="mock-menu" style="`)
//line views/vtheme/Mockup.html:37
	qw422016.E().S(menuStyle)
//line views/vtheme/Mockup.html:37
	qw422016.N().S(`"><div class="mock-link" style="`)
//line views/vtheme/Mockup.html:38
	qw422016.E().S(menuLinkStyle)
//line views/vtheme/Mockup.html:38
	qw422016.N().S(`">A</div><div class="mock-link-selected" style="`)
//line views/vtheme/Mockup.html:39
	qw422016.E().S(menuLinkSelectedStyle)
//line views/vtheme/Mockup.html:39
	qw422016.N().S(`">B</div><div class="mock-link" style="`)
//line views/vtheme/Mockup.html:40
	qw422016.E().S(menuLinkStyle)
//line views/vtheme/Mockup.html:40
	qw422016.N().S(`">C</div><div class="mock-link" style="`)
//line views/vtheme/Mockup.html:41
	qw422016.E().S(menuLinkStyle)
//line views/vtheme/Mockup.html:41
	qw422016.N().S(`">D</div></div><div class="mock-main" style="`)
//line views/vtheme/Mockup.html:43
	qw422016.E().S(mainStyle)
//line views/vtheme/Mockup.html:43
	qw422016.N().S(`"><div>Welcome!</div><div><div class="mock-muted" style="`)
//line views/vtheme/Mockup.html:46
	qw422016.E().S(mutedStyle)
//line views/vtheme/Mockup.html:46
	qw422016.N().S(`">Here's some links:</div><ul><li class="mock-link" style="`)
//line views/vtheme/Mockup.html:48
	qw422016.E().S(linkStyle)
//line views/vtheme/Mockup.html:48
	qw422016.N().S(`">New</li><li class="mock-link" style="`)
//line views/vtheme/Mockup.html:49
	qw422016.E().S(linkStyle)
//line views/vtheme/Mockup.html:49
	qw422016.N().S(`">Also New</li><li class="mock-link-visited" style="`)
//line views/vtheme/Mockup.html:50
	qw422016.E().S(linkVisitedStyle)
//line views/vtheme/Mockup.html:50
	qw422016.N().S(`">Visited</li></ul></div></div>`)
//line views/vtheme/Mockup.html:54
	vutil.StreamIndent(qw422016, true, indent)
//line views/vtheme/Mockup.html:54
	qw422016.N().S(`</div>`)
//line views/vtheme/Mockup.html:56
}

//line views/vtheme/Mockup.html:56
func WriteMockupColors(qq422016 qtio422016.Writer, mode string, c *theme.Colors, indent int, ps *cutil.PageState) {
//line views/vtheme/Mockup.html:56
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Mockup.html:56
	StreamMockupColors(qw422016, mode, c, indent, ps)
//line views/vtheme/Mockup.html:56
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Mockup.html:56
}

//line views/vtheme/Mockup.html:56
func MockupColors(mode string, c *theme.Colors, indent int, ps *cutil.PageState) string {
//line views/vtheme/Mockup.html:56
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Mockup.html:56
	WriteMockupColors(qb422016, mode, c, indent, ps)
//line views/vtheme/Mockup.html:56
	qs422016 := string(qb422016.B)
//line views/vtheme/Mockup.html:56
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Mockup.html:56
	return qs422016
//line views/vtheme/Mockup.html:56
}

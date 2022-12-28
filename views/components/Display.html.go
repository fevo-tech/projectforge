// Code generated by qtc from "Display.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/Display.html:2
package components

//line views/components/Display.html:2
import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/lib/filter"
	"projectforge.dev/projectforge/app/util"
)

//line views/components/Display.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Display.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Display.html:14
func StreamDisplayTimestamp(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:14
	qw422016.N().S(`<span class="nowrap">`)
//line views/components/Display.html:15
	qw422016.E().S(util.TimeToFull(value))
//line views/components/Display.html:15
	qw422016.N().S(`</span>`)
//line views/components/Display.html:16
}

//line views/components/Display.html:16
func WriteDisplayTimestamp(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:16
	StreamDisplayTimestamp(qw422016, value)
//line views/components/Display.html:16
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:16
}

//line views/components/Display.html:16
func DisplayTimestamp(value *time.Time) string {
//line views/components/Display.html:16
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:16
	WriteDisplayTimestamp(qb422016, value)
//line views/components/Display.html:16
	qs422016 := string(qb422016.B)
//line views/components/Display.html:16
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:16
	return qs422016
//line views/components/Display.html:16
}

//line views/components/Display.html:18
func StreamDisplayTimestampDay(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:19
	qw422016.E().S(util.TimeToYMD(value))
//line views/components/Display.html:20
}

//line views/components/Display.html:20
func WriteDisplayTimestampDay(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:20
	StreamDisplayTimestampDay(qw422016, value)
//line views/components/Display.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:20
}

//line views/components/Display.html:20
func DisplayTimestampDay(value *time.Time) string {
//line views/components/Display.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:20
	WriteDisplayTimestampDay(qb422016, value)
//line views/components/Display.html:20
	qs422016 := string(qb422016.B)
//line views/components/Display.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:20
	return qs422016
//line views/components/Display.html:20
}

//line views/components/Display.html:22
func StreamDisplayUUID(qw422016 *qt422016.Writer, value *uuid.UUID) {
//line views/components/Display.html:23
	if value != nil {
//line views/components/Display.html:24
		qw422016.E().S(value.String())
//line views/components/Display.html:25
	}
//line views/components/Display.html:26
}

//line views/components/Display.html:26
func WriteDisplayUUID(qq422016 qtio422016.Writer, value *uuid.UUID) {
//line views/components/Display.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:26
	StreamDisplayUUID(qw422016, value)
//line views/components/Display.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:26
}

//line views/components/Display.html:26
func DisplayUUID(value *uuid.UUID) string {
//line views/components/Display.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:26
	WriteDisplayUUID(qb422016, value)
//line views/components/Display.html:26
	qs422016 := string(qb422016.B)
//line views/components/Display.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:26
	return qs422016
//line views/components/Display.html:26
}

//line views/components/Display.html:28
func StreamDisplayStringArray(qw422016 *qt422016.Writer, value []string) {
//line views/components/Display.html:29
	if len(value) == 0 {
//line views/components/Display.html:29
		qw422016.N().S(`<em>empty</em>`)
//line views/components/Display.html:31
	}
//line views/components/Display.html:33
	max := 3
	display := value
	var extra int
	if len(value) > max {
		extra = len(value) - max
		display = display[:max]
	}

//line views/components/Display.html:41
	if extra > 0 {
//line views/components/Display.html:41
		qw422016.N().S(`<span title="`)
//line views/components/Display.html:41
		qw422016.E().S(strings.Join(value, `, `))
//line views/components/Display.html:41
		qw422016.N().S(`">`)
//line views/components/Display.html:41
	}
//line views/components/Display.html:42
	for idx, v := range display {
//line views/components/Display.html:43
		if idx > 0 {
//line views/components/Display.html:43
			qw422016.N().S(`,`)
//line views/components/Display.html:43
			qw422016.N().S(` `)
//line views/components/Display.html:43
		}
//line views/components/Display.html:44
		qw422016.E().S(v)
//line views/components/Display.html:45
	}
//line views/components/Display.html:46
	if extra > 0 {
//line views/components/Display.html:46
		qw422016.N().S(`, <em>and`)
//line views/components/Display.html:46
		qw422016.N().S(` `)
//line views/components/Display.html:46
		qw422016.N().D(extra)
//line views/components/Display.html:46
		qw422016.N().S(` `)
//line views/components/Display.html:46
		qw422016.N().S(`more...</em>`)
//line views/components/Display.html:46
	}
//line views/components/Display.html:47
	if extra > 0 {
//line views/components/Display.html:47
		qw422016.N().S(`</span>`)
//line views/components/Display.html:47
	}
//line views/components/Display.html:48
}

//line views/components/Display.html:48
func WriteDisplayStringArray(qq422016 qtio422016.Writer, value []string) {
//line views/components/Display.html:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:48
	StreamDisplayStringArray(qw422016, value)
//line views/components/Display.html:48
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:48
}

//line views/components/Display.html:48
func DisplayStringArray(value []string) string {
//line views/components/Display.html:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:48
	WriteDisplayStringArray(qb422016, value)
//line views/components/Display.html:48
	qs422016 := string(qb422016.B)
//line views/components/Display.html:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:48
	return qs422016
//line views/components/Display.html:48
}

//line views/components/Display.html:50
func StreamDisplayDiffs(qw422016 *qt422016.Writer, value util.Diffs) {
//line views/components/Display.html:51
	if len(value) == 0 {
//line views/components/Display.html:51
		qw422016.N().S(`<em>no changes</em>`)
//line views/components/Display.html:53
	} else {
//line views/components/Display.html:53
		qw422016.N().S(`<table><tbody>`)
//line views/components/Display.html:56
		for _, d := range value {
//line views/components/Display.html:56
			qw422016.N().S(`<tr><td><code>`)
//line views/components/Display.html:58
			qw422016.E().S(d.Path)
//line views/components/Display.html:58
			qw422016.N().S(`</code></td><td><code class="error">`)
//line views/components/Display.html:59
			qw422016.E().S(d.Old)
//line views/components/Display.html:59
			qw422016.N().S(`</code></td><td>→</td><td><code class="success">`)
//line views/components/Display.html:61
			qw422016.E().S(d.New)
//line views/components/Display.html:61
			qw422016.N().S(`</code></td></tr>`)
//line views/components/Display.html:63
		}
//line views/components/Display.html:63
		qw422016.N().S(`</tbody></table>`)
//line views/components/Display.html:66
	}
//line views/components/Display.html:67
}

//line views/components/Display.html:67
func WriteDisplayDiffs(qq422016 qtio422016.Writer, value util.Diffs) {
//line views/components/Display.html:67
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:67
	StreamDisplayDiffs(qw422016, value)
//line views/components/Display.html:67
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:67
}

//line views/components/Display.html:67
func DisplayDiffs(value util.Diffs) string {
//line views/components/Display.html:67
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:67
	WriteDisplayDiffs(qb422016, value)
//line views/components/Display.html:67
	qs422016 := string(qb422016.B)
//line views/components/Display.html:67
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:67
	return qs422016
//line views/components/Display.html:67
}

//line views/components/Display.html:69
func StreamDisplayMaps(qw422016 *qt422016.Writer, maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) {
//line views/components/Display.html:70
	if len(maps) == 0 {
//line views/components/Display.html:70
		qw422016.N().S(`<em>no results</em>`)
//line views/components/Display.html:72
	} else {
//line views/components/Display.html:72
		qw422016.N().S(`<table><thead><tr>`)
//line views/components/Display.html:76
		for _, k := range maps[0].Keys() {
//line views/components/Display.html:77
			StreamTableHeaderSimple(qw422016, "map", k, k, "", params, nil, ps)
//line views/components/Display.html:78
		}
//line views/components/Display.html:78
		qw422016.N().S(`</tr></thead><tbody>`)
//line views/components/Display.html:82
		for _, m := range maps {
//line views/components/Display.html:82
			qw422016.N().S(`<tr>`)
//line views/components/Display.html:84
			for _, k := range m.Keys() {
//line views/components/Display.html:86
				res := ""
				switch t := m[k].(type) {
				case string:
					res = t
				case []byte:
					res = string(t)
				default:
					res = fmt.Sprint(m[k])
				}

//line views/components/Display.html:96
				if preserveWhitespace {
//line views/components/Display.html:96
					qw422016.N().S(`<td style="white-space: pre;">`)
//line views/components/Display.html:97
					qw422016.E().S(res)
//line views/components/Display.html:97
					qw422016.N().S(`</td>`)
//line views/components/Display.html:98
				} else {
//line views/components/Display.html:98
					qw422016.N().S(`<td>`)
//line views/components/Display.html:99
					qw422016.E().S(res)
//line views/components/Display.html:99
					qw422016.N().S(`</td>`)
//line views/components/Display.html:100
				}
//line views/components/Display.html:101
			}
//line views/components/Display.html:101
			qw422016.N().S(`</tr>`)
//line views/components/Display.html:103
		}
//line views/components/Display.html:103
		qw422016.N().S(`</tbody></table>`)
//line views/components/Display.html:106
	}
//line views/components/Display.html:107
}

//line views/components/Display.html:107
func WriteDisplayMaps(qq422016 qtio422016.Writer, maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) {
//line views/components/Display.html:107
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:107
	StreamDisplayMaps(qw422016, maps, params, preserveWhitespace, ps)
//line views/components/Display.html:107
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:107
}

//line views/components/Display.html:107
func DisplayMaps(maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) string {
//line views/components/Display.html:107
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:107
	WriteDisplayMaps(qb422016, maps, params, preserveWhitespace, ps)
//line views/components/Display.html:107
	qs422016 := string(qb422016.B)
//line views/components/Display.html:107
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:107
	return qs422016
//line views/components/Display.html:107
}

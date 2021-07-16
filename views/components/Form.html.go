// Code generated by qtc from "Form.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/Form.html:1
package components

//line views/components/Form.html:1
import (
	"github.com/kyleu/projectforge/app/util"
	"github.com/kyleu/projectforge/views/vutil"
)

//line views/components/Form.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Form.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Form.html:6
func StreamFormInput(qw422016 *qt422016.Writer, key string, id string, value string) {
//line views/components/Form.html:7
	if id == "" {
//line views/components/Form.html:7
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:8
		qw422016.E().S(key)
//line views/components/Form.html:8
		qw422016.N().S(`" value="`)
//line views/components/Form.html:8
		qw422016.E().S(value)
//line views/components/Form.html:8
		qw422016.N().S(`" />`)
//line views/components/Form.html:9
	} else {
//line views/components/Form.html:9
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:10
		qw422016.E().S(id)
//line views/components/Form.html:10
		qw422016.N().S(`" name="`)
//line views/components/Form.html:10
		qw422016.E().S(key)
//line views/components/Form.html:10
		qw422016.N().S(`" value="`)
//line views/components/Form.html:10
		qw422016.E().S(value)
//line views/components/Form.html:10
		qw422016.N().S(`" />`)
//line views/components/Form.html:11
	}
//line views/components/Form.html:12
}

//line views/components/Form.html:12
func WriteFormInput(qq422016 qtio422016.Writer, key string, id string, value string) {
//line views/components/Form.html:12
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:12
	StreamFormInput(qw422016, key, id, value)
//line views/components/Form.html:12
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:12
}

//line views/components/Form.html:12
func FormInput(key string, id string, value string) string {
//line views/components/Form.html:12
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:12
	WriteFormInput(qb422016, key, id, value)
//line views/components/Form.html:12
	qs422016 := string(qb422016.B)
//line views/components/Form.html:12
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:12
	return qs422016
//line views/components/Form.html:12
}

//line views/components/Form.html:14
func StreamFormInputNumber(qw422016 *qt422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:15
	if id == "" {
//line views/components/Form.html:15
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:16
		qw422016.E().S(key)
//line views/components/Form.html:16
		qw422016.N().S(`" type="number" value="`)
//line views/components/Form.html:16
		qw422016.E().V(value)
//line views/components/Form.html:16
		qw422016.N().S(`" />`)
//line views/components/Form.html:17
	} else {
//line views/components/Form.html:17
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:18
		qw422016.E().S(id)
//line views/components/Form.html:18
		qw422016.N().S(`" name="`)
//line views/components/Form.html:18
		qw422016.E().S(key)
//line views/components/Form.html:18
		qw422016.N().S(`" type="number" value="`)
//line views/components/Form.html:18
		qw422016.E().V(value)
//line views/components/Form.html:18
		qw422016.N().S(`" />`)
//line views/components/Form.html:19
	}
//line views/components/Form.html:20
}

//line views/components/Form.html:20
func WriteFormInputNumber(qq422016 qtio422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:20
	StreamFormInputNumber(qw422016, key, id, value)
//line views/components/Form.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:20
}

//line views/components/Form.html:20
func FormInputNumber(key string, id string, value interface{}) string {
//line views/components/Form.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:20
	WriteFormInputNumber(qb422016, key, id, value)
//line views/components/Form.html:20
	qs422016 := string(qb422016.B)
//line views/components/Form.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:20
	return qs422016
//line views/components/Form.html:20
}

//line views/components/Form.html:22
func StreamFormInputTimestamp(qw422016 *qt422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:23
	if id == "" {
//line views/components/Form.html:23
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:24
		qw422016.E().S(key)
//line views/components/Form.html:24
		qw422016.N().S(`" type="datetime-local" value="`)
//line views/components/Form.html:24
		qw422016.E().V(value)
//line views/components/Form.html:24
		qw422016.N().S(`" />`)
//line views/components/Form.html:25
	} else {
//line views/components/Form.html:25
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:26
		qw422016.E().S(id)
//line views/components/Form.html:26
		qw422016.N().S(`" name="`)
//line views/components/Form.html:26
		qw422016.E().S(key)
//line views/components/Form.html:26
		qw422016.N().S(`" type="datetime-local" value="`)
//line views/components/Form.html:26
		qw422016.E().V(value)
//line views/components/Form.html:26
		qw422016.N().S(`" />`)
//line views/components/Form.html:27
	}
//line views/components/Form.html:28
}

//line views/components/Form.html:28
func WriteFormInputTimestamp(qq422016 qtio422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:28
	StreamFormInputTimestamp(qw422016, key, id, value)
//line views/components/Form.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:28
}

//line views/components/Form.html:28
func FormInputTimestamp(key string, id string, value interface{}) string {
//line views/components/Form.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:28
	WriteFormInputTimestamp(qb422016, key, id, value)
//line views/components/Form.html:28
	qs422016 := string(qb422016.B)
//line views/components/Form.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:28
	return qs422016
//line views/components/Form.html:28
}

//line views/components/Form.html:30
func StreamFormTextarea(qw422016 *qt422016.Writer, key string, id string, value string) {
//line views/components/Form.html:31
	if id == "" {
//line views/components/Form.html:31
		qw422016.N().S(`<textarea name="`)
//line views/components/Form.html:32
		qw422016.E().S(key)
//line views/components/Form.html:32
		qw422016.N().S(`">`)
//line views/components/Form.html:32
		qw422016.E().S(value)
//line views/components/Form.html:32
		qw422016.N().S(`</textarea>`)
//line views/components/Form.html:33
	} else {
//line views/components/Form.html:33
		qw422016.N().S(`<textarea id="`)
//line views/components/Form.html:34
		qw422016.E().S(id)
//line views/components/Form.html:34
		qw422016.N().S(`" name="`)
//line views/components/Form.html:34
		qw422016.E().S(key)
//line views/components/Form.html:34
		qw422016.N().S(`">`)
//line views/components/Form.html:34
		qw422016.E().S(value)
//line views/components/Form.html:34
		qw422016.N().S(`</textarea>`)
//line views/components/Form.html:35
	}
//line views/components/Form.html:36
}

//line views/components/Form.html:36
func WriteFormTextarea(qq422016 qtio422016.Writer, key string, id string, value string) {
//line views/components/Form.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:36
	StreamFormTextarea(qw422016, key, id, value)
//line views/components/Form.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:36
}

//line views/components/Form.html:36
func FormTextarea(key string, id string, value string) string {
//line views/components/Form.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:36
	WriteFormTextarea(qb422016, key, id, value)
//line views/components/Form.html:36
	qs422016 := string(qb422016.B)
//line views/components/Form.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:36
	return qs422016
//line views/components/Form.html:36
}

//line views/components/Form.html:38
func StreamFormSelect(qw422016 *qt422016.Writer, key string, id string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:39
	if id == "" {
//line views/components/Form.html:39
		qw422016.N().S(`<select name="`)
//line views/components/Form.html:40
		qw422016.E().S(key)
//line views/components/Form.html:40
		qw422016.N().S(`">`)
//line views/components/Form.html:41
	} else {
//line views/components/Form.html:41
		qw422016.N().S(`<select name="`)
//line views/components/Form.html:42
		qw422016.E().S(key)
//line views/components/Form.html:42
		qw422016.N().S(`" id="`)
//line views/components/Form.html:42
		qw422016.E().S(id)
//line views/components/Form.html:42
		qw422016.N().S(`">`)
//line views/components/Form.html:43
	}
//line views/components/Form.html:44
	for idx, opt := range opts {
//line views/components/Form.html:46
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:51
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Form.html:52
		if opt == value {
//line views/components/Form.html:52
			qw422016.N().S(`<option selected="selected" value="`)
//line views/components/Form.html:53
			qw422016.E().S(opt)
//line views/components/Form.html:53
			qw422016.N().S(`">`)
//line views/components/Form.html:53
			qw422016.E().S(title)
//line views/components/Form.html:53
			qw422016.N().S(`</option>`)
//line views/components/Form.html:54
		} else {
//line views/components/Form.html:54
			qw422016.N().S(`<option value="`)
//line views/components/Form.html:55
			qw422016.E().S(opt)
//line views/components/Form.html:55
			qw422016.N().S(`">`)
//line views/components/Form.html:55
			qw422016.E().S(title)
//line views/components/Form.html:55
			qw422016.N().S(`</option>`)
//line views/components/Form.html:56
		}
//line views/components/Form.html:57
	}
//line views/components/Form.html:58
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:58
	qw422016.N().S(`</select>`)
//line views/components/Form.html:60
}

//line views/components/Form.html:60
func WriteFormSelect(qq422016 qtio422016.Writer, key string, id string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:60
	StreamFormSelect(qw422016, key, id, value, opts, titles, indent)
//line views/components/Form.html:60
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:60
}

//line views/components/Form.html:60
func FormSelect(key string, id string, value string, opts []string, titles []string, indent int) string {
//line views/components/Form.html:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:60
	WriteFormSelect(qb422016, key, id, value, opts, titles, indent)
//line views/components/Form.html:60
	qs422016 := string(qb422016.B)
//line views/components/Form.html:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:60
	return qs422016
//line views/components/Form.html:60
}

//line views/components/Form.html:62
func StreamFormRadio(qw422016 *qt422016.Writer, key string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:63
	for idx, opt := range opts {
//line views/components/Form.html:65
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:70
		vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:71
		if opt == value {
//line views/components/Form.html:71
			qw422016.N().S(`<label><input type="radio" name="`)
//line views/components/Form.html:72
			qw422016.E().S(key)
//line views/components/Form.html:72
			qw422016.N().S(`" value="`)
//line views/components/Form.html:72
			qw422016.E().S(opt)
//line views/components/Form.html:72
			qw422016.N().S(`" checked="checked" />`)
//line views/components/Form.html:72
			qw422016.N().S(` `)
//line views/components/Form.html:72
			qw422016.E().S(title)
//line views/components/Form.html:72
			qw422016.N().S(`</label>`)
//line views/components/Form.html:73
		} else {
//line views/components/Form.html:73
			qw422016.N().S(`<label><input type="radio" name="`)
//line views/components/Form.html:74
			qw422016.E().S(key)
//line views/components/Form.html:74
			qw422016.N().S(`" value="`)
//line views/components/Form.html:74
			qw422016.E().S(opt)
//line views/components/Form.html:74
			qw422016.N().S(`" />`)
//line views/components/Form.html:74
			qw422016.N().S(` `)
//line views/components/Form.html:74
			qw422016.E().S(title)
//line views/components/Form.html:74
			qw422016.N().S(`</label>`)
//line views/components/Form.html:75
		}
//line views/components/Form.html:76
	}
//line views/components/Form.html:77
}

//line views/components/Form.html:77
func WriteFormRadio(qq422016 qtio422016.Writer, key string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:77
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:77
	StreamFormRadio(qw422016, key, value, opts, titles, indent)
//line views/components/Form.html:77
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:77
}

//line views/components/Form.html:77
func FormRadio(key string, value string, opts []string, titles []string, indent int) string {
//line views/components/Form.html:77
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:77
	WriteFormRadio(qb422016, key, value, opts, titles, indent)
//line views/components/Form.html:77
	qs422016 := string(qb422016.B)
//line views/components/Form.html:77
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:77
	return qs422016
//line views/components/Form.html:77
}

//line views/components/Form.html:79
func StreamFormCheckbox(qw422016 *qt422016.Writer, key string, values []string, opts []string, titles []string, indent int) {
//line views/components/Form.html:80
	for idx, opt := range opts {
//line views/components/Form.html:82
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:87
		vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:88
		if util.StringArrayContains(values, opt) {
//line views/components/Form.html:88
			qw422016.N().S(`<label><input type="checkbox" name="`)
//line views/components/Form.html:89
			qw422016.E().S(key)
//line views/components/Form.html:89
			qw422016.N().S(`" value="`)
//line views/components/Form.html:89
			qw422016.E().S(opt)
//line views/components/Form.html:89
			qw422016.N().S(`" checked="checked" />`)
//line views/components/Form.html:89
			qw422016.N().S(` `)
//line views/components/Form.html:89
			qw422016.E().S(title)
//line views/components/Form.html:89
			qw422016.N().S(`</label>`)
//line views/components/Form.html:90
		} else {
//line views/components/Form.html:90
			qw422016.N().S(`<label><input type="checkbox" name="`)
//line views/components/Form.html:91
			qw422016.E().S(key)
//line views/components/Form.html:91
			qw422016.N().S(`" value="`)
//line views/components/Form.html:91
			qw422016.E().S(opt)
//line views/components/Form.html:91
			qw422016.N().S(`" />`)
//line views/components/Form.html:91
			qw422016.N().S(` `)
//line views/components/Form.html:91
			qw422016.E().S(title)
//line views/components/Form.html:91
			qw422016.N().S(`</label>`)
//line views/components/Form.html:92
		}
//line views/components/Form.html:93
	}
//line views/components/Form.html:94
}

//line views/components/Form.html:94
func WriteFormCheckbox(qq422016 qtio422016.Writer, key string, values []string, opts []string, titles []string, indent int) {
//line views/components/Form.html:94
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:94
	StreamFormCheckbox(qw422016, key, values, opts, titles, indent)
//line views/components/Form.html:94
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:94
}

//line views/components/Form.html:94
func FormCheckbox(key string, values []string, opts []string, titles []string, indent int) string {
//line views/components/Form.html:94
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:94
	WriteFormCheckbox(qb422016, key, values, opts, titles, indent)
//line views/components/Form.html:94
	qs422016 := string(qb422016.B)
//line views/components/Form.html:94
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:94
	return qs422016
//line views/components/Form.html:94
}

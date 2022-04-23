// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"context"
	"io/fs"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"golang.org/x/exp/slices"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/lib/menu"
	"projectforge.dev/projectforge/app/util"
	"projectforge.dev/projectforge/doc"
	"projectforge.dev/projectforge/views/vdoc"
)

func Docs(rc *fasthttp.RequestCtx) {
	act("docs", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		pth, _ := RCRequiredString(rc, "path", false)
		if pth == "" {
			return "", errors.New("invalid path")
		}

		split := util.StringSplitAndTrim(pth, "/")
		bc := []string{"docs"}
		bc = append(bc, split...)

		x, err := doc.HTML(pth + ".md")
		if err != nil {
			return "", errors.Wrapf(err, "unable to load documentation from [%s]", pth)
		}
		return render(rc, as, &vdoc.MarkdownPage{Title: pth, HTML: x}, ps, bc...)
	})
}

var docMenuCached *menu.Item

func docMenu(ctx context.Context, as *app.State) *menu.Item {
	if docMenuCached == nil {
		docMenuCached = docMenuCreate(ctx, as)
	}
	return docMenuCached
}

func docMenuCreate(ctx context.Context, as *app.State) *menu.Item {
	var paths []string
	err := fs.WalkDir(doc.FS, ".", func(path string, d fs.DirEntry, err error) error {
		paths = append(paths, path)
		return err
	})
	if err != nil {
		as.Logger.Errorf("unable to build documentation menu: %+v", err)
	}
	slices.Sort(paths)

	ret := &menu.Item{Key: "docs", Title: "Documentation", Icon: "folder"}
	for _, p := range paths {
		if p == "." || !strings.HasSuffix(p, ".md") {
			continue
		}
		split := strings.Split(p, "/")
		p = strings.TrimSuffix(p, ".md")
		mi := ret
		for idx, comp := range split {
			name := strings.TrimSuffix(comp, ".md")
			addFolder := func() {
				i := &menu.Item{Key: name, Title: util.StringToCamel(name), Icon: "folder"}
				mi.Children = append(mi.Children, i)
				slices.SortFunc(mi.Children, func(l *menu.Item, r *menu.Item) bool {
					return l.Title < r.Title
				})
				mi = i
			}
			if idx == len(split)-1 {
				if strings.HasSuffix(comp, ".md") {
					r := "/" + path.Join("docs", p)
					title := util.StringToCamel(name)
					b, err := doc.FS.ReadFile(p + ".md")
					if err != nil {
						panic(err)
					}

					lines := strings.Split(string(b), "\n")
					for _, line := range lines {
						if strings.HasPrefix(line, "#") {
							for strings.HasPrefix(line, "#") {
								line = line[1:]
							}
							title = strings.TrimSpace(line)
							break
						}
					}

					mi.Children = append(mi.Children, &menu.Item{Key: name, Title: title, Icon: "file", Route: r})
				} else {
					addFolder()
				}
			} else {
				if kid := mi.Children.Get(comp); kid == nil {
					addFolder()
				} else {
					mi = kid
				}
			}
		}
	}
	slices.SortFunc(ret.Children, func(l *menu.Item, r *menu.Item) bool {
		return l.Title < r.Title
	})

	return ret
}

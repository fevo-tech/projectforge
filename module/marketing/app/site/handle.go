package site

import (
	"github.com/valyala/fasthttp"

	"{{{ .Package }}}/app"
	"{{{ .Package }}}/app/controller/cutil"
	"{{{ .Package }}}/app/site/download"
	"{{{ .Package }}}/app/util"
	"{{{ .Package }}}/views"
	"{{{ .Package }}}/views/layout"
	"{{{ .Package }}}/views/vsite"
)

func siteData(result string, kvs ...string) map[string]interface{} {
	ret := map[string]interface{}{"app": util.AppName, "result": result}
	for i := 0; i < len(kvs); i += 2 {
		ret[kvs[i]] = kvs[i+1]
	}
	return ret
}

func Handle(path []string, ctx *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (string, layout.Page, []string, error) {
	if len(path) == 0 {
		ps.Data = siteData("Welcome to the marketing site!")
		return "", &vsite.Index{}, path, nil
	}
	var page layout.Page
	switch path[0] {
	case keyIntro:
		ps.Data = siteData("This static page is an introduction to " + util.AppName)
		page = &views.Debug{}
	case keyDownload:
		dls := download.DownloadLinks(as.BuildInfo.Version)
		ps.Data = dls
		page = &vsite.Download{Links: dls}
	case keyInstall:
		ps.Data = siteData("This static page contains installation instructions")
		page = &vsite.Installation{}
	case keyQuickStart:
		ps.Data = siteData("This static page show how to get started with " + util.AppName)
		page = &views.Debug{}
	case keyContrib:
		ps.Data = siteData("This static page describes how to build " + util.AppName)
		page = &views.Debug{}
	default:
		ps.Data = "TODO!"
		page = &views.Debug{}
	}
	return "", page, path, nil
}
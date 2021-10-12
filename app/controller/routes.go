package controller

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/projectforge/app/telemetry/httpmetrics"
	"github.com/kyleu/projectforge/app/util"
)

//nolint
func AppRoutes() fasthttp.RequestHandler {
	r := router.New()

	r.GET("/", Home)
	r.GET("/about", About)
	r.GET("/theme", ThemeList)
	r.GET("/theme/{key}", ThemeEdit)
	r.POST("/theme/{key}", ThemeSave)
	r.GET(defaultSearchPath, Search)

	r.GET(defaultProfilePath, Profile)
	r.POST(defaultProfilePath, ProfileSave)
	r.GET("/auth/{key}", AuthDetail)
	r.GET("/auth/callback/{key}", AuthCallback)
	r.GET("/auth/logout/{key}", AuthLogout)

	// $PF_SECTION_START(routes)$
	r.GET("/doctor", Doctor)
	r.GET("/doctor/all", DoctorRunAll)
	r.GET("/doctor/{key}", DoctorRun)

	r.GET("/p", ProjectList)
	r.GET("/p/new", ProjectForm)
	r.POST("/p/new", ProjectCreate)
	r.GET("/p/{key}", ProjectDetail)
	r.GET("/p/{key}/edit", ProjectEdit)
	r.POST("/p/{key}/edit", ProjectSave)
	r.GET("/p/{key}/fs", ProjectFileRoot)
	r.GET("/p/{key}/fs/{path:*}", ProjectFile)
	r.GET("/p/{key}/svg", SVGList)
	r.GET("/p/{key}/svg/x/add", SVGAdd)
	r.GET("/p/{key}/svg/x/build", SVGBuild)
	r.GET("/p/{key}/svg/{icon}", SVGDetail)
	r.GET("/p/{key}/svg/{icon}/setapp", SVGSetApp)
	r.GET("/p/{key}/svg/{icon}/remove", SVGRemove)

	r.GET("/run/{act}", RunAllActions)
	r.GET("/run/{tgt}/{act}", RunAction)

	r.GET("/m", ModuleList)
	r.GET("/m/{key}", ModuleDetail)
	r.GET("/m/{key}/fs", ModuleFileRoot)
	r.GET("/m/{key}/fs/{path:*}", ModuleFile)

	r.GET("/test", TestList)
	r.GET("/test/{key}", TestRun)
	// $PF_SECTION_END(routes)$

	r.GET("/admin", Admin)
	r.GET("/admin/sandbox", SandboxList)
	r.GET("/admin/sandbox/{key}", SandboxRun)
	r.GET("/admin/settings", Settings)
	r.GET("/admin/{path:*}", Admin)

	r.GET("/favicon.ico", Favicon)
	r.GET("/robots.txt", RobotsTxt)
	r.GET("/assets/{_:*}", Static)

	r.OPTIONS("/", Options)
	r.OPTIONS("/{_:*}", Options)
	r.NotFound = NotFound

	p := httpmetrics.NewMetrics(util.AppKey)
	return fasthttp.CompressHandlerBrotliLevel(p.WrapHandler(r), fasthttp.CompressBrotliBestSpeed, fasthttp.CompressBestSpeed)
}

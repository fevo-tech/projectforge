package controller

import (
	"net/url"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"{{{ .Package }}}/app"
	"{{{ .Package }}}/app/controller/cutil"
	"{{{ .Package }}}/app/lib/theme"
	"{{{ .Package }}}/app/lib/user"
	"{{{ .Package }}}/app/util"
	"{{{ .Package }}}/views/vprofile"
)

func Profile(rc *fasthttp.RequestCtx) {
	act("profile", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(rc, as, ps)
	})
}
{{{ if .HasModule "marketing" }}}
func ProfileSite(rc *fasthttp.RequestCtx) {
	actSite("profile", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(rc, as, ps)
	})
}
{{{ end }}}
func profileAction(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (string, error) {
	ps.Title = "Profile"
	ps.Data = ps.Profile
	thm := as.Themes.Get(ps.Profile.Theme){{{ if .HasModule "oauth" }}}

	prvs, err := as.Auth.Providers()
	if err != nil {
		return "", errors.Wrap(err, "can't load providers")
	}{{{ end }}}

	redir := "/"
	ref := string(rc.Request.Header.Peek("Referer"))
	if ref != "" {
		u, err := url.Parse(ref)
		if err == nil && u != nil && u.Path != defaultProfilePath {
			redir = u.Path
		}
	}

	page := &vprofile.Profile{Profile: ps.Profile, Theme: thm, {{{ if .HasModule "oauth" }}}Providers: prvs, {{{ end }}}Referrer: redir}
	return render(rc, as, page, ps, "Profile")
}

func ProfileSave(rc *fasthttp.RequestCtx) {
	act("profile.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		n := ps.Profile.Clone()

		referrer := frm.GetStringOpt("referrer")
		if referrer == "" {
			referrer = defaultProfilePath
		}

		n.Name = frm.GetStringOpt("name")
		n.Mode = frm.GetStringOpt("mode")
		n.Theme = frm.GetStringOpt("theme")
		if n.Theme == theme.ThemeDefault.Key {
			n.Theme = ""
		}

		err = cutil.SaveProfile(n, rc, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}

		return returnToReferrer("Saved profile", referrer, rc, ps)
	})
}

func loadProfile(session util.ValueMap) (*user.Profile, error) {
	x, ok := session["profile"]
	if !ok {
		return user.DefaultProfile.Clone(), nil
	}
	s, ok := x.(string)
	if !ok {
		return user.DefaultProfile.Clone(), nil
	}
	p := &user.Profile{}
	err := util.FromJSON([]byte(s), p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func returnToReferrer(msg string, dflt string, rc *fasthttp.RequestCtx, ps *cutil.PageState) (string, error) {
	refer := ""
	referX, ok := ps.Session[cutil.ReferKey]
	if ok {
		refer, ok = referX.(string)
		if ok {
			_ = cutil.RemoveFromSession(cutil.ReferKey, rc, ps.Session, ps.Logger)
		}
	}
	if refer == "" {
		refer = dflt
	}
	return flashAndRedir(true, msg, refer, rc, ps)
}

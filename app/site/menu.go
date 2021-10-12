// Code generated by projectforge.dev using code from the [marketing] module, which is under license [CC0]
package site

import (
	"github.com/kyleu/projectforge/app/menu"
	"github.com/kyleu/projectforge/app/user"
)

const (
	keyIntro      = "introduction"
	keyInstall    = "install"
	keyDownload   = "download"
	keyQuickStart = "quickstart"
	keyContrib    = "contributing"
	keyTech       = "technology"
)

func Menu(p *user.Profile, a user.Accounts) menu.Items {
	return menu.Items{
		{Key: keyIntro, Title: "Introduction", Icon: "heart", Route: "/"+keyIntro},
		{Key: keyInstall, Title: "Install", Icon: "code", Route: "/"+keyInstall},
		{Key: keyDownload, Title: "Download", Icon: "download", Route: "/"+keyDownload},
		{Key: keyQuickStart, Title: "Quick Start", Icon: "bolt", Route: "/"+keyQuickStart},
		{Key: keyContrib, Title: "Contributing", Icon: "cog", Route: "/"+keyContrib},
		{Key: keyTech, Title: "Technology", Icon: "cog", Route: "/"+keyTech},
	}
}

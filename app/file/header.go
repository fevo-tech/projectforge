package file

import (
	"strings"

	"go.uber.org/zap"
)

const headerContent = "Code generated by Project Forge, see https://projectforge.dev for details."

func contentWithHeader(t Type, c string, logger *zap.SugaredLogger) string {
	switch t.Key {
	case TypeConf.Key:
		return c
	case TypeCSS.Key:
		return "/* " + headerContent + " */\n" + c
	case TypeDocker.Key:
		return "# " + headerContent + "\n" + c
	case TypeGo.Key:
		return "// " + headerContent + "\n" + c
	case TypeGoMod.Key:
		return "// " + headerContent + "\n" + c
	case TypeHCL.Key:
		return "# " + headerContent + "\n" + c
	case TypeHTML.Key:
		return "<!-- " + headerContent + " -->\n" + c
	case TypeIcons.Key:
		return c
	case TypeIgnore.Key:
		return c
	case TypeJavaScript.Key:
		return "// " + headerContent + "\n" + c
	case TypeJSON.Key:
		return c
	case TypeMakefile.Key:
		return "# " + headerContent + "\n" + c
	case TypeMarkdown.Key:
		return "<!--- " + headerContent + " -->\n" + c
	case TypePList.Key:
		return c
	case TypeShell.Key:
		return secondLine(c, "# "+headerContent)
	case TypeSVG.Key:
		return c
	case TypeTypeScript.Key:
		return "// " + headerContent + "\n" + c
	case TypeYAML.Key:
		return "# " + headerContent + "\n" + c
	default:
		logger.Warnf("unhandled header for file type [%s]", t.Title)
		return c
	}
}

func secondLine(content string, rplc string) string {
	idx := strings.Index(content, "\n")
	if idx == -1 {
		return content
	}
	return content[0:idx] + "\n" + rplc + content[idx:]
}

package project

import (
	"strings"

	"golang.org/x/exp/slices"

	"projectforge.dev/projectforge/app/lib/theme"
	"projectforge.dev/projectforge/app/project/export/model"
	"projectforge.dev/projectforge/app/util"
)

var DefaultIcon = "code"

type Project struct {
	Key     string   `json:"key"`
	Name    string   `json:"name,omitempty"`
	Icon    string   `json:"icon,omitempty"`
	Exec    string   `json:"exec,omitempty"`
	Version string   `json:"version"`
	Package string   `json:"package,omitempty"`
	Args    string   `json:"args,omitempty"`
	Port    int      `json:"port,omitempty"`
	Modules []string `json:"modules"`
	Ignore  []string `json:"ignore,omitempty"`
	Tags    []string `json:"tags"`

	Info  *Info        `json:"info,omitempty"`
	Theme *theme.Theme `json:"theme,omitempty"`
	Build *Build       `json:"build,omitempty"`

	ExportArgs *model.Args   `json:"-"`
	Config     util.ValueMap `json:"-"`
	Path       string        `json:"-"`
	Parent     string        `json:"-"`
	Error      string        `json:"-"`
}

func NewProject(key string, path string) *Project {
	return &Project{Key: key, Version: "0.0.0", Path: path}
}

func (p *Project) Title() string {
	if p.Name == "" {
		return p.Key
	}
	return p.Name
}

func (p *Project) NameSafe() string {
	return strings.ReplaceAll(strings.ReplaceAll(p.Name, "-", ""), " ", "")
}

func (p *Project) Executable() string {
	if p.Exec == "" {
		return p.Key
	}
	return p.Exec
}

func (p *Project) CleanKey() string {
	return clean(p.Key)
}

func (p *Project) IconSafe() string {
	if _, ok := util.SVGLibrary[p.Icon]; ok {
		return p.Icon
	}
	return DefaultIcon
}

func (p *Project) DescriptionSafe() string {
	if p.Info == nil {
		return ""
	}
	if p.Info.Description == "" {
		return p.Info.Summary
	}
	return p.Info.Description
}

func (p *Project) HasModule(key string) bool {
	return slices.Contains(p.Modules, key)
}

func (p *Project) ToMap() util.ValueMap {
	return util.ValueMap{
		"key": p.Key, "name": p.Name, "icon": p.Icon, "exec": p.Exec,
		"version": p.Version, "package": p.Package, "args": p.Key, "port": p.Port,
		"modules": p.Modules, "path": p.Path,
	}
}

func (p *Project) WebPath() string {
	return "/p/" + p.Key
}

func (p *Project) ModuleArgExport(pSvc *Service, logger util.Logger) (*model.Args, error) {
	if p.ExportArgs == nil && p.HasModule("export") {
		var err error
		p.ExportArgs, err = pSvc.loadExportArgs(pSvc.GetFilesystem(p), logger)
		if err != nil {
			return nil, err
		}
		p.ExportArgs.Modules = p.Modules
	}
	return p.ExportArgs, nil
}

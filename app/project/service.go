package project

import (
	"strings"
	"sync"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"projectforge.dev/projectforge/app/lib/filesystem"
	"projectforge.dev/projectforge/app/util"
)

const (
	ConfigDir          = "." + util.AppKey
	additionalFilename = ConfigDir + "/additional-projects.json"
)

type Service struct {
	cache       map[string]*Project
	cacheLock   sync.RWMutex
	filesystems map[string]filesystem.FileLoader
	fsLock      sync.RWMutex
}

func NewService() *Service {
	return &Service{cache: map[string]*Project{}, filesystems: map[string]filesystem.FileLoader{}}
}

func (s *Service) GetFilesystem(prj *Project) filesystem.FileLoader {
	s.fsLock.Lock()
	defer s.fsLock.Unlock()
	curr, ok := s.filesystems[prj.Key]
	if ok {
		return curr
	}
	fs := filesystem.NewFileSystem(prj.Path)
	s.filesystems[prj.Key] = fs
	return fs
}

func (s *Service) Refresh(logger util.Logger) (Projects, error) {
	s.cache = map[string]*Project{}
	root, err := s.add(".", nil, logger)
	if err != nil {
		return nil, err
	}
	if add, ok := s.getAdditional(logger); ok {
		for _, a := range add {
			if _, err := s.add(a, root, logger); err != nil {
				return nil, err
			}
		}
	}
	return s.Projects(), nil
}

func (s *Service) Get(key string) (*Project, error) {
	s.cacheLock.Lock()
	ret, ok := s.cache[key]
	s.cacheLock.Unlock()
	if ok {
		return ret, nil
	}
	return nil, errors.Errorf("no project with key [%s] found among %d candidates [%s]", key, len(s.cache), strings.Join(s.Keys(), ", "))
}

func (s *Service) Keys() []string {
	s.cacheLock.Lock()
	keys := make(map[string]string, len(s.cache))
	titles := make([]string, 0, len(s.cache))
	for k, v := range s.cache {
		tl := strings.ToLower(v.Title())
		keys[tl] = k
		titles = append(titles, tl)
	}
	s.cacheLock.Unlock()
	slices.Sort(titles)
	ret := make([]string, 0, len(titles))
	for _, title := range titles {
		ret = append(ret, keys[title])
	}
	return ret
}

func (s *Service) Projects() Projects {
	keys := s.Keys()
	ret := make(Projects, 0, len(keys))
	for _, key := range keys {
		p, _ := s.Get(key)
		ret = append(ret, p)
	}
	return ret
}

func (s *Service) ByPath(path string) *Project {
	s.cacheLock.Lock()
	defer s.cacheLock.Unlock()
	for _, v := range s.cache {
		if v.Path == path {
			return v
		}
	}
	return nil
}

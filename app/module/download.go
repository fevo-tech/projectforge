package module

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kyleu/projectforge/app/filesystem"
	"github.com/pkg/errors"
)

func (s *Service) Download(key string, url string) error {
	if url == "" {
		base := "http://localhost:40002"
		if o := os.Getenv("projectforge_update_url"); o != "" {
			base = o
		}
		url = fmt.Sprintf(base+ "/projectforge.module.%s.zip", key)
	}
	s.logger.Infof("downloading module [%s] from URL [%s]", key, url)
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return errors.Wrapf(err, "invalid URL [%s] for module [%s]", url, key)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "unable to retrieve module [%s] from [%s]", key, url)
	}
	defer func() { _ = response.Body.Close() }()

	if response.StatusCode != http.StatusOK {
		return errors.Errorf("module [%s] load request to [%s] returned status code [%d]", key, url, response.StatusCode)
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.Errorf("unable to read body from module [%s] load request to [%s]", key, url)
	}

	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return errors.Errorf("unable to unzip body from module [%s] load request to [%s]", key, url)
	}

	for _, f := range r.File {
		fn := filepath.Join(key, f.Name)
		if f.FileInfo().IsDir() {
			err = s.config.CreateDirectory(fn)
			if err != nil {
				return errors.Errorf("unable to create directory [%s] for module [%s] from [%s]", fn, key, url)
			}
			continue
		}
		o, err := f.Open()
		if err != nil {
			return errors.Errorf("unable to read file [%s] from module [%s] from [%s]", fn, key, url)
		}
		content, err := io.ReadAll(o)
		err = s.config.WriteFile(fn, content, filesystem.DefaultMode, false)
		if err != nil {
			return errors.Errorf("unable to write file [%s] for module [%s] from [%s]", fn, key, url)
		}
	}

	return nil
}


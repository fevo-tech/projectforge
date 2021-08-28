package search

import (
	"strings"
	"sync"

	"github.com/pkg/errors"
)

type Provider func(*Params) (Results, error)

func Search(params *Params) (Results, []error) {
	// $PF_SECTION_START(search_functions)$
	testFunc := func(p *Params) (Results, error) {
		return Results{{URL: "/search?q=test", Title: "Test Result", Match: p.Q + "!!!"}}, errors.New("!!!")
	}

	allProviders := []Provider{testFunc}
	// $PF_SECTION_END(search_functions)$

	ret := Results{}
	var errs []error
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(allProviders))
	params.Q = strings.TrimSpace(params.Q)

	for _, p := range allProviders {
		f := p
		go func() {
			res, err := f(params)
			mu.Lock()
			if err != nil {
				errs = append(errs, err)
			}
			ret = append(ret, res...)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	ret.Sort()
	return ret, errs
}
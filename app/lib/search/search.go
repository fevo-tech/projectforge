// Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/lib/search/result"
	"projectforge.dev/projectforge/app/lib/telemetry"
	"projectforge.dev/projectforge/app/util"
)

type Provider func(context.Context, *Params, *app.State, *cutil.PageState, util.Logger) (result.Results, error)

func Search(ctx context.Context, params *Params, as *app.State, page *cutil.PageState) (result.Results, []error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "search", page.Logger)
	defer span.Complete()

	if params.Q == "" {
		return nil, nil
	}
	var allProviders []Provider
	// $PF_SECTION_START(search_functions)$
	projectFunc := func(ctx context.Context, p *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		return as.Services.Projects.Search(ctx, p.Q, logger)
	}
	moduleFunc := func(ctx context.Context, p *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		return as.Services.Modules.Search(ctx, p.Q, logger)
	}
	allProviders = append(allProviders, projectFunc, moduleFunc)
	// $PF_SECTION_END(search_functions)$
	if len(allProviders) == 0 {
		return nil, []error{errors.New("no search providers configured")}
	}

	params.Q = strings.TrimSpace(params.Q)

	results, errs := util.AsyncCollect(allProviders, func(item Provider) (result.Results, error) {
		return item(ctx, params, as, page, logger)
	})

	ret := make(result.Results, 0, len(results)*len(results))
	for _, x := range results {
		ret = append(ret, x...)
	}

	ret.Sort()
	return ret, errs
}

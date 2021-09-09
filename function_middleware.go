package starlarkhelper

import (
	"context"
	"go.starlark.net/starlark"
)

type Middleware func(ctx context.Context, h *Helper, do BuiltinFunc) (starlark.Value, error)

func ChainMiddleware(middlewares ...Middleware) Middleware {
	if len(middlewares) == 0 {
		return nil
	} else if len(middlewares) == 1 {
		return middlewares[0]
	} else {
		return func(ctx context.Context, h *Helper, do BuiltinFunc) (starlark.Value, error) {
			return middlewares[0](ctx, h, getChainMiddleware(middlewares, 0, do))
		}
	}
}

func getChainMiddleware(middlewares []Middleware, curr int, finalDo BuiltinFunc) BuiltinFunc {
	if curr == len(middlewares)-1 {
		return finalDo
	}

	return func(ctx context.Context, h *Helper) (starlark.Value, error) {
		return middlewares[curr+1](ctx, h, getChainMiddleware(middlewares, curr+1, finalDo))
	}
}

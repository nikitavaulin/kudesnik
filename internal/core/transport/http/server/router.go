package core_http_server

import (
	"fmt"
	"net/http"

	tools_jwt "github.com/nikitavaulin/kudesnik/internal/core/tools/jwt"
	core_http_middleware "github.com/nikitavaulin/kudesnik/internal/core/transport/http/middleware"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion  ApiVersion
	jwtProvider *tools_jwt.JwtProvider
	middleware  []core_http_middleware.Middleware
}

func NewAPIVersionRouter(apiVersion ApiVersion, jwt *tools_jwt.JwtProvider, middlewares ...core_http_middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:    http.NewServeMux(),
		apiVersion:  apiVersion,
		jwtProvider: jwt,
		middleware:  middlewares,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		r.Handle(pattern, route.WithMiddleware(r.jwtProvider))
	}
}

func (r *APIVersionRouter) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(r, r.middleware...)
}

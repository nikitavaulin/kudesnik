package core_http_server

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	tools_jwt "github.com/nikitavaulin/kudesnik/internal/core/tools/jwt"
	core_http_middleware "github.com/nikitavaulin/kudesnik/internal/core/transport/http/middleware"
)

type Route struct {
	Method       string
	Path         string
	Handler      http.HandlerFunc
	Middlewares  []core_http_middleware.Middleware
	RequiresAuth bool
	AllowedRoles []domain.Role
}

func (r *Route) WithMiddleware(jwtProvider *tools_jwt.JwtProvider) http.Handler {
	handler := r.Handler
	if r.RequiresAuth {

		if len(r.AllowedRoles) > 0 {
			handler = core_http_middleware.Authorize(r.AllowedRoles...)(handler)
		}

		handler = core_http_middleware.Authenticate(jwtProvider)(handler)
	}

	return core_http_middleware.ChainMiddleware(handler, r.Middlewares...)
}

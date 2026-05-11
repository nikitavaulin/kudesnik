package core_http_middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	tools_jwt "github.com/nikitavaulin/kudesnik/internal/core/tools/jwt"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

func Authenticate(jwtProvider *tools_jwt.JwtProvider) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPReponseHandler(log, w)

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "missing authorization token")
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := jwtProvider.DecodeClaims(token)
			if err != nil {
				responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "invalid authorization token")
				return
			}

			ctx = domain.RoleToContext(ctx, domain.Role(claims.Role))
			ctx = domain.EmailToContext(ctx, claims.Email)
			ctx = domain.UserIDToContext(ctx, uuid.MustParse(claims.ID))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Authorize(allowedRoles ...domain.Role) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPReponseHandler(log, w)

			userRole, ok := domain.RoleFromContext(ctx)
			if !ok {
				responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "role is empty")
				return
			}

			if slices.Contains(allowedRoles, userRole) {
				next.ServeHTTP(w, r)
				return
			}

			responseHandler.ErrorResponse(core_errors.ErrForbidden, "no access")
		})
	}
}

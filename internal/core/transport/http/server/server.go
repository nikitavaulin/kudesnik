package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	tools_jwt "github.com/nikitavaulin/kudesnik/internal/core/tools/jwt"
	core_http_middleware "github.com/nikitavaulin/kudesnik/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux         *http.ServeMux
	config      *HTTPServerConfig
	log         *core_logger.Logger
	middlewares []core_http_middleware.Middleware
	jwtProvider *tools_jwt.JwtProvider
	// staticDirPath string
	// staticURL     string
}

func NewHTTPServer(
	config *HTTPServerConfig,
	log *core_logger.Logger,
	jwtProvider *tools_jwt.JwtProvider,
	middlewares ...core_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:         http.NewServeMux(),
		config:      config,
		log:         log,
		middlewares: middlewares,
		jwtProvider: jwtProvider,
		// staticURL:   config.StaticURL,
	}
}

func (s *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)
		s.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router.WithMiddleware()),
		)
	}
}

func (s *HTTPServer) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		s.mux.Handle(pattern, route.WithMiddleware(s.jwtProvider))
	}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(s.mux, s.middlewares...)

	server := &http.Server{
		Addr:    s.config.Address,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.log.Warn("start HTTP server", zap.String("address", s.config.Address))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and serve http server: %w", err)
		}
	case <-ctx.Done():
		s.log.Warn("shutdown HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownDuration)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		s.log.Warn("HTTP server stopped successfully")
	}

	return nil
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_middleware "github.com/nikitavaulin/kudesnik/internal/core/transport/http/middleware"
	core_http_server "github.com/nikitavaulin/kudesnik/internal/core/transport/http/server"
	products_transport_http "github.com/nikitavaulin/kudesnik/internal/features/products/transport/http"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Hello, kudesnik app!")

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Printf("failed to create application logger: %s\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Starting Kudesnik application!")

	productsTransportHTTP := products_transport_http.NewProductsHTTPHandler(nil)
	productsRouters := productsTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(
		productsRouters...,
	)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewHTTPServerConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/nikitavaulin/kudesnik/internal/core/transport/http/middleware"
	core_http_server "github.com/nikitavaulin/kudesnik/internal/core/transport/http/server"
	product_categories_repository "github.com/nikitavaulin/kudesnik/internal/features/product_categories/repository"
	product_categories_service "github.com/nikitavaulin/kudesnik/internal/features/product_categories/service"
	product_categories_transport_http "github.com/nikitavaulin/kudesnik/internal/features/product_categories/transport"
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

	pool, err := core_postgres_pool.NewConnectionPool(ctx, core_postgres_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to create postgresql connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("Starting Kudesnik application!")

	logger.Debug("initializing features", zap.String("feature", "Product-Categories"))

	productCategoriesRepo := product_categories_repository.NewProductCategoriesRepository(pool)
	productCategoriesService := product_categories_service.NewProductCategoriesService(productCategoriesRepo)
	productCategoriesHTTPTransport := product_categories_transport_http.NewProductCategoryHTTPHandler(productCategoriesService)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(
		productCategoriesHTTPTransport.Routes()...,
	)

	logger.Debug("initializing HTTP server")
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

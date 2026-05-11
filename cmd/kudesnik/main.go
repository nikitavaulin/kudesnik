package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_pgx_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool/pgx"
	tools_jwt "github.com/nikitavaulin/kudesnik/internal/core/tools/jwt"
	core_http_middleware "github.com/nikitavaulin/kudesnik/internal/core/transport/http/middleware"
	core_http_server "github.com/nikitavaulin/kudesnik/internal/core/transport/http/server"
	admin_repository_postgres "github.com/nikitavaulin/kudesnik/internal/features/admin/repository/postgres"
	admin_service "github.com/nikitavaulin/kudesnik/internal/features/admin/service"
	admin_transport_http "github.com/nikitavaulin/kudesnik/internal/features/admin/transport/http"
	product_categories_repository "github.com/nikitavaulin/kudesnik/internal/features/product_categories/repository"
	product_categories_service "github.com/nikitavaulin/kudesnik/internal/features/product_categories/service"
	product_categories_transport_http "github.com/nikitavaulin/kudesnik/internal/features/product_categories/transport"
	products_repository_postges "github.com/nikitavaulin/kudesnik/internal/features/products/repository/postgres"
	products_service "github.com/nikitavaulin/kudesnik/internal/features/products/service"
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

	pool, err := core_pgx_pool.NewPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to create pgx connection pool", zap.Error(err))
	}
	defer pool.Close()

	jwtProvider := tools_jwt.NewJWTProvider(tools_jwt.NewConfigMust())

	logger.Debug("Starting Kudesnik application!")

	logger.Debug("initializing features")

	productsCategoriesRepo := product_categories_repository.NewProductCategoriesRepository(pool)
	productsCategoriesService := product_categories_service.NewProductCategoriesService(productsCategoriesRepo)
	productsCategoriesHTTPTransport := product_categories_transport_http.NewProductCategoryHTTPHandler(productsCategoriesService)

	productsRepo := products_repository_postges.NewProductsRepositoryPostgres(pool)
	productsSevice := products_service.NewProductsService(productsRepo)
	productsTransport := products_transport_http.NewProductsHTTPHandler(productsSevice)

	adminsRepo := admin_repository_postgres.NewAdminRepositoryPostgres(pool)
	adminsService := admin_service.NewAdminServie(adminsRepo, jwtProvider)
	adminsTransport := admin_transport_http.NewAdminTrasnsportHTTPHandler(adminsService)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1, jwtProvider)
	apiVersionRouter.RegisterRoutes(productsCategoriesHTTPTransport.Routes()...)
	apiVersionRouter.RegisterRoutes(productsTransport.Routes()...)
	apiVersionRouter.RegisterRoutes(adminsTransport.Routes()...)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewHTTPServerConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

}

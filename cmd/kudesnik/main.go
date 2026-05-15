package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_pgx_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool/pgx"
	tools_jwt "github.com/nikitavaulin/kudesnik/internal/core/tools/jwt"
	core_http_middleware "github.com/nikitavaulin/kudesnik/internal/core/transport/http/middleware"
	core_http_server "github.com/nikitavaulin/kudesnik/internal/core/transport/http/server"
	admin_repository_postgres "github.com/nikitavaulin/kudesnik/internal/features/admin/repository/postgres"
	admin_service "github.com/nikitavaulin/kudesnik/internal/features/admin/service"
	admin_transport_http "github.com/nikitavaulin/kudesnik/internal/features/admin/transport/http"
	customer_requests_repository_postgres "github.com/nikitavaulin/kudesnik/internal/features/customer_requests/repository/postgres"
	customer_requests_service "github.com/nikitavaulin/kudesnik/internal/features/customer_requests/service"
	customer_requests_transport_http "github.com/nikitavaulin/kudesnik/internal/features/customer_requests/transport/http"
	image_local_storage "github.com/nikitavaulin/kudesnik/internal/features/images/repository/local"
	images_service "github.com/nikitavaulin/kudesnik/internal/features/images/service"
	images_transport_http "github.com/nikitavaulin/kudesnik/internal/features/images/transport/http"
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

	projectRoot, err := getProjectRoot()
	if err != nil {
		logger.Fatal("failed to get project root", zap.Error(err))
	}

	uploadsPath, err := getUploadsFolder()
	if err != nil {
		logger.Fatal("failed to get uploads path", zap.Error(err))
	}

	absUploadPath := filepath.Join(projectRoot, uploadsPath)

	if err := os.MkdirAll(absUploadPath, 0755); err != nil {
		logger.Fatal("failed to create uploads directory", zap.Error(err), zap.String("upload path", absUploadPath))
	}

	jwtProvider := tools_jwt.NewJWTProvider(tools_jwt.NewConfigMust())

	logger.Debug("Starting Kudesnik application!")
	logger.Debug("initializing features")

	localStorage := image_local_storage.NewLocalStorage(image_local_storage.NewLocalStorageConfig(projectRoot, "/static", uploadsPath))
	imageServie := images_service.NewImageService(localStorage)
	imageTransport := images_transport_http.NewImageTransportHTTPHandler(imageServie)

	productsCategoriesRepo := product_categories_repository.NewProductCategoriesRepository(pool)
	productsCategoriesService := product_categories_service.NewProductCategoriesService(productsCategoriesRepo)
	productsCategoriesHTTPTransport := product_categories_transport_http.NewProductCategoryHTTPHandler(productsCategoriesService)

	productsRepo := products_repository_postges.NewProductsRepositoryPostgres(pool)
	productsSevice := products_service.NewProductsService(productsRepo)
	productsTransport := products_transport_http.NewProductsHTTPHandler(productsSevice, imageServie)

	adminsRepo := admin_repository_postgres.NewAdminRepositoryPostgres(pool)
	adminsService := admin_service.NewAdminServie(adminsRepo, jwtProvider)
	adminsTransport := admin_transport_http.NewAdminTrasnsportHTTPHandler(adminsService)

	customerRequestsRepo := customer_requests_repository_postgres.NewCustomerRequestsRepository(pool)
	customerRequestsService := customer_requests_service.NewCustomerRequestsService(customerRequestsRepo, productsRepo)
	customerRequestsTransport := customer_requests_transport_http.NewCustomerRequestsTransportHTTP(customerRequestsService)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1, jwtProvider)
	apiVersionRouter.RegisterRoutes(productsCategoriesHTTPTransport.Routes()...)
	apiVersionRouter.RegisterRoutes(productsTransport.Routes()...)
	apiVersionRouter.RegisterRoutes(adminsTransport.Routes()...)
	apiVersionRouter.RegisterRoutes(customerRequestsTransport.Routes()...)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewHTTPServerConfigMust(),
		logger,
		jwtProvider,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	httpServer.RegisterAPIRouters(apiVersionRouter)
	httpServer.RegisterRoutes(imageTransport.Routes()...)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

}

func getProjectRoot() (string, error) {
	root := os.Getenv("PROJECT_ROOT")
	if root == "" {
		return "", fmt.Errorf("failed to get env var PROJECT_ROOT")
	}
	return root, nil
}

func getUploadsFolder() (string, error) {
	folderPath := os.Getenv("UPLOADS_FOLDER")
	if folderPath == "" {
		return "", fmt.Errorf("failed to get env var UPLOADS_FOLDER")
	}
	return "/" + folderPath, nil
}

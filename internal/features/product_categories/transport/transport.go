package product_categories_transport_http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_http_server "github.com/nikitavaulin/kudesnik/internal/core/transport/http/server"
)

type ProductCategoryHTTPHandler struct {
	categoriesService ProductCategoryService
}

type ProductCategoryService interface {
	CreateProductCategory(
		ctx context.Context,
		category domain.ProductCategory,
	) (domain.ProductCategory, error)

	GetProductCategories(ctx context.Context, limit, offset *int) ([]domain.ProductCategory, error)

	GetProductCategory(ctx context.Context, categoryID uuid.UUID) (domain.ProductCategory, error)

	DeleteProductCategory(ctx context.Context, categoryID uuid.UUID) error
}

func NewProductCategoryHTTPHandler(categoriesService ProductCategoryService) *ProductCategoryHTTPHandler {
	return &ProductCategoryHTTPHandler{
		categoriesService: categoriesService,
	}
}

func (h *ProductCategoryHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/product-categories",
			Handler: h.CreateCategory,
		},
		{
			Method:  http.MethodGet,
			Path:    "/product-categories",
			Handler: h.GetCategories,
		},
		{
			Method:  http.MethodGet,
			Path:    "/product-categories/{id}",
			Handler: h.GetProductCategory,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/product-categories/{id}",
			Handler: h.DeleteProductCategory,
		},
	}
}

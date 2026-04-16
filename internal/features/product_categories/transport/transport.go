package product_categories_transport_http

import (
	"context"
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_http_server "github.com/nikitavaulin/kudesnik/internal/core/transport/http/server"
)

type ProductCategoryHTTPHandler struct {
	categoryService ProductCategoryService
}

type ProductCategoryService interface {
	CreateProductCategory(
		ctx context.Context,
		category domain.ProductCategory,
	) (domain.ProductCategory, error)
}

func NewProductCategoryHTTPHandler(categoryService ProductCategoryService) *ProductCategoryHTTPHandler {
	return &ProductCategoryHTTPHandler{
		categoryService: categoryService,
	}
}

func (h *ProductCategoryHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/product-categories",
			Handler: h.CreateCategory,
		},
	}
}

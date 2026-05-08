package products_transport_http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_http_server "github.com/nikitavaulin/kudesnik/internal/core/transport/http/server"
)

type ProductsHTTPHandler struct {
	productsService ProductsService
}

type ProductsService interface {
	CreateProduct(ctx context.Context, product domain.BaseProduct) (domain.BaseProduct, error)
	GetProduct(ctx context.Context, id uuid.UUID) (domain.BaseProduct, error)
	GetProducts(ctx context.Context, categoryID *uuid.UUID, limit, offset *int) ([]domain.BaseProduct, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	DeleteProducts(ctx context.Context, IDs []uuid.UUID) error
	UpdateProductVisability(ctx context.Context, id uuid.UUID, isVisible bool) error
	UpdateProductsVisability(ctx context.Context, IDs []uuid.UUID, isVisible bool) error
}

func NewProductsHTTPHandler(productsService ProductsService) *ProductsHTTPHandler {
	return &ProductsHTTPHandler{
		productsService: productsService,
	}
}

func (h *ProductsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/products",
			Handler: h.CreateProduct,
		},
		{
			Method:  http.MethodGet,
			Path:    "/products",
			Handler: h.GetProducts,
		},
		{
			Method:  http.MethodGet,
			Path:    "/products/{id}",
			Handler: h.GetProduct,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/products/{id}",
			Handler: h.DeleteProduct,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/products",
			Handler: h.DeleteProducts,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/products/visibility/{id}",
			Handler: h.UpdateProductVisability,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/products/visibility",
			Handler: h.UpdateProductsVisability,
		},
	}
}

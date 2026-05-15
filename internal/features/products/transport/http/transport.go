package products_transport_http

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_http_server "github.com/nikitavaulin/kudesnik/internal/core/transport/http/server"
)

type ProductsHTTPHandler struct {
	productsService ProductsService
	imageService    ImageService
}

type ProductsService interface {
	CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error)

	GetProducts(ctx context.Context, category *domain.ProductCategoryCode, order *string, limit, offset *int) ([]domain.ProductBaseDetailed, error)
	GetProduct(ctx context.Context, id uuid.UUID, category domain.ProductCategoryCode) (domain.ProductDetailed, error)
	GetProductBase(ctx context.Context, id uuid.UUID) (domain.ProductBase, error)

	DeleteProduct(ctx context.Context, id uuid.UUID) error
	DeleteProducts(ctx context.Context, IDs []uuid.UUID) error

	UpdateProductVisability(ctx context.Context, id uuid.UUID, isVisible bool) error
	UpdateProductsVisability(ctx context.Context, IDs []uuid.UUID, isVisible bool) error
	PatchProduct(ctx context.Context, id uuid.UUID, patch domain.ProductPatch) (domain.Product, error)
}

type ImageService interface {
	DeleteProductImages(ctx context.Context, imagePath, thumbnailPath string) error
	UploadImage(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*domain.ImageUploadResult, error)
}

func NewProductsHTTPHandler(productsService ProductsService, imageService ImageService) *ProductsHTTPHandler {
	return &ProductsHTTPHandler{
		productsService: productsService,
		imageService:    imageService,
	}
}

func (h *ProductsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:       http.MethodPost,
			Path:         "/products/{category_code}",
			Handler:      h.CreateProduct,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole, domain.ManagerRole},
		},
		{
			Method:  http.MethodGet,
			Path:    "/products",
			Handler: h.GetProducts,
		},
		{
			Method:  http.MethodGet,
			Path:    "/products/{category_code}/{id}",
			Handler: h.GetProduct,
		},
		{
			Method:       http.MethodDelete,
			Path:         "/products/{id}",
			Handler:      h.DeleteProduct,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole, domain.ManagerRole},
		},
		{
			Method:       http.MethodDelete,
			Path:         "/products",
			Handler:      h.DeleteProducts,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole, domain.ManagerRole},
		},
		{
			Method:       http.MethodPatch,
			Path:         "/products/visibility/p/{id}",
			Handler:      h.UpdateProductVisability,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.ManagerRole, domain.AdminRole},
		},
		{
			Method:       http.MethodPatch,
			Path:         "/products/visibility",
			Handler:      h.UpdateProductsVisability,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole, domain.ManagerRole},
		},
		{
			Method:       http.MethodPatch,
			Path:         "/products/{category_code}/{id}",
			Handler:      h.PatchProduct,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole, domain.ManagerRole},
		},
		{
			Method:       http.MethodPost,
			Path:         "/products/image/{id}",
			Handler:      h.UploadProductImage,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.ManagerRole, domain.AdminRole},
		},
		{
			Method:       http.MethodDelete,
			Path:         "/products/image/{id}",
			Handler:      h.DeleteProductImage,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.ManagerRole, domain.AdminRole},
		},
	}
}

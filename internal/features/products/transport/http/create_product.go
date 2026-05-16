package products_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type CreateProductResponse ProductDTOResponse

// CreateProduct godoc
// @Summary Создать товар
// @Description Создать товар (требуется указать конкретную категорию)
// @Security BearerAuth
// @Tags products
// @Accept json
// @Produce json
// @Param category_code path string true "код категории товар которой создается (windows, entrance-doors, interior-doors, balconies, others)"
// @Param request body domain.ProductBase true "CreateProduct тело запроса, если category_code=others"
// @Success 201 {object} CreateProductResponse "Созданный товар"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /products/{category_code} [post]
func (h *ProductsHTTPHandler) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(logger, rw)

	logger.Debug("invoke create product handler")

	category, err := core_http_request.GetCategoryCodeFromPath(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category_code")
		return
	}

	product := domain.GetProductEmptyInstance(string(category))

	if err := core_http_request.DecodeRequest(r, &product); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	product, err = h.productsService.CreateProduct(ctx, product)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create product")
		return
	}

	responseHandler.JSONResponse(product, http.StatusCreated)
}

package products_transport_http

import (
	"net/http"

	_ "github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type GetProductResponseDTO ProductDTOResponse

// GetProduct godoc
// @Summary Получить товар
// @Description Получить товар (требуется указать конкретную категорию)
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "ID товара" Format(uuid)
// @Param category_code path string true "код категории товар которой создается (windows, entrance-doors, interior-doors, balconies, others)"
// @Success 200 {object} domain.ProductDetailed "Созданный товар"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /products/{category_code}/{id} [get]
func (h *ProductsHTTPHandler) GetProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get product handler")

	category, err := core_http_request.GetCategoryCodeFromPath(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category_code")
		return
	}

	productID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get productID")
		return
	}

	product, err := h.productsService.GetProduct(ctx, productID, category)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get product by ID")
		return
	}

	responseHandler.JSONResponse(product, http.StatusOK)
}

package product_categories_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type GetProductCategoryResponse ProductCategoryDTOResponse

// GetProductCategory godoc
// @Summary Получить категорию товаров
// @Description Получить категорию товаров по коду
// @Tags product-categories
// @Produce json
// @Param category_code path string true "Код получаемой категории"
// @Success 200 {object} GetProductCategoryResponse "полученная категория"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /product-categories/{category_code} [get]
func (h *ProductCategoryHTTPHandler) GetProductCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get product category handler")

	categoryCode, err := core_http_request.GetCategoryCodeFromPath(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category_code")
		return
	}

	category, err := h.categoriesService.GetProductCategory(ctx, categoryCode)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category by ID")
		return
	}

	categoryDTO := dtoFromDomain(category)
	responseHandler.JSONResponse(categoryDTO, http.StatusOK)
}

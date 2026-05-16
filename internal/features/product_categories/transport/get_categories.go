package product_categories_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type GetCategoriesResponse []ProductCategoryDTOResponse

// GetCategories godoc
// @Summary Получить список категорий
// @Description Получить список категорий, есть пагинация.
// @Tags product-categories
// @Produce json
// @Param limit query int false "лимит возвращаемых категорий"
// @Param offset query int false "смещение возвращаемых категорий"
// @Success 200 {object} GetCategoriesResponse "полученные категории"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /product-categories [get]
func (h *ProductCategoryHTTPHandler) GetCategories(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get product categories handler")

	limit, offset, err := core_http_request.GetLimitOffsetParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset params")
		return
	}

	categories, err := h.categoriesService.GetProductCategories(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get product categories")
		return
	}

	categoriesDTO := dtosFromDomains(categories)
	responseHandler.JSONResponse(categoriesDTO, http.StatusOK)
}

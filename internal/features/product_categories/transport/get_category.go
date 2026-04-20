package product_categories_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
	core_http_utils "github.com/nikitavaulin/kudesnik/internal/core/transport/http/utils"
)

type GetProductCategoryResponse ProductCategoryDTOResponse

func (h *ProductCategoryHTTPHandler) GetProductCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get product category handler")

	categoryID, err := core_http_utils.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get categoryID")
		return
	}

	category, err := h.categoriesService.GetProductCategory(ctx, categoryID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category by ID")
		return
	}

	categoryDTO := dtoFromDomain(category)
	responseHandler.JSONResponse(categoryDTO, http.StatusOK)
}

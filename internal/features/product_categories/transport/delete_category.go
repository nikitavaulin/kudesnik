package product_categories_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

func (h *ProductCategoryHTTPHandler) DeleteProductCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke delete product categories handler")

	categoryCode, err := core_http_request.GetCategoryCodeFromPath(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category_code")
		return
	}

	if err := h.categoriesService.DeleteProductCategory(ctx, categoryCode); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete category")
		return
	}

	responseHandler.NoContentResponse()
}

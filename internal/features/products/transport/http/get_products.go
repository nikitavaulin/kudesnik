package products_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type GetProductsResponseDTO []ProductDTOResponse

func (h *ProductsHTTPHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get products handler")

	categoryID, err := core_http_request.GetUUIDQueryParam(r, "category_id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get categoryID param")
		return
	}

	limit, offset, err := core_http_request.GetLimitOffsetParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset params")
		return
	}

	products, err := h.productsService.GetProducts(ctx, categoryID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get products")
		return
	}

	// categoriesDTO := GetProductsResponseDTO(productsDtoFromDomain(products...))
	responseHandler.JSONResponse(products, http.StatusOK)
}

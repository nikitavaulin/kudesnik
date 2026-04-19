package product_categories_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
	core_http_utils "github.com/nikitavaulin/kudesnik/internal/core/transport/http/utils"
)

type GetCategoriesResponse []ProductCategoryDTOResponse

func (h *ProductCategoryHTTPHandler) GetCategories(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get product categories handler")

	limit, offset, err := getLimitOffsetParams(r)
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

func getLimitOffsetParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get limit query param: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get offset query param: %w", err)
	}

	return limit, offset, nil
}

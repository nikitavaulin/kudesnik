package products_transport_http

import (
	"fmt"
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

// GetProducts godoc
// @Summary Получить список товаров
// @Description Получить список товаров, есть фильтры
// @Tags products
// @Accept json
// @Produce json
// @Param category query string false "код категории товар (windows, entrance-doors, interior-doors, balconies, others)"
// @Param limit query int false "лимит возвращаемых категорий"
// @Param offset query int false "смещение возвращаемых категорий"
// @Param order query string false "сортировка по возрастанию или убыванию (asc, desc)"
// @Success 200 {array} domain.ProductBaseDetailed "Список товаров"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /products [get]
func (h *ProductsHTTPHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get products handler")

	categoryCode, err := getCategoryCodeQueryParam(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category code param")
		return
	}

	limit, offset, err := core_http_request.GetLimitOffsetParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset params")
		return
	}

	order, err := core_http_request.GetOrderQueryParam(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get order param")
		return
	}

	products, err := h.productsService.GetProducts(ctx, categoryCode, order, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get products")
		return
	}

	responseHandler.JSONResponse(products, http.StatusOK)
}

func getCategoryCodeQueryParam(r *http.Request) (*domain.ProductCategoryCode, error) {
	value := r.URL.Query().Get("category")
	if value == "" {
		return nil, nil
	}
	if err := domain.ValidateCategoryCode(value); err != nil {
		return nil, fmt.Errorf("unknown category code: %w", core_errors.ErrInvalidArgument)
	}
	code := domain.ProductCategoryCode(value)
	return &code, nil
}

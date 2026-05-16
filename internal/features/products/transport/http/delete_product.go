package products_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

// DeleteProduct godoc
// @Summary Удалить товар
// @Description Удалить товар по ID.
// @Security BearerAuth
// @Tags products
// @Produce json
// @Param id path string true "ID товара" Format(uuid)
// @Success 204
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /products/{id} [delete]
func (h *ProductsHTTPHandler) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke delete product handler")

	productID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get productID")
		return
	}

	if err := h.productsService.DeleteProduct(ctx, productID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete product")
		return
	}

	responseHandler.NoContentResponse()
}

// DeleteProducts godoc
// @Summary Удалить товары
// @Description Удалить несколько товаров по списку их ID.
// @Security BearerAuth
// @Tags products
// @Produce json
// @Param request body []ProductIdDTORequest true "DeleteProducts тело запроса"
// @Success 204
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /products [delete]
func (h *ProductsHTTPHandler) DeleteProducts(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke delete products handler")

	var productIDsDTO []ProductIdDTORequest

	if err := core_http_request.DecodeRequest(r, &productIDsDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode HTTP request")
		return
	}

	productIDs := UUIDsFromDTOs(productIDsDTO...)

	if err := h.productsService.DeleteProducts(ctx, productIDs); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete products")
		return
	}

	responseHandler.NoContentResponse()
}

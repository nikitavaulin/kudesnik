package customer_requests_transport_http

import (
	"net/http"

	_ "github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

// GetCustomerRequest godoc
// @Summary Получить заявку
// @Description Получить заявку по ID
// @Security BearerAuth
// @Tags customer-requests
// @Produce json
// @Param id path string true "ID получаемой заявки" Format(uuid)
// @Success 200 {object} domain.CustomerRequestDetailed "полученная заявка"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /requests/{id} [get]
func (h *CustomerRequestsTransportHTTP) GetCustomerRequest(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get customer request handler")

	id, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get customer ID")
		return
	}

	customerRequest, err := h.requestsService.GetCustomerRequest(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get customer request")
		return
	}

	responseHandler.JSONResponse(customerRequest, http.StatusOK)
}

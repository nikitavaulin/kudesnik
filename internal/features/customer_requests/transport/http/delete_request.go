package customer_requests_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

// DeleteCustomerRequest godoc
// @Summary Удалить заявку
// @Description Удалить заявку по ID. Может только superadmin
// @Security BearerAuth
// @Tags customer-requests
// @Produce json
// @Param id path string true "ID удаляемой заявки" Format(uuid)
// @Success 204
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /requests/{id} [delete]
func (h *CustomerRequestsTransportHTTP) DeleteCustomerRequest(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke delete customer reques handler")

	requestID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get customer request ID")
		return
	}

	if err := h.requestsService.DeleteCustomerRequest(ctx, requestID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete customer request")
		return
	}

	responseHandler.NoContentResponse()
}

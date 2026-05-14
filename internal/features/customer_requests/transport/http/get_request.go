package customer_requests_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

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

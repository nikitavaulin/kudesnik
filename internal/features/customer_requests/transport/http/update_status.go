package customer_requests_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type UpdateReqStatusDTO struct {
	Status domain.CustomerRequestStatus `json:"status"`
}

func (h *CustomerRequestsTransportHTTP) UpdateCustomerRequestStatus(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke update customer request status handler")

	requestID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get request ID")
		return
	}

	adminID, ok := domain.UserIDFromContext(ctx)
	if !ok {
		responseHandler.ErrorResponse(err, "failed to get admin ID from context")
		return
	}

	var dto UpdateReqStatusDTO

	if err := core_http_request.DecodeRequest(r, &dto); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode HTTP request")
		return
	}

	if err := h.requestsService.UpdateCustomerRequestStatus(ctx, requestID, adminID, dto.Status); err != nil {
		responseHandler.ErrorResponse(err, "failed to update customer request status")
		return
	}

	responseHandler.NoContentResponse()
}

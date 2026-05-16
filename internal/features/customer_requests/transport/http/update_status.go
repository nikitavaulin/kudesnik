package customer_requests_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
	"go.uber.org/zap"
)

type UpdateReqStatusDTO struct {
	Status domain.CustomerRequestStatus `json:"status"`
}

// UpdateCustomerRequestStatus godoc
// @Summary Изменить статус заявки
// @Description Изменить статус заявки (new, in_progress, completed, cancelled) в системе
// @Security BearerAuth
// @Tags customer-requests
// @Accept json
// @Produce json
// @Param id path string true "ID обновляемой заявки" Format(uuid)
// @Param request body UpdateReqStatusDTO true "UpdateCustomerRequestStatus тело запроса"
// @Success 204
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /requests/status/{id} [patch]
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

	log.Debug(
		"Admin has changed customer request status",
		zap.String("Admin", adminID.String()),
		zap.String("CustomerRequest", requestID.String()),
		zap.String("New status", string(dto.Status)),
	)

	responseHandler.NoContentResponse()
}

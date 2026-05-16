package customer_requests_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
	"go.uber.org/zap"
)

// PatchCustomerRequest godoc
// @Summary Изменить заявку
// @Description Изменить заявку (желаемые дату и время, доп. комментарий) в системе
// @Security BearerAuth
// @Tags customer-requests
// @Accept json
// @Produce json
// @Param id path string true "ID обновляемой заявки" Format(uuid)
// @Param request body domain.CustomerRequestPatch true "PatchCustomerRequest тело запроса"
// @Success 204
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /requests/{id} [patch]
func (h *CustomerRequestsTransportHTTP) PatchCustomerRequest(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke patch customer request handler")

	adminID, ok := domain.UserIDFromContext(ctx)
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "failed to get admin ID from context")
		return
	}

	requestID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get customer request ID")
		return
	}

	var patch domain.CustomerRequestPatch

	if err := core_http_request.DecodeRequest(r, &patch); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode patch request")
		return
	}

	if err := h.requestsService.PatchCustomerRequest(ctx, requestID, patch); err != nil {
		responseHandler.ErrorResponse(err, "failed to patch customer request")
		return
	}

	log.Debug(
		"Admin has patched customer request",
		zap.String("Admin", adminID.String()),
		zap.String("Customer request", requestID.String()),
	)

	responseHandler.NoContentResponse()
}

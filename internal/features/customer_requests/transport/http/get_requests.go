package customer_requests_transport_http

import (
	"fmt"
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

func (h *CustomerRequestsTransportHTTP) GetCustomerRequests(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	status, err := getStatusQueryParam(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get customer request status from query param")
		return
	}

	phone, err := getPhoneNumberQueryParam(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get customer phone number from query param")
		return
	}

	adminID, err := core_http_request.GetUUIDQueryParam(r, "admin")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get adminID params")
		return
	}

	limit, offset, err := core_http_request.GetLimitOffsetParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset params")
		return
	}

	requests, err := h.requestsService.GetCustomerRequests(ctx, status, phone, adminID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get customer requests")
		return
	}

	responseHandler.JSONResponse(requests, http.StatusOK)
}

func getStatusQueryParam(r *http.Request) (*domain.CustomerRequestStatus, error) {
	value := core_http_request.GetStringParamOrNil(r, "status")
	if value == nil {
		return nil, nil
	}
	status := domain.CustomerRequestStatus(*value)
	if !domain.IsCustomerRequestStatus(status) {
		return nil, fmt.Errorf("unknown customer request status: %w", core_errors.ErrInvalidArgument)
	}
	return &status, nil
}

func getPhoneNumberQueryParam(r *http.Request) (*string, error) {
	value := core_http_request.GetStringParamOrNil(r, "customer")
	if value == nil {
		return nil, nil
	}
	phone := "+" + *value
	if err := core_validation.ValidatePhoneNumber(phone); err != nil {
		return nil, err
	}
	return &phone, nil
}

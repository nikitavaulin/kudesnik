package customer_requests_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

// CreateCustomerRequest godoc
// @Summary Создать заявку клиента
// @Description Создать заявку клиента (не требует авторизации)
// @Tags customer-requests
// @Accept json
// @Produce json
// @Param request body domain.CustomerRequest true "CreateCustomerRequest тело запроса"
// @Success 201 {object} CustomerRequestIDResponseDTO "ID созданной заявки"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /requests [post]
func (h *CustomerRequestsTransportHTTP) CreateCustomerRequest(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	var requestDTO CustomerRequestDTO

	if err := core_http_request.DecodeRequest(r, &requestDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode customer request create request")
		return
	}

	customer, err := domain.NewCustomer(requestDTO.CustomerPhoneNumber, requestDTO.CustomerFullname)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed create customer")
		return
	}

	customerRequest := domain.NewCustomerRequest(
		requestDTO.CustomerPhoneNumber,
		requestDTO.DesiredDate,
		requestDTO.DesiredTime,
		requestDTO.ExtraComment,
		requestDTO.ProductID,
	)

	id, err := h.requestsService.CreateCustomerRequest(ctx, *customerRequest, *customer)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create customer request")
		return
	}

	responseDTO := CustomerRequestIDResponseDTO{id.String()}
	responseHandler.JSONResponse(responseDTO, http.StatusOK)
}

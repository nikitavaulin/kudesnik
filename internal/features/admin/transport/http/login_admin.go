package admin_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type LoginAdminRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginAdminResponseDTO struct {
	Token string `json:"token"`
}

func (h *AdminTrasnsportHTTPHandler) LoginAdmin(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	var loginRequestDTO LoginAdminRequestDTO

	if err := core_http_request.DecodeRequest(r, &loginRequestDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode admin login request")
		return
	}

	token, err := h.adminService.LoginAdmin(
		ctx,
		loginRequestDTO.Email,
		loginRequestDTO.Password,
	)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to login admin")
		return
	}

	tokenDTO := LoginAdminResponseDTO{token}
	responseHandler.JSONResponse(tokenDTO, http.StatusOK)
}

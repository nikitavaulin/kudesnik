package admin_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type CreateAdminRequestDTO struct {
	Email     string      `json:"email"`
	FullName  string      `json:"full_name"`
	Password  string      `json:"password"`
	AdminType domain.Role `json:"admin_type"`
}

func (h *AdminTrasnsportHTTPHandler) CreateAdmin(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	var adminRequestDTO CreateAdminRequestDTO

	if err := core_http_request.DecodeRequest(r, &adminRequestDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode admin create request")
		return
	}

	admin, err := h.adminService.CreateAdmin(
		ctx,
		adminRequestDTO.Email,
		adminRequestDTO.FullName,
		adminRequestDTO.Password,
		adminRequestDTO.AdminType,
	)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create admin")
		return
	}

	adminDTO := toAdminResponseDTO(admin)
	responseHandler.JSONResponse(adminDTO, http.StatusOK)
}

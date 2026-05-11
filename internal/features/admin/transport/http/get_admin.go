package admin_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

func (h *AdminTrasnsportHTTPHandler) GetAdmin(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get admin handler")

	id, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get admin ID")
		return
	}

	admin, err := h.adminService.GetAdminByID(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get admin by ID")
		return
	}

	adminDTO := toAdminResponseDTO(admin)
	responseHandler.JSONResponse(adminDTO, http.StatusOK)
}

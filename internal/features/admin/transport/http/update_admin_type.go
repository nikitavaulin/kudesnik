package admin_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type UpdateAdminTypeRequestDTO struct {
	AdminType domain.Role `json:"admin_type"`
}

func (h *AdminTrasnsportHTTPHandler) UpdateAdminType(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke update admin type handler")

	adminID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get admin ID")
		return
	}

	var dto UpdateAdminTypeRequestDTO

	if err := core_http_request.DecodeRequest(r, &dto); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode HTTP request")
		return
	}

	if err := h.adminService.UpdateAdminType(ctx, adminID, dto.AdminType); err != nil {
		responseHandler.ErrorResponse(err, "failed to update admin type")
		return
	}

	responseHandler.NoContentResponse()
}

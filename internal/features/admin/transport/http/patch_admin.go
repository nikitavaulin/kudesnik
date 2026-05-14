package admin_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
	core_http_types "github.com/nikitavaulin/kudesnik/internal/core/transport/http/types"
)

type AdminPatchRequestDTO struct {
	Email     core_http_types.Nullable[string]      `json:"email"`
	FullName  core_http_types.Nullable[string]      `json:"full_name"`
	AdminType core_http_types.Nullable[domain.Role] `json:"admin_type"`
}

func (h *AdminTrasnsportHTTPHandler) PatchAdmin(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke patch admin handler")

	id, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get admin ID")
		return
	}

	var adminPatchDTO AdminPatchRequestDTO

	if err := core_http_request.DecodeRequest(r, &adminPatchDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode patch request")
		return
	}

	patch := patchDomainFromDTO(adminPatchDTO)

	patchedAdmin, err := h.adminService.PatchAdmin(ctx, id, patch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch admin")
		return
	}

	adminDTO := toAdminResponseDTO(patchedAdmin)
	responseHandler.JSONResponse(adminDTO, http.StatusOK)
}

func patchDomainFromDTO(patchDTO AdminPatchRequestDTO) domain.AdminPatch {
	return domain.NewAdminPatch(
		patchDTO.Email.ToDomain(),
		patchDTO.FullName.ToDomain(),
		patchDTO.AdminType.ToDomain(),
	)
}

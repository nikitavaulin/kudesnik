package admin_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

// GetAdmin godoc
// @Summary Получить админа
// @Description Получить админа по ID
// @Security BearerAuth
// @Tags admins
// @Produce json
// @Param id path string true "ID получаемого админа" Format(uuid)
// @Success 200 {object} AdminsRepsonseDTO "полученный admin"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /admins/{id} [get]
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

// GetAdminAfterAuth godoc
// @Summary Получить админа
// @Description Получить админа по ID из jwt после авторизации
// @Security BearerAuth
// @Tags admins
// @Produce json
// @Success 200 {object} AdminsRepsonseDTO "полученный admin"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /admins/profile [get]
func (h *AdminTrasnsportHTTPHandler) GetAdminAfterAuth(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get admin handler")

	id, ok := domain.UserIDFromContext(ctx)
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "failed to get admin ID from context")
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

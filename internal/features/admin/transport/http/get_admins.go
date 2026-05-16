package admin_transport_http

import (
	"fmt"
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

// GetAdmins godoc
// @Summary Получить список админов
// @Description Получить список админов. Доступен фильтр по типу админа (superadmin или manager)
// @Security BearerAuth
// @Tags admins
// @Produce json
// @Param admin_type query string false "тип получаемого админа (superadmin или manager)"
// @Success 200 {array} AdminsRepsonseDTO "полученный список админов"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /admins [get]
func (h *AdminTrasnsportHTTPHandler) GetAdmins(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get admins handler")

	adminType, err := getAdminTypeQueryParam(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get admin_type")
		return
	}

	admins, err := h.adminService.GetAdmins(ctx, adminType)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get admins")
		return
	}

	adminsDTO := toAdminsResponseDTO(admins)
	responseHandler.JSONResponse(adminsDTO, http.StatusOK)
}

func getAdminTypeQueryParam(r *http.Request) (*domain.Role, error) {
	value := core_http_request.GetStringParamOrNil(r, "admin_type")
	if value == nil {
		return nil, nil
	}
	role := domain.Role(*value)
	if !domain.IsRoleAdmin(role) {
		return nil, fmt.Errorf("unknown admin type: %w", core_errors.ErrInvalidArgument)
	}
	return &role, nil
}

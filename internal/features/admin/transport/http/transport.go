package admin_transport_http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_http_server "github.com/nikitavaulin/kudesnik/internal/core/transport/http/server"
)

type AdminTrasnsportHTTPHandler struct {
	adminService AdminService
}

type AdminService interface {
	CreateAdmin(ctx context.Context, email, fullname, password string, role domain.Role) (domain.Admin, error)
	LoginAdmin(ctx context.Context, email string, password string) (string, error)

	GetAdminByID(ctx context.Context, id uuid.UUID) (domain.Admin, error)
	GetAdmins(ctx context.Context, adminType *domain.Role) ([]domain.Admin, error)

	UpdateAdminType(ctx context.Context, id uuid.UUID, adminType domain.Role) error
	PatchAdmin(ctx context.Context, id uuid.UUID, patch domain.AdminPatch) (domain.Admin, error)
}

func NewAdminTrasnsportHTTPHandler(adminService AdminService) *AdminTrasnsportHTTPHandler {
	return &AdminTrasnsportHTTPHandler{
		adminService: adminService,
	}
}

func (h *AdminTrasnsportHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:       http.MethodPost,
			Path:         "/admins",
			Handler:      h.CreateAdmin,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole},
		},
		{
			Method:       http.MethodGet,
			Path:         "/admins/{id}",
			Handler:      h.GetAdmin,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole, domain.ManagerRole},
		},
		{
			Method:  http.MethodPost,
			Path:    "/admins/auth",
			Handler: h.LoginAdmin,
		},
		{
			Method:       http.MethodGet,
			Path:         "/admins",
			Handler:      h.GetAdmins,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole},
		},
		{
			Method:       http.MethodPatch,
			Path:         "/admins/role/{id}",
			Handler:      h.UpdateAdminType,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole},
		},
		{
			Method:       http.MethodPatch,
			Path:         "/admins/{id}",
			Handler:      h.PatchAdmin,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole},
		},
	}
}

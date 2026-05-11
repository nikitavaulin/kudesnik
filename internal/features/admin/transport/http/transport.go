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
	GetAdminByID(ctx context.Context, id uuid.UUID) (domain.Admin, error)
	LoginAdmin(ctx context.Context, email string, password string) (string, error)
}

func NewAdminTrasnsportHTTPHandler(adminService AdminService) *AdminTrasnsportHTTPHandler {
	return &AdminTrasnsportHTTPHandler{
		adminService: adminService,
	}
}

func (h *AdminTrasnsportHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/admins",
			Handler: h.CreateAdmin,
		},
		{
			Method:  http.MethodGet,
			Path:    "/admins/{id}",
			Handler: h.GetAdmin,
		},
		{
			Method:  http.MethodPost,
			Path:    "/admins/auth",
			Handler: h.LoginAdmin,
		},
	}
}

package admin_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	tools_jwt "github.com/nikitavaulin/kudesnik/internal/core/tools/jwt"
)

type AdminService struct {
	adminRepo   AdminRepository
	jwtProvider *tools_jwt.JwtProvider
}

type AdminRepository interface {
	CreateAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)

	GetAdminByID(ctx context.Context, id uuid.UUID) (domain.Admin, error)
	GetAdminByEmail(ctx context.Context, email string) (domain.Admin, error)
	GetAdmins(ctx context.Context, adminType *domain.Role) ([]domain.Admin, error)

	UpdateAdminType(ctx context.Context, id uuid.UUID, adminType domain.Role) error
	PatchAdmin(ctx context.Context, id uuid.UUID, patchedAdmin domain.Admin) (domain.Admin, error)
}

func NewAdminServie(adminRepo AdminRepository, jwtProvider *tools_jwt.JwtProvider) *AdminService {
	return &AdminService{
		adminRepo:   adminRepo,
		jwtProvider: jwtProvider,
	}
}

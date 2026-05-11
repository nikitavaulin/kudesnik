package admin_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	tools_passwordhasher "github.com/nikitavaulin/kudesnik/internal/core/tools/passwordhasher"
)

func (s *AdminService) CreateAdmin(ctx context.Context, email, fullname, password string, role domain.Role) (domain.Admin, error) {
	if err := domain.ValidateAdmin(email, fullname, password, role); err != nil {
		return domain.Admin{}, fmt.Errorf("invalid admin params: %w", err)
	}

	passwordHash, err := tools_passwordhasher.HashPassword(password)
	if err != nil {
		return domain.Admin{}, fmt.Errorf("failed to hash password: %w", err)
	}

	admin := *domain.NewAdmin(email, fullname, passwordHash, role)

	admin, err = s.adminRepo.CreateAdmin(ctx, admin)
	if err != nil {
		return domain.Admin{}, fmt.Errorf("failed to create admin in repo: %w", err)
	}

	return admin, err
}

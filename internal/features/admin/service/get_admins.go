package admin_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *AdminService) GetAdmins(ctx context.Context, adminType *domain.Role) ([]domain.Admin, error) {
	admins, err := s.adminRepo.GetAdmins(ctx, adminType)
	if err != nil {
		return []domain.Admin{}, fmt.Errorf("failed to get admins from repo: %w", err)
	}
	return admins, nil
}

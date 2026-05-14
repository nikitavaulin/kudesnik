package admin_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *AdminService) GetAdminByID(ctx context.Context, id uuid.UUID) (domain.Admin, error) {
	admin, err := s.adminRepo.GetAdminByID(ctx, id)
	if err != nil {
		return domain.Admin{}, fmt.Errorf("failed to get admin from repo: %w", err)
	}
	return admin, nil
}

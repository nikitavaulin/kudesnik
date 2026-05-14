package admin_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (s *AdminService) UpdateAdminType(ctx context.Context, id uuid.UUID, adminType domain.Role) error {
	if !domain.IsRoleAdmin(adminType) {
		return fmt.Errorf("unknown admin type: %w", core_errors.ErrInvalidArgument)
	}

	if err := s.adminRepo.UpdateAdminType(ctx, id, adminType); err != nil {
		return fmt.Errorf("failed to update in repo: %w", err)
	}

	return nil
}

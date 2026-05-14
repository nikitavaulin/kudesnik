package admin_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *AdminService) PatchAdmin(ctx context.Context, id uuid.UUID, patch domain.AdminPatch) (domain.Admin, error) {
	admin, err := s.GetAdminByID(ctx, id)
	if err != nil {
		return domain.Admin{}, fmt.Errorf("failed to get admin from repo: %w", err)
	}

	if err := patch.Validate(); err != nil {
		return domain.Admin{}, fmt.Errorf("invalid patch: %w", err)
	}

	if err := admin.ApplyPatch(patch); err != nil {
		return domain.Admin{}, fmt.Errorf("failed to apply patch: %w", err)
	}

	patched, err := s.adminRepo.PatchAdmin(ctx, id, admin)
	if err != nil {
		return domain.Admin{}, fmt.Errorf("failed to update admin in repo: %w", err)
	}

	return patched, nil
}

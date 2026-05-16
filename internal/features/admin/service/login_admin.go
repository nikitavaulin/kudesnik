package admin_service

import (
	"context"
	"fmt"

	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	tools_jwt "github.com/nikitavaulin/kudesnik/internal/core/tools/jwt"
	tools_passwordhasher "github.com/nikitavaulin/kudesnik/internal/core/tools/passwordhasher"
)

func (s *AdminService) LoginAdmin(ctx context.Context, email string, password string) (string, error) {
	admin, err := s.adminRepo.GetAdminByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("failed to get admin by email from repo: %w", err)
	}

	if !tools_passwordhasher.VerifyPassword(password, admin.PasswordHash) {
		return "", fmt.Errorf("incorrect password: %w", core_errors.ErrUnauthorized)
	}

	token, err := s.jwtProvider.GenerateToken(tools_jwt.NewClaims(admin))
	if err != nil {
		return "", fmt.Errorf("failed to generate jwt token: %w", err)
	}

	return token, err
}

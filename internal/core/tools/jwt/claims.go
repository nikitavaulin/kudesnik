package tools_jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type Claims struct {
	jwt.RegisteredClaims
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewClaims(admin domain.Admin) *Claims {
	return &Claims{
		ID:    admin.ID.String(),
		Email: admin.Email,
		Role:  string(admin.AdminType),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

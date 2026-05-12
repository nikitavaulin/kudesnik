package tools_jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JwtProvider struct {
	secret []byte
}

func NewJWTProvider(config Config) *JwtProvider {
	return &JwtProvider{
		secret: config.Secret,
	}
}

func (j *JwtProvider) GenerateToken(claims *Claims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := jwtToken.SignedString(j.secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign jwt token: %w", err)
	}

	return signedToken, nil
}

func (j *JwtProvider) DecodeClaims(token string) (*Claims, error) {
	claims := &Claims{}

	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt token: %w", err)
	}

	if !jwtToken.Valid {
		return nil, fmt.Errorf("invalid jwt token")
	}

	return claims, nil
}

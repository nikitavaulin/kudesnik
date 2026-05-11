package domain

import (
	"context"

	"github.com/google/uuid"
)

type emailsContextKey struct{}

var emailContextKey = emailsContextKey{}

func EmailToContext(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, emailContextKey, email)
}

func EmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(emailContextKey).(string)
	return email, ok
}

type rolesContextKey struct{}

var roleContextKey = rolesContextKey{}

func RoleToContext(ctx context.Context, role Role) context.Context {
	return context.WithValue(ctx, roleContextKey, role)
}

func RoleFromContext(ctx context.Context) (Role, bool) {
	role, ok := ctx.Value(roleContextKey).(Role)
	return role, ok
}

type usersIDContextKey struct{}

var userIDContextKey = usersIDContextKey{}

func UserIDToContext(ctx context.Context, id uuid.UUID) context.Context {
	return context.WithValue(ctx, userIDContextKey, id)
}

func UserIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	id, ok := ctx.Value(userIDContextKey).(uuid.UUID)
	return id, ok
}

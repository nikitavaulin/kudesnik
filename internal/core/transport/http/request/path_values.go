package core_http_request

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func GetUUIDFromPath(r *http.Request, key string) (uuid.UUID, error) {
	value := r.PathValue(key)
	if value == "" {
		return uuid.Nil, fmt.Errorf(
			"ID value is empty (key=%s): %w",
			key,
			core_errors.ErrInvalidArgument,
		)
	}
	ID, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf(
			"failed to parse ID, got: %s (key=%s): %w",
			value,
			key,
			core_errors.ErrInvalidArgument,
		)
	}
	return ID, nil
}

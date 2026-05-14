package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func GetIntQueryParam(request *http.Request, key string) (*int, error) {
	value := request.URL.Query().Get(key)
	if value == "" {
		return nil, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf(
			"param: %s with key: %s could not be convert to integer: %v: %w",
			value, key,
			err, core_errors.ErrInvalidArgument,
		)
	}

	return &intValue, nil
}

func GetUUIDQueryParam(request *http.Request, key string) (*uuid.UUID, error) {
	value := request.URL.Query().Get(key)
	if value == "" {
		return nil, nil
	}

	ID, err := uuid.Parse(value)
	if err != nil {
		return nil, fmt.Errorf(
			"param: %s with key: %s could not be convert to UUID: %v: %w",
			value, key,
			err, core_errors.ErrInvalidArgument,
		)
	}

	return &ID, nil
}

func GetLimitOffsetParams(r *http.Request) (*int, *int, error) {
	limit, err := GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get limit query param: %w", err)
	}

	offset, err := GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get offset query param: %w", err)
	}

	return limit, offset, nil
}

func GetStringQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

// GetStringParamOrNil return string query param or nil by key
func GetStringParamOrNil(r *http.Request, key string) *string {
	value := GetStringQueryParam(r, key)
	if len(value) == 0 {
		return nil
	}
	return &value
}

func GetOrderQueryParam(r *http.Request) (*string, error) {
	value := GetStringQueryParam(r, "order")
	if value == "" {
		return nil, nil
	}
	if value == domain.AscendingOrder || value == domain.DescendingOrder {
		return &value, nil
	}
	return nil, fmt.Errorf("unknown order: %w", core_errors.ErrInvalidArgument)
}

package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

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

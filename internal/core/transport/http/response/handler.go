package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPReponseHandler struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPReponseHandler(log *core_logger.Logger, rw http.ResponseWriter) *HTTPReponseHandler {
	return &HTTPReponseHandler{
		log: log,
		rw:  rw,
	}
}

func (h *HTTPReponseHandler) JSONResponse(responseBody any, statusCode int) {
	h.rw.WriteHeader(statusCode)
	if err := json.NewEncoder(h.rw).Encode(responseBody); err != nil {
		h.log.Error("failded to write HTTP response", zap.Error(err))
	}
}

func (h *HTTPReponseHandler) ErrorResponse(err error, msg string) {
	var statusCode int
	var logFunc func(string, ...zap.Field)

	switch {
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug

	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn

	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn

	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))
	h.errorResponse(statusCode, err, msg)

}

func (h *HTTPReponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPReponseHandler) errorResponse(statusCode int, err error, msg string) {
	h.log.Error(msg, zap.Error(err))

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	h.JSONResponse(response, statusCode)
}

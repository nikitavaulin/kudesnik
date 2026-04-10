package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (h *HTTPReponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}

}

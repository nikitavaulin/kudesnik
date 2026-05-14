package images_transport_http

import (
	"context"
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_http_server "github.com/nikitavaulin/kudesnik/internal/core/transport/http/server"
)

type ImagesTransportHTTPHandler struct {
	imagesService ImagesService
}

type ImagesService interface {
	GetImage(ctx context.Context, imagePath string) (*domain.Image, error)
}

func NewImageTransportHTTPHandler(imagesService ImagesService) *ImagesTransportHTTPHandler {
	return &ImagesTransportHTTPHandler{
		imagesService: imagesService,
	}
}

func (h *ImagesTransportHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/static/",
			Handler: h.GetImage,
		},
	}
}

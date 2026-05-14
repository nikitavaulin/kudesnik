package products_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

func (h *ProductsHTTPHandler) DeleteProductImage(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke delete product image handler")

	productID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get productID")
		return
	}

	product, err := h.productsService.GetProductBase(ctx, productID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get product")
		return
	}

	if err := h.imageService.DeleteProductImages(ctx, *product.ImageURL, *product.ThumbnailURL); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete product image")
		return
	}

	patch := &domain.ProductBasePatch{}
	patch.ImageURL.Value = nil
	patch.ImageURL.Set = true
	patch.ThumbnailURL.Value = nil
	patch.ThumbnailURL.Set = true

	if _, err = h.productsService.PatchProduct(ctx, productID, patch); err != nil {
		responseHandler.ErrorResponse(err, "failed to patch product")
		return
	}

	responseHandler.NoContentResponse()
}

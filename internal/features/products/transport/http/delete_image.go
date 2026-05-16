package products_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

// DeleteProductImage godoc
// @Summary Удалить фото товара
// @Description Удалить фото товара по ID товара.
// @Security BearerAuth
// @Tags products
// @Produce json
// @Param id path string true "ID товара для удаления фото" Format(uuid)
// @Success 204
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /products/image/{id} [delete]
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

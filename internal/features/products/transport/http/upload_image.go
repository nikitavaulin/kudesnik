package products_transport_http

import (
	"fmt"
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type UploadProductImageResponseDTO struct {
	ImageURL     string `json:"image_url" example:"uploads/products/1778765484773106300_90c1e8db-4a4a-44c3-961a-6ad2c4b043f6.jpg"`
	ThumbnailURL string `json:"thumbnail_url" example:"uploads/products/1778765484773106300_90c1e8db-4a4a-44c3-961a-6ad2c4b043f6.jpg"`
}

// UploadProductImage godoc
// @Summary Загрузить фото товара
// @Description Загрузить фото товара по ID
// @Security BearerAuth
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "ID товара" Format(uuid)
// @Success 201 {object} UploadProductImageResponseDTO "UploadProductImage тело ответа"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /products/image/{id} [post]
func (h *ProductsHTTPHandler) UploadProductImage(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke upload product image handler")

	productID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get productID")
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		responseHandler.ErrorResponse(fmt.Errorf("%w: %w", err, core_errors.ErrInvalidArgument), "File too large or invalid form")
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		responseHandler.ErrorResponse(err, "no image file provided")
		return
	}
	defer file.Close()

	uploadResult, err := h.imageService.UploadImage(ctx, file, header)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to upload image")
		return
	}

	patch := &domain.ProductBasePatch{}
	patch.ImageURL.Value = &uploadResult.ImagePath
	patch.ImageURL.Set = true
	patch.ThumbnailURL.Value = &uploadResult.ImagePath
	patch.ThumbnailURL.Set = true

	if _, err = h.productsService.PatchProduct(ctx, productID, patch); err != nil {
		_ = h.imageService.DeleteProductImages(ctx, uploadResult.ImagePath, uploadResult.ThumbnailPath)
		responseHandler.ErrorResponse(err, "failed to update product")
		return
	}

	responseDTO := UploadProductImageResponseDTO{
		ImageURL:     uploadResult.ImageURL,
		ThumbnailURL: uploadResult.ThumbnailURL,
	}

	responseHandler.JSONResponse(responseDTO, http.StatusCreated)
}

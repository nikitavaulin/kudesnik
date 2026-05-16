package images_transport_http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
	"go.uber.org/zap"
)

// GetImage godoc
// @Summary      Получить изображение
// @Description  Возвращает изображение по указанному пути. Поддерживает serve content.
// @Tags         images
// @Param        path   path      string  true  "Путь к изображению"  Format(path)
// @Header       200    {string}  Content-Type  "MIME тип изображения (image/jpeg, image/png)"
// @Header       200    {string}  Content-Length  "Размер изображения в байтах"
// @Success      200    {file}    binary  "Изображение успешно получено"
// @Failure      400    {object}  core_http_response.ErrorResponse "Неверный путь к изображению"
// @Failure      404    {object}  core_http_response.ErrorResponse  "Изображение не найдено"
// @Failure      500    {object}  core_http_response.ErrorResponse  "Внутренняя ошибка сервера"
// @Router       /static/{path} [get]
func (h *ImagesTransportHTTPHandler) GetImage(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get image handler")

	imgPath, err := getImagePathFromURL(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get img path")
		return
	}
	log.Warn(imgPath)

	image, err := h.imagesService.GetImage(ctx, imgPath)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get image")
		return
	}
	defer image.Content.Close()

	if file, ok := image.Content.(*os.File); ok {
		stat, _ := file.Stat()
		http.ServeContent(rw, r, imgPath, stat.ModTime(), file)
		return
	}

	rw.Header().Set("Content-Type", image.MimeType)
	rw.Header().Set("Content-Length", strconv.FormatInt(image.Size, 10))
	// TODO: cache control

	if _, err := io.Copy(rw, image.Content); err != nil {
		log.Error("failed to send image", zap.Error(err))
		return
	}
}

func getImagePathFromURL(r *http.Request) (string, error) {
	imgPath := strings.TrimPrefix(r.URL.Path, "/static/")
	if imgPath == "" {
		return "", fmt.Errorf("img path is empty: %w", core_errors.ErrInvalidArgument)
	}
	return imgPath, nil
}

package images_service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	"go.uber.org/zap"
)

func (s *ImageService) UploadImage(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*domain.ImageUploadResult, error) {
	log := core_logger.FromContext(ctx)

	if header.Size > 10<<20 {
		return nil, fmt.Errorf("invalid image size (exceeds 10 Mb): %w", core_errors.ErrInvalidArgument)
	}

	contentType := header.Header.Get("Content-Type")
	if err := validateImageType(contentType); err != nil {
		return nil, err
	}

	imagePath, err := s.storage.Save(ctx, file, header.Filename)
	if err != nil {
		return nil, fmt.Errorf("failed to save image: %w", err)
	}

	thumbnailPath, err := s.createThumbnail(ctx, file, header.Filename)
	if err != nil {
		log.Error("failed to create thumbnail", zap.Error(err))
	}
	if thumbnailPath == "" {
		thumbnailPath = imagePath
	}

	return &domain.ImageUploadResult{
		ImageURL:      s.storage.GetURL(imagePath),
		ImagePath:     imagePath,
		ThumbnailURL:  s.storage.GetURL(thumbnailPath),
		ThumbnailPath: thumbnailPath,
	}, nil
}

func validateImageType(contentType string) error {
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}
	if !allowedTypes[contentType] {
		return fmt.Errorf("only JPEG and PNG images allowed: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (s *ImageService) createThumbnail(ctx context.Context, file multipart.File, filename string) (string, error) {
	// if seeker, ok := file.(io.Seeker); ok {
	// 	seeker.Seek(0, 0)
	// }

	// img, err := imaging.Decode(file)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to decode img: %w", err)
	// }

	// thumbnail := imaging.Fill(img, 300, 300, imaging.Center, imaging.Lanczos)

	// extIdx := strings.LastIndex(filename, ".")
	// if extIdx == -1 {
	// 	extIdx = len(filename)
	// }

	// thumbName := filename[:extIdx] + "_thumb" + filename[extIdx:]
	return "", nil
}

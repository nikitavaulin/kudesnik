package images_service

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ImageService) ValidatePath(imagePath string) error {
	cleanPath := filepath.Clean(imagePath)

	if strings.Contains(cleanPath, "..") {
		return fmt.Errorf("invalid path: directory traversal detected")
	}

	// matched, _ := regexp.MatchString(`^[a-zA-Z0-9./_: -]+$`, cleanPath)
	// if !matched {
	// 	return fmt.Errorf("invalid path: contains illegal characters")
	// }

	return nil
}

func (s *ImageService) GetMimeType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".bmp":
		return "image/bmp"
	case ".svg":
		return "image/svg+xml"
	default:
		return "application/octet-stream"
	}
}

func (s *ImageService) GetImage(ctx context.Context, imagePath string) (*domain.Image, error) {
	if err := s.ValidatePath(imagePath); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	image, err := s.storage.Get(ctx, imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}

	image.MimeType = s.GetMimeType(imagePath)

	// image.CacheControl = "private, max-age=86400"
	return image, nil
}

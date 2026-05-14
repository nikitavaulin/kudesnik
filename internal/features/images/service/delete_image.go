package images_service

import (
	"context"
	"fmt"
)

func (s *ImageService) DeleteProductImages(ctx context.Context, imagePath, thumbnailPath string) error {
	if imagePath != "" {
		if err := s.storage.Delete(ctx, imagePath); err != nil {
			return fmt.Errorf("delete image: %w", err)
		}
	}

	if thumbnailPath != "" && thumbnailPath != imagePath {
		if err := s.storage.Delete(ctx, thumbnailPath); err != nil {
			return fmt.Errorf("delete thumbnail: %w", err)
		}
	}

	return nil
}

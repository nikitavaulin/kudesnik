package image_local_storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *LocalStorage) Get(ctx context.Context, imagePath string) (*domain.Image, error) {
	fullPath := filepath.Join(r.basePath, imagePath)

	file, err := os.Open(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("image not found: %w: %w", err, core_errors.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to open image: %w", err)
	}

	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	return &domain.Image{
		Path:    imagePath,
		Content: file,
		Size:    stat.Size(),
	}, nil
}

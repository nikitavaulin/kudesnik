package image_storage

import (
	"context"
	"io"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type ImageStorage interface {
	Save(ctx context.Context, file io.Reader, originalFilename string) (string, error)
	Delete(ctx context.Context, filePath string) error
	GetURL(path string) string
	Get(ctx context.Context, imagePath string) (*domain.Image, error)
}

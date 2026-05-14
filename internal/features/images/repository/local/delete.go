package image_local_storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

func (s *LocalStorage) Delete(ctx context.Context, filePath string) error {
	if filePath == "" {
		return nil
	}

	fullPath := filepath.Join(s.basePath, filePath)

	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

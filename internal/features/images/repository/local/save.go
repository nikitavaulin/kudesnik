package image_local_storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func (s *LocalStorage) Save(ctx context.Context, file io.Reader, originalFilename string) (string, error) {
	safeName := generateSafeName(originalFilename)

	fullDirPath := filepath.Join(s.basePath, s.uploadPath)

	if err := os.MkdirAll(fullDirPath, 0755); err != nil {
		return "", fmt.Errorf("failed to create uploads directory: %w", err)
	}

	fullPath := filepath.Join(fullDirPath, safeName)

	dest, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dest.Close()

	if _, err := io.Copy(dest, file); err != nil {
		return "", fmt.Errorf("failed to copy file: %w", err)
	}

	relativePath := filepath.Join(s.uploadPath, safeName)
	relativePath = filepath.ToSlash(relativePath)

	return relativePath, nil
}

func generateSafeName(filename string) string {
	ext := filepath.Ext(filename)
	safeName := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), uuid.New().String(), ext)
	return safeName
}

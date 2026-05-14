package image_local_storage

import (
	"fmt"
	"path/filepath"
	"strings"
)

type LocalStorage struct {
	basePath   string
	baseURL    string
	uploadPath string
}

func NewLocalStorage(config *LocalStorageConfig) *LocalStorage {
	return &LocalStorage{
		basePath:   config.BasePath,
		baseURL:    config.BaseURL,
		uploadPath: config.UploadPath,
	}
}

func (s *LocalStorage) GetURL(path string) string {
	if path == "" {
		return ""
	}
	baseURL := strings.TrimRight(s.baseURL, "/")

	urlPath := filepath.ToSlash(path)
	urlPath = strings.TrimLeft(urlPath, "/")

	url := fmt.Sprintf("%s/%s", baseURL, urlPath)
	return url
}

package domain

import "io"

type ImageUploadResult struct {
	ImageURL  string
	ImagePath string

	ThumbnailURL  string
	ThumbnailPath string
}

type Image struct {
	// ID       string
	Path     string
	Content  io.ReadCloser
	Size     int64
	MimeType string
	// CacheControl string
}

package images_service

import (
	image_storage "github.com/nikitavaulin/kudesnik/internal/features/images/repository"
)

type ImageService struct {
	storage image_storage.ImageStorage
}

func NewImageService(storage image_storage.ImageStorage) *ImageService {
	return &ImageService{
		storage: storage,
	}
}

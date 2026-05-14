package image_local_storage

type LocalStorageConfig struct {
	BasePath   string
	BaseURL    string
	UploadPath string
}

func NewLocalStorageConfig(projectRoot string, baseURL string, uploadPath string) *LocalStorageConfig {
	return &LocalStorageConfig{
		BasePath:   projectRoot,
		BaseURL:    baseURL,
		UploadPath: uploadPath,
	}
}

package utils

func IsFileExtensionValid(fileExtension string) bool {
	allowedExtensions := []string{"jpg", "jpeg", "png", "pdf"}
	for _, extension := range allowedExtensions {
		if extension == fileExtension {
			return true
		}
	}
	return false
}

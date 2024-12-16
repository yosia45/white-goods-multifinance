package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"white-goods-multifinace/constants"
)

func GetFileTypeByExtension(file *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(file.Filename)
	ext = ext[1:]

	for fileType, extensions := range constants.FileExtensions {
		for _, validExt := range extensions {
			if ext == validExt {
				return fileType, nil
			}
		}
	}

	return "", fmt.Errorf("file type not recognized for extension %s", ext)
}

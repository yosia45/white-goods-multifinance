package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func SaveUploadFile(file *multipart.FileHeader, path string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%s-%d%s", uuid.New().String(), time.Now().Unix(), filepath.Ext(file.Filename))
	filePath := filepath.Join(path, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return filePath, nil
}

package utils

import (
	"fmt"
	"os"
)

func DeleteFile(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("no file to delete")
	}

	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

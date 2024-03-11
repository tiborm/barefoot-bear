package bfb_io

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func WriteFile(dirPath string, fileName string, content []byte) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirPath, 0o755)
		if errDir != nil {
			return fmt.Errorf("error creating directory: %w", errDir)
		}
	}

	err = os.WriteFile(
		filepath.Join(dirPath, fileName),
		content, 0o644,
	)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}
	return nil
}

func IsFileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else {
			return false, fmt.Errorf("error checking file in cache: %w", err)
		}
	}

	return true, nil
}

func GetFile(filePath string) ([]byte, error) {
	res, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading from file cache: %w", err)
	}

	return res, nil
}

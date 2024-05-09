package filehandler

import (
	"errors"
	"os"
	"path/filepath"
)

type FileHandler struct {}

func (fh FileHandler) GetFile(filePath string) ([]byte, error) {
	res, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (fh FileHandler) IsFileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (fh FileHandler) WriteFile(outputDirectory string, fileName string, content []byte) error {
	if err := createDirIfNotExists(outputDirectory); err != nil {
		return err
	}

	return writeFile(filepath.Join(outputDirectory, fileName), content)
}

func writeFile(filePath string, content []byte) error {
	return os.WriteFile(filePath, content, 0o644)
}

func createDirIfNotExists(outputDirectory string) error {
	if _, err := os.Stat(outputDirectory); err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(outputDirectory, 0o755)
		}
		return err
	}
	return nil
}

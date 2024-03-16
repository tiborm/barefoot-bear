package bfbio

import (
	"os"
	"path/filepath"
)

func WriteFile(outputDirectory string, fileName string, content []byte) error {
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
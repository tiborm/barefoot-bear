package bfbio

import (
	"errors"
	"os"
)

func GetFile(filePath string) ([]byte, error) {
	res, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func IsFileExists(filePath string) (bool, error) {
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
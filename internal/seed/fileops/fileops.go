package fileops

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/tiborm/barefoot-bear/internal/seed/jsonops"
)

func ReadFilenamesFromFolder(folderPath string) []string {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	var filenames []string
	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}

	return filenames
}

func ReadJsonFile[J jsonops.JsonResponseStruct](jsonFile string) (*J, error) {
	file, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s %v", jsonFile, err)
	}

	var result *J
	if err := json.Unmarshal(file, &result); err != nil {
		return nil, fmt.Errorf("failed to decode: %s %v", jsonFile, err)
	}

	return result, nil
}
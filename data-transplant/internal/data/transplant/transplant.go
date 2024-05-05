package transplant

import (
	"fmt"
	"path/filepath"

	"github.com/tiborm/barefoot-bear/internal/data/transplant/bfbio"
	"github.com/tiborm/barefoot-bear/internal/filters"
	"github.com/tiborm/barefoot-bear/internal/params"
)

func FetchAndStore(params params.FetchAndStoreParams, IDs []string) ([]string, error) {
	var fetchedBytes []byte
	var id string
	var iterations int
	allProductIDs := make([]string, 0)

	if IDs == nil {
		iterations = 1
	} else {
		iterations = len(IDs)
	}

	for i := 0; i < iterations; i++ {
		if IDs == nil {
			id = ""
		} else {
			id = IDs[i]
		}

		fileName := fmt.Sprintf("%s%s", id, params.StoreParams.FileNameExtension)
		filePath := filepath.Join(params.StoreParams.FolderPath, fileName)

		isCached, err := bfbio.IsFileExists(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to verify file in cache: %w", err)
		}

		if params.FetchParams.ForceFetch || !isCached {
			fetchedBytes, err := params.FetchFn(id, params.FetchParams)
			if err != nil {
				return nil, fmt.Errorf("error occurred while fetching: %w", err)
			}
			err = bfbio.WriteFile(params.StoreParams.FolderPath, fileName, fetchedBytes)
			if err != nil {
				return nil, fmt.Errorf("error writing to file: %w", err)
			}
		}

		if len(fetchedBytes) == 0 {
			fetchedBytes, err = readCategoriesFromFile(filePath)
			if err != nil {
				return nil, fmt.Errorf("error reading file: %w", err)
			}
		}

		extractedIDs, err := params.IDExtractorFn(fetchedBytes)
		if err != nil {
			return nil, fmt.Errorf("error extracting IDs: %w", err)
		}

		fmt.Println("Cleaning IDs")
		allProductIDs = append(allProductIDs, filters.ApplyAllCleaner(extractedIDs)...)
	}
	return allProductIDs, nil
}

func readCategoriesFromFile(file string) ([]byte, error) {
	categoriesByteArray, err := bfbio.GetFile(file)
	if err != nil {
		return nil, err
	}

	return categoriesByteArray, nil
}

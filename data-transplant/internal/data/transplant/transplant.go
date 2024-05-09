package transplant

import (
	"fmt"
	"path/filepath"

	"github.com/tiborm/barefoot-bear/internal/data/transplant/bfbio"
	"github.com/tiborm/barefoot-bear/internal/filters"
	"github.com/tiborm/barefoot-bear/internal/params"
)

func FetchAndStore(IDs []string, params params.FetchAndStoreParams) ([]string, error) {
	var fetchedBytes []byte
	var currentID string
	var iterations int
	allProductIDs := make([]string, 0)

	if IDs == nil {
		iterations = 1
	} else {
		iterations = len(IDs)
	}

	for i := 0; i < iterations; i++ {
		if IDs != nil {
			currentID = IDs[i]	
		}

		fileName := fmt.Sprintf("%s%s", currentID, params.StoreParams.FileNameExtension)
		filePath := filepath.Join(params.StoreParams.FolderPath, fileName)

		isCached, err := bfbio.IsFileExists(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to verify file in cache: %w", err)
		}

		if (params.FetchParams.ForceFetch || !isCached) && params.FetchFn != nil {
			fetchedBytes, err = params.FetchFn(currentID, params.FetchParams)
			if err != nil {
				return nil, fmt.Errorf("error occurred while fetching: %w", err)
			}
			err = bfbio.WriteFile(params.StoreParams.FolderPath, fileName, fetchedBytes)
			if err != nil {
				return nil, fmt.Errorf("error writing to file: %w", err)
			}
		}

		if len(fetchedBytes) == 0 && isCached {
			fetchedBytes, err = bfbio.GetFile(filePath)
			if err != nil {
				return nil, fmt.Errorf("error reading file: %w", err)
			}
		}

		var extractedIDs []string
		if params.IDExtractorFn != nil {
			extractedIDs, err = params.IDExtractorFn(fetchedBytes)
			if err != nil {
				return nil, fmt.Errorf("error extracting IDs: %w", err)
			}
			allProductIDs = append(allProductIDs, filters.ApplyAllCleaner(extractedIDs)...)
		}
	}

	return allProductIDs, nil
}


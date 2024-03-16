package category

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tiborm/barefoot-bear/internal/data/transplant/bfbio"
	"github.com/tiborm/barefoot-bear/internal/filters"
	"github.com/tiborm/barefoot-bear/internal/model"
)

func GetCategories(url string, outputDirectory string, fileName string, forceFetch bool) ([]string, error) {
	var catIds []string
	var categoryBytes []byte
	filePath := filepath.Join(outputDirectory, fileName)

	iscached, err := bfbio.IsFileExists(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to verify category file in cache: %w", err)
	}

	if forceFetch || !iscached {
		categoryBytes, err = fetchCategoriesFromURL(url) //
		if err != nil {
			return nil, fmt.Errorf("error occurred while fetching categories: %w", err)
		}
		err = saveCategoriesToFile(outputDirectory, fileName, categoryBytes)
		if err != nil {
			return nil, fmt.Errorf("error writing categories to file: %w", err)
		}
	}

	if len(categoryBytes) == 0 {
		categoryBytes, err = readCategoriesFromFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("error reading file: %w", err)
		}
	}

	catIds = *getCategoryIDs(categoryBytes)

	log.Println("Cleaning category IDs")
	return cleanUpIDs(catIds), nil
}

func cleanUpIDs(ids []string) []string {
	return filters.ApplyAllCleaner(ids)
}

func readCategoriesFromFile(file string) ([]byte, error) {
	categoriesByteArray, err := bfbio.GetFile(file)
	if err != nil {
		return nil, err
	}

	return categoriesByteArray, nil
}

func fetchCategoriesFromURL(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	categoriesByteArray, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return categoriesByteArray, nil
}

func getCategoryIDs(categoriesByteArray []byte) *[]string {
	var categories []model.Category
	json.Unmarshal(categoriesByteArray, &categories)

	log.Printf("Fetched %d main categories", len(categories))

	allCategories := getSubIDsInDepth(categories, &[]string{})

	log.Printf("Fetched %d categories in total, including sub-categories", len(*allCategories))
	return allCategories
}

func saveCategoriesToFile(outputDirectory string, fileName string, categories []byte) error {
	return bfbio.WriteFile(outputDirectory, fileName, categories)
}

// getSubIDsInDepth is a helper function to extract all sub-category ID
func getSubIDsInDepth(categories []model.Category, ids *[]string) *[]string {
	for _, category := range categories {
		*ids = append(*ids, category.ID)
		if category.Subs != nil {
			getSubIDsInDepth(category.Subs, ids)
		}
	}

	return ids
}

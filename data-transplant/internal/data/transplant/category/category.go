package category

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tiborm/barefoot-bear/constants"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/bfb_io"
	"github.com/tiborm/barefoot-bear/internal/model"
)

type Category interface {
	ReadCategoriesFromFile(file string) *[]string
	FetchCategoriesFromURL(url string) *[]string
}

func GetCategories(forceFetch bool) ([]string, error) {
	filePath := filepath.Join(constants.CategoryFolderPath, constants.CategoryFileName)

	isFileExist, err := bfb_io.IsFileExists(filePath)
	if err != nil {
		return nil, fmt.Errorf("error occurred while checking if the category file exists in the file cache: %w", err)
	}

	if forceFetch || !isFileExist {
		catIds, err := FetchCategoriesFromURL(constants.CategoryURL)
		if err != nil {
			return nil, fmt.Errorf("error occurred while fetching categories: %w", err)
		} else {
			return catIds, nil
		}
	} else {
		catIds, err := ReadCategoriesFromFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("error reading file: %w", err)
		} else {
			return catIds, nil
		}
	}
}

// ReadCategoriesFromFile reads categories from the given file
func ReadCategoriesFromFile(file string) ([]string, error) {
	categoriesByteArray, err := bfb_io.GetFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading cached category file: %w", err)
	}

	var categories []model.Category
	json.Unmarshal(categoriesByteArray, &categories)

	return *getSubsInDepth(categories, &[]string{}), nil
}

// FetchCategoriesFromURL fetches categories from the given URL
func FetchCategoriesFromURL(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching categories: %w", err)
	}

	categoriesByteArray, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("error reading categories response: %w", err)
	}

	bfb_io.WriteFile(constants.CategoryFolderPath, constants.CategoryFileName, categoriesByteArray)

	var categories []model.Category
	json.Unmarshal(categoriesByteArray, &categories)

	log.Printf("Fetched %d main categories", len(categories))

	allCategories := getSubsInDepth(categories, &[]string{})

	log.Printf("Fetched %d categories in total, including sub-categories", len(*allCategories))

	return *allCategories, nil
}

// getSubsInDepth is a helper function to extract all sub-categories from a category
func getSubsInDepth(categories []model.Category, ids *[]string) *[]string {
	for _, category := range categories {
		*ids = append(*ids, category.ID)
		if category.Subs != nil {
			getSubsInDepth(category.Subs, ids)
		}
	}

	return ids
}

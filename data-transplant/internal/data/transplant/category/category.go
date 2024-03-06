package category

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/tiborm/barefoot-bear/internal/model"
)

type Category interface {
	ReadCategoriesFromFile(file string) *[]string
	FetchCategoriesFromURL(url string) *[]string
}

// ReadCategoriesFromFile reads categories from the given file
func ReadCategoriesFromFile(file string) (*[]string, error) {
	categoriesJSONFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer categoriesJSONFile.Close()
	
	categoriesByteArray, err := io.ReadAll(categoriesJSONFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var categories []model.Category
	json.Unmarshal(categoriesByteArray, &categories)

	return getSubsInDepth(categories, &[]string{}), nil
}

// FetchCategoriesFromURL fetches categories from the given URL
func FetchCategoriesFromURL(url string) (*[]string, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error fetching categories: %w", err)
	}

	categoriesByteArray, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var categories []model.Category
	json.Unmarshal(categoriesByteArray, &categories)

	log.Printf("Fetched %d main categories", len(categories))

	allCategories := getSubsInDepth(categories, &[]string{})

	log.Printf("Fetched %d categories in total, including sub-categories", len(*allCategories))

	return allCategories, nil
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

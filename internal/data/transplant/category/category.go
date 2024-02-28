package category

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/tiborm/barefoot-bear/internal/model"
)

type CleanFunc func(categories *[]string) *[]string

type Category interface {
	ReadCategoriesFromFile(file string) *[]string
	FetchCategoriesFromURL(url string) *[]string
	CleanUpCategories(categories *[]string, criterias []CleanFunc) *[]string
}

// ReadCategoriesFromFile reads categories from the given file
func ReadCategoriesFromFile(file string) *[]string {
	categoriesJSONFile, err := os.Open(file)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	categoriesByteArray, _ := io.ReadAll(categoriesJSONFile)
	defer categoriesJSONFile.Close()

	var categories []model.Category
	json.Unmarshal(categoriesByteArray, &categories)

	return getSubsInDepth(categories, &[]string{})
}

// FetchCategoriesFromURL fetches categories from the given URL
func FetchCategoriesFromURL(url string) *[]string {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	categoriesByteArray, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var categories []model.Category
	json.Unmarshal(categoriesByteArray, &categories)

	log.Printf("Fetched %d main categories", len(categories))

	allCategories := getSubsInDepth(categories, &[]string{})

	log.Printf("Fetched %d categories in total, including sub-categories", len(*allCategories))

	return allCategories
}

func ApplyAllCleaner(categories *[]string) *[]string {
	cleaners := []CleanFunc{
		RemoveDulications, RemoveItemsWithSpecChars,
	}

	categories = CleanUpCategories(categories, cleaners)
	return categories
}

// CleanUpCategories applies the given criteria to the list of categories
func CleanUpCategories(categories *[]string, cleaners []CleanFunc) *[]string {
	// result := []string{}

	for _, clean := range cleaners {
		categories = clean(categories)
	}

	log.Printf("Cleaned up categories, %d remaining", len(*categories))
	
	// return &result
	return categories
}

// RemoveDulications removes duplicate items from the list
func RemoveDulications(categories *[]string) *[]string {
	seen := make(map[string]bool)
	result := []string{}

	for _, cat := range *categories {
		if _, ok := seen[cat]; !ok {
			seen[cat] = true
			result = append(result, cat)
		}
	}

	log.Printf("Removed %d duplicate categories", len(*categories)-len(result))

	return &result
}

// RemoveItemsWithSpecChars removes items from the list that contain special characters
func RemoveItemsWithSpecChars(categories *[]string) *[]string {
	result := []string{}
	
	pattern := "^[a-zA-Z0-9]+$"
	re, _ := regexp.Compile(pattern)

	for _, cat := range *categories {
		if re.Match([]byte(cat)) {
			result = append(result, cat)
		}
	}

	log.Printf("Removed %d categories with special characters", len(*categories)-len(result))

	return &result
}

// getSubsInDepth is a helper function to extract all sub-categories from a category
func getSubsInDepth(categories []model.Category, ids *[]string) *[]string {
	for _, category := range categories {
		*ids = append(*ids, category.Id)
		if category.Subs != nil {
			getSubsInDepth(category.Subs, ids)
		}
	}

	return ids
}

package filters // CleanUpCategories applies the given criteria to the list of categories

import (
	"log"
	"regexp"
)

type FilterFn func(categories *[]string) *[]string

func ApplyAllCleaner(categories *[]string) *[]string {
	cleaners := []FilterFn{
		RemoveDulications, RemoveItemsWithSpecChars,
	}

	categories = CleanUpCategories(categories, cleaners)
	return categories
}

// CleanUpCategories applies the given criteria to the list of categories
func CleanUpCategories(categories *[]string, cleaners []FilterFn) *[]string {
	for _, clean := range cleaners {
		categories = clean(categories)
	}

	log.Printf("Cleaned up categories, %d remaining", len(*categories))
	
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
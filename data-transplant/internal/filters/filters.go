package filters // CleanUpCategories applies the given criteria to the list of categories

import (
	"regexp"
)

type FilterFn func(IDs []string) []string

func ApplyAllCleaner(IDs []string) []string {
	cleaners := []FilterFn{
		RemoveDulications, RemoveItemsWithSpecChars,
	}

	IDs = CleanUpIDs(IDs, cleaners)
	return IDs
}

// CleanUpIDs applies the given criteria to the list of categories
func CleanUpIDs(IDs []string, cleaners []FilterFn) []string {
	for _, clean := range cleaners {
		IDs = clean(IDs)
	}

	return IDs
}

// RemoveDulications removes duplicate items from the list
func RemoveDulications(IDs []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, cat := range IDs {
		if _, ok := seen[cat]; !ok {
			seen[cat] = true
			result = append(result, cat)
		}
	}

	return result
}

// RemoveItemsWithSpecChars removes items from the list that contain special characters
func RemoveItemsWithSpecChars(IDs []string) []string {
	result := []string{}

	pattern := "^[a-zA-Z0-9]+$"
	re, _ := regexp.Compile(pattern)

	for _, cat := range IDs {
		if re.Match([]byte(cat)) {
			result = append(result, cat)
		}
	}

	return result
}

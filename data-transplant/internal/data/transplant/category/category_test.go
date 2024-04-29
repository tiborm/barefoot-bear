package category

import (
	"slices"
	"testing"

	"github.com/tiborm/barefoot-bear/internal/model"
)

func TestGetCategories(t *testing.T) {
	// Test data
	url := "http://localhost:8080/categories"
	outputDirectory := "test"
	fileName := "categories.json"
	forceFetch := false

	// Test fetching categories
	_, err := GetCategories(url, outputDirectory, fileName, forceFetch)

	if err != nil {
		t.Errorf("Failed to fetch categories: %s", err.Error())
	}

}

func TestExtractingSubCategories(t *testing.T) {
	// Test data
	category := model.Category{
		ID:   "1",
		Name: "Furniture",
		Subs: []model.Category{
			{
				ID:   "1.1",
				Name: "Chairs",
				Subs: nil,
			},
			{
				ID:   "1.2",
				Name: "Tables",
				Subs: []model.Category{
					{
						ID:   "1.3.1",
						Name: "Coffee Tables",
						Subs: []model.Category{
							{
								ID:   "1.4.1",
								Name: "Wooden Coffee Tables",
								Subs: nil,
							},
						},
					},
				},
			},
		},
	}

	// Test extracting sub-categories
	IDs := getSubIDsInDepth([]model.Category{category}, &[]string{})

	if len(*IDs) != 5 {
		t.Errorf("Failed to extract all sub-categories")
	}

	for _, ID := range []string{"1", "1.1", "1.2", "1.3.1", "1.4.1"} {
		if !slices.Contains(*IDs, ID) {
			t.Errorf("Failed to extract sub-category: %s", ID)
		}
	}
}

func TestCleaningUpIDs(t *testing.T) {
	testIDs := []string{"ID1", "ID2", "ID3", "ID3", "ID3", "ID///4", "ID....5", "ID&^6"}

	cleanedIDs := cleanUpIDs(testIDs)

	if len(cleanedIDs) != 3 {
		t.Errorf("Failed to remove duplicate IDs")
	}

	for i, ID := range(cleanedIDs) {
		if cleanedIDs[i] != testIDs[i] {
			t.Errorf("Failed to remove special characters from ID: %s", ID)
		}
	}
}
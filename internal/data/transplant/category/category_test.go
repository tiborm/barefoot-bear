package category

import (
	"slices"
	"testing"

	"github.com/tiborm/barefoot-bear/internal/model"
)

func TestExtractingSubCategories(t *testing.T) {
	// Test data
	category := model.Category{
		Id:   "1",
		Name: "Furniture",
		Subs: []model.Category{
			{
				Id:   "1.1",
				Name: "Chairs",
				Subs: nil,
			},
			{
				Id:   "1.2",
				Name: "Tables",
				Subs: []model.Category{
					{
						Id:   "1.3.1",
						Name: "Coffee Tables",
						Subs: []model.Category{
							{
								Id:   "1.4.1",
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
	ids := getSubsInDepth([]model.Category{category}, &[]string{})

	if len(*ids) != 5 {
		t.Errorf("Failed to extract all sub-categories")
	}

	for _, id := range []string{"1", "1.1", "1.2", "1.3.1", "1.4.1"} {
		if !slices.Contains(*ids, id) {
			t.Errorf("Failed to extract sub-category: %s", id)
		}
	}
}

func TestCleanUpCategories(t *testing.T) {
	// Test data
	categories := &[]string{"1", "1", "2", "2", "54303DA", "54303DA", "6542/432342", "6542\\432342", "6542:432342"}

	// Test cleaning up categories
	cleanedCategories := ApplyAllCleaner(categories)

	if len(*cleanedCategories) != 3 {
		t.Errorf("Failed to clean up categories: Expected no. of categories: 3, Actual no. of categories: %d", len(*cleanedCategories))
	}
	if !slices.Contains(*cleanedCategories, "1") {
		t.Errorf("Failed to remove duplicate: 1")
	}
	if !slices.Contains(*cleanedCategories, "2") {
		t.Errorf("Failed to remove duplicate: 2")
	}
	if !slices.Contains(*cleanedCategories, "54303DA") {
		t.Errorf("Failed to remove duplicate: 54303DA")
	}
}

func TestIdsShouldBeUnique(t *testing.T) {
	uniqueList := RemoveDulications(&[]string{"1", "1", "2", "2", "54303DA", "54303DA"})

	if len(*uniqueList) != 3 {
		t.Errorf("Failed to remove duplicates")
	}
	if !slices.Contains(*uniqueList, "1") {
		t.Errorf("Failed to remove duplicate: 1")
	}
	if !slices.Contains(*uniqueList, "2") {
		t.Errorf("Failed to remove duplicate: 2")
	}
	if !slices.Contains(*uniqueList, "54303DA") {
		t.Errorf("Failed to remove duplicate: 54303DA")
	}
}

func TestCategoryIdsShouldBeFileNameCompatibile(t *testing.T) {
	result := RemoveItemsWithSpecChars(&[]string{"1", "2", "54303DA", "6542/432342", "6542\\432342", "6542:432342", "6542*432342", "6542?432342", "6542\"432342", "6542<432342", "6542>432342", "6542|432342", "6542 432342", "6542\t432342", "6542\n432342", "6542\r"})

	if len(*result) != 3 {
		t.Errorf("Failed to remove special characters")
	}
	if !slices.Contains(*result, "1") {
		t.Errorf("Failed to remove duplicate: 1")
	}
	if !slices.Contains(*result, "2") {
		t.Errorf("Failed to remove duplicate: 2")
	}
	if !slices.Contains(*result, "54303DA") {
		t.Errorf("Failed to remove duplicate: 54303DA")
	}
}

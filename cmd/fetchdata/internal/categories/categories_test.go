package categories

import (
	"encoding/json"
	"slices"
	"testing"
)

// FIXME: generated test, please verify manually
func TestCategory(t *testing.T) {
	// Test data
	category := Category{
		Id:   "1",
		Name: "Furniture",
		Subs: []Category{
			{
				Id:   "1.1",
				Name: "Chairs",
				Subs: nil,
			},
			{
				Id:   "1.2",
				Name: "Tables",
				Subs: nil,
			},
		},
	}

	// Test JSON marshaling
	data, err := json.Marshal(category)
	if err != nil {
		t.Errorf("Failed to marshal category: %v", err)
	}

	// Test JSON unmarshaling
	var category2 Category
	err = json.Unmarshal(data, &category2)
	if err != nil {
		t.Errorf("Failed to unmarshal category: %v", err)
	}

	// Test equality
	// if category != category2 {
	// 	t.Errorf("Category does not match after marshaling and unmarshaling")
	// }
}

func TestExtractingSubCategories(t *testing.T) {
	// Test data
	category := Category{
		Id:   "1",
		Name: "Furniture",
		Subs: []Category{
			{
				Id:   "1.1",
				Name: "Chairs",
				Subs: nil,
			},
			{
				Id:   "1.2",
				Name: "Tables",
				Subs: []Category{
					{
						Id:  "1.3.1",
						Name: "Coffee Tables",
						Subs: []Category{
							{
								Id: "1.4.1",
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
	ids := getSubsInDepth([]Category{category}, &[]string{})

	if len(*ids) != 5 {
		t.Errorf("Failed to extract all sub-categories")
	}

	for _, id := range []string{"1", "1.1", "1.2", "1.3.1", "1.4.1"} {
		if !slices.Contains(*ids, id) {
			t.Errorf("Failed to extract sub-category: %s", id)
		}
	}
}

func TestIdsShouldBeUnique(t *testing.T) {
	t.Error("Not implemented")
}

func TestIdsShouldBeFileNameCompatibile(t *testing.T) {
	t.Error("Not implemented")
}

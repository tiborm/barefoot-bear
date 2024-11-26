package fetch

import (
	"slices"
	"testing"

	"github.com/tiborm/barefoot-bear/pkg/model"
)

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

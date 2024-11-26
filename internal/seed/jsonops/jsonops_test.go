package jsonops

import (
	"testing"

	"github.com/tiborm/barefoot-bear/internal/model"
)

func TestExtractCategories(t *testing.T) {
	categoriesJSON := &model.CategoryJsonResponse{
		{ID: "1", Name: "Category 1", Subs: []model.Category{
			{ID: "1-1", Name: "SubCategory 1-1", Subs: []model.Category{
				{ID: "1-1-1", Name: "SubSubCategory 1-1-1"},
				{ID: "1-1-2", Name: "SubSubCategory 1-1-2"},
			}},
			{ID: "1-2", Name: "SubCategory 1-2", Subs: []model.Category{
				{ID: "1-2-1", Name: "SubSubCategory 1-2-1"},
				{ID: "1-2-2", Name: "SubSubCategory 1-2-2"},
			}},
		}},
		{ID: "2", Name: "Category 2", Subs: []model.Category{
			{ID: "2-1", Name: "SubCategory 2-1", Subs: []model.Category{
				{ID: "2-1-1", Name: "SubSubCategory 2-1-1"},
				{ID: "2-1-2", Name: "SubSubCategory 2-1-2"},
			}},
			{ID: "2-2", Name: "SubCategory 2-2", Subs: []model.Category{
				{ID: "2-2-1", Name: "SubSubCategory 2-2-1"},
				{ID: "2-2-2", Name: "SubSubCategory 2-2-2"},
			}},
		}},
	}

	expected := []*model.Category{
		{ID: "1", Name: "Category 1", ChildIDs: []string{"1-1", "1-2"}},
		{ID: "2", Name: "Category 2", ChildIDs: []string{"2-1", "2-2"}},
		{ID: "1-1", Name: "SubCategory 1-1", ChildIDs: []string{"1-1-1", "1-1-2"}},
		{ID: "1-2", Name: "SubCategory 1-2", ChildIDs: []string{"1-2-1", "1-2-2"}},
		{ID: "2-1", Name: "SubCategory 2-1", ChildIDs: []string{"2-1-1", "2-1-2"}},
		{ID: "2-2", Name: "SubCategory 2-2", ChildIDs: []string{"2-2-1", "2-2-2"}},
		{ID: "1-1-1", Name: "SubSubCategory 1-1-1", ChildIDs: []string{}},
		{ID: "1-1-2", Name: "SubSubCategory 1-1-2", ChildIDs: []string{}},
		{ID: "1-2-1", Name: "SubSubCategory 1-2-1", ChildIDs: []string{}},
		{ID: "1-2-2", Name: "SubSubCategory 1-2-2", ChildIDs: []string{}},
		{ID: "2-1-1", Name: "SubSubCategory 2-1-1", ChildIDs: []string{}},
		{ID: "2-1-2", Name: "SubSubCategory 2-1-2", ChildIDs: []string{}},
		{ID: "2-2-1", Name: "SubSubCategory 2-2-1", ChildIDs: []string{}},
		{ID: "2-2-2", Name: "SubSubCategory 2-2-2", ChildIDs: []string{}},
	}

	result := ExtractCategories(categoriesJSON)

	// Check if the expected number of categories is returned
	if len(result) != len(expected) {
		t.Errorf("expected %d categories, got %d", len(expected), len(result))
	}

	// Check by id, if all expected categories are returned
	for _, category := range result {
		if !includes(expected, category.ID) {
			t.Errorf("expected category %v is missing from results", category.ID)
		}
	}

	// check childids for each category
	for _, category := range result {
		for _, childID := range category.ChildIDs {
			parentID := childID[:len(childID)-2]

			if !includes(result, childID) {
				t.Errorf("expected child category %v is missing from results", childID)
			}

			if !includes(expected, parentID) {
				if !includes(result, childID) {
					t.Errorf("expected parent category %v is missing from results", parentID)
				}
			}
		}
	}

	// all result subs filend should be emptied
	for _, category := range result {
		if category.Subs != nil {
			t.Errorf("expected category %v should have empty Subs field", category.ID)
		}
	}
}

func includes(categories []*model.Category, id string) bool {
	for _, category := range categories {
		if category.ID == id {
			return true
		}
	}
	return false
}

package main

import (
	"encoding/json"
	"testing"
)

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
	if category != category2 {
		t.Errorf("Category does not match after marshaling and unmarshaling")
	}
}

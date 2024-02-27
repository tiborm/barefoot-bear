package categories

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Category struct {
	Id   string     `json:"id"`
	Name string     `json:"name"`
	Subs []Category `json:"subs"`
}

func ReadCategoriesFromFile(file string) *[]string{
	categoriesJSONFile, err := os.Open("categories.json")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	categoriesByteArray, _ := io.ReadAll(categoriesJSONFile)
	defer categoriesJSONFile.Close()

	var categories []Category
	json.Unmarshal(categoriesByteArray, &categories)

	return getSubsInDepth(categories, &[]string{})
}

func FilterCategories() {
	
}

func getSubsInDepth(categories []Category, ids *[]string) *[]string {
	for _, category := range categories {
		fmt.Println(category.Id)
		*ids = append(*ids, category.Id)
		if category.Subs != nil {
			getSubsInDepth(category.Subs, ids)
		}
	}
	
	return ids
}
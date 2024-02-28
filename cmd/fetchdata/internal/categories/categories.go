package categories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Category struct {
	Id   string     `json:"id"`
	Name string     `json:"name"`
	Subs []Category `json:"subs"`
}

func ReadCategoriesFromFile(file string) *[]string{
	categoriesJSONFile, err := os.Open(file)

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
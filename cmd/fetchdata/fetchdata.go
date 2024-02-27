package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/tiborm/barefoot-bear/cmd/fetchdata/internal/searchtemplate"

	// "github.com/mpvl/unique"
)

type Category struct {
	Id   string     `json:"id"`
	Name string     `json:"name"`
	Subs []Category `json:"subs"`
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

func fetchProducsByCategory(categoryId string) {
	// Fetch products by category
	// var searchJsonMap map[string]json.RawMessage
	var searchJsonMap map[string]interface{}
	json.Unmarshal(searchtemplate.SearchJSONTemplate, &searchJsonMap)

	searchJsonMap["searchParameters"].(map[string]interface{})["input"] = categoryId

	payload, _ := json.Marshal(searchJsonMap)

	response, err := http.Post(
		"https://sik.search.blue.cdtapps.com/gb/en/search?c=listaf&v=20240110",
		"application/json",
		bytes.NewBuffer(payload),
	)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var responseJson map[string]interface{}
	json.Unmarshal(body, &responseJson)

	// TODO Fetch only if file yet not exists (state sync is not a concern)
	// TODO Add logging
	// TODO separate fetching and writing to file
	// TODO refactor path to constant, maybe in a separate file	
	err = os.WriteFile("./json-output/products/"+categoryId+".json", body, 0644)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func main() {
	categoriesJSONFile, err := os.Open("ikea-categories.json")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	categoriesByteArray, _ := io.ReadAll(categoriesJSONFile)
	defer categoriesJSONFile.Close()

	var categories []Category
	json.Unmarshal(categoriesByteArray, &categories)

	
	catIds := getSubsInDepth(categories, &[]string{})

	// fmt.Println(len(*catIds), " categories found")
	// unique.Strings(catIds)
	// fmt.Println(len(*catIds), " unique categories found")
	
	// fmt.Println(len(*catIds), " categories found")
	
	// unique.Strings(catIds)
	// Out of 1504 categories, 23 are unique? That's not right

	// TODO filer ids with / character

	for _, catId := range *catIds {
		time.Sleep(6 * time.Second) // Sleep for 6 seconds to avoid rate limiting
		fetchProducsByCategory(catId)
	}

	// TODO give some feedback to the user about the progress of the fetching process
	// like "Fetching products for category 1 of 100"
}

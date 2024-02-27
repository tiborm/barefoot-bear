package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Category struct {
	Id   string     `json:"id"`
	Name string     `json:"name"`
	Subs []Category `json:"subs"`
}

func getSubsInDepth(categories []Category, depth int, ids []string) []string {
	if depth > 0 {
		for _, category := range categories {
			fmt.Println(category.Id)
			if category.Subs == nil {
				ids = append(ids, category.Id)
			}
			getSubsInDepth(category.Subs, depth-1, ids)
		}
	} 
	return ids
}

func fetchProducsByCategory(categoryId string) {
	// Fetch products by category
	// var searchJsonMap map[string]json.RawMessage
	var searchJsonMap map[string]interface{}
	json.Unmarshal(SearchJSONTemplate, &searchJsonMap)

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

	err = os.WriteFile("./products/" + categoryId+".json", body, 0644)

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

	for _, category := range categories {
		fmt.Println(category.Name)
	}

	catIds := getSubsInDepth(categories, 7, []string{})

	for _, catId := range catIds {
		time.Sleep(6 * time.Second)
		fetchProducsByCategory(catId)
	}
}

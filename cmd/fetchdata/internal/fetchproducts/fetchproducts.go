package fetchproducts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tiborm/barefoot-bear/cmd/fetchdata/internal/searchtemplate"
)

func FetchProducsByCategory(categoryId string) {
	var searchJsonMap map[string]interface{}
	json.Unmarshal(searchtemplate.SearchJSONTemplate, &searchJsonMap)

	searchJsonMap["searchParameters"].(map[string]interface{})["input"] = categoryId

	payload, _ := json.Marshal(searchJsonMap)

	response, err := http.Post(
		searchtemplate.Url,
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
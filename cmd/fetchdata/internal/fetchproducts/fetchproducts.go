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

	// TODO Fetch only if file yet not exists (state sync is not a concern)
	// TODO Add logging
	// TODO separate fetching and writing to file
	// TODO refactor path to constant, maybe in a separate file	
	// TODO add folder path to config
	err = os.WriteFile("./json/products/"+categoryId+".json", body, 0644)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
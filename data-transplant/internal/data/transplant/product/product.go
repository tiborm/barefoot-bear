package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tiborm/barefoot-bear/internal/filters"
	"github.com/tiborm/barefoot-bear/internal/model"
	"github.com/tiborm/barefoot-bear/internal/params"
)

func FetchProductsFromAPI(id string, params params.FetchParams) ([]byte, error) {
	var searchJsonMap map[string]interface{}
	json.Unmarshal(params.PostPayload, &searchJsonMap)

	searchJsonMap["searchParameters"].(map[string]interface{})["input"] = id

	payload, _ := json.Marshal(searchJsonMap)

	res, err := http.Post(
		params.URL,
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching products by category: %w", err)
	}
	defer res.Body.Close()

	time.Sleep(time.Duration(params.FetchSleepTime) * time.Second)

	prodContent, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading fetched products: %w", err)
	}

	return prodContent, nil
}

func ExtractProductIds(rawJson []byte) ([]string, error) {
	var productsResponse model.ProductResponse
	err := json.Unmarshal(rawJson, &productsResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling products: %w", err)
	}

	if len(productsResponse.Results) == 0 {
		return []string{}, nil
	}

	products := productsResponse.Results[0].Items
	productsIDs := make([]string, len(products))
	for i, product := range products {
		productsIDs[i] = product.Product.Id
	}

	productsIDs = cleanUpIDs(productsIDs)

	return productsIDs, nil
}

func cleanUpIDs(ids []string) []string {
	return filters.ApplyAllCleaner(ids)
}

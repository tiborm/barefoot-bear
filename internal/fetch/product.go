package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tiborm/barefoot-bear/pkg/model"
)

type Sleeper interface {
	Sleep(time.Duration)
}

type ProductFetcher struct {
	sleeper Sleeper
	poster  Poster
}

type Poster interface {
	Post(url, contentType string, body io.Reader) (*http.Response, error)
}

func NewProductFetcher(sleeper Sleeper, poster Poster) *ProductFetcher {
	return &ProductFetcher{sleeper: sleeper, poster: poster}
}

func (pf ProductFetcher) Fetch(id string, params FetchParams) ([]byte, error) {
	var searchJsonMap map[string]interface{}
	json.Unmarshal(params.PostPayload, &searchJsonMap)

	searchJsonMap["searchParameters"].(map[string]interface{})["input"] = id

	payload, _ := json.Marshal(searchJsonMap)

	res, err := pf.poster.Post(
		params.URL,
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching products by category: %w", err)
	}
	defer res.Body.Close()

	pf.sleeper.Sleep(time.Duration(params.FetchSleepTime) * time.Second)

	prodContent, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading fetched products: %w", err)
	}

	return prodContent, nil
}

func (pf ProductFetcher) GetIDs(rawJson []byte) ([]string, error) {
	var productsResponse model.ProductJsonResponse
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

	return productsIDs, nil
}

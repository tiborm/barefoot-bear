package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/tiborm/barefoot-bear/internal/data/transplant/bfbio"
	"github.com/tiborm/barefoot-bear/internal/filters"
	"github.com/tiborm/barefoot-bear/internal/model"
	"github.com/tiborm/barefoot-bear/internal/params"
)

func GetAllProducts(catIds []string, params params.FetchAndStoreParams, forceFetch bool, fetchSleepTime float64) ([]string, error) {
	allProductIDs := make([]string, 0)
	
	for _, catId := range catIds {
		prodIDs, err := getProductsByCatID(
			catId,
			params,
			forceFetch,
		)
		if err != nil {
			return nil, fmt.Errorf("error getting products for category: %s, %w", catId, err)
		}
		allProductIDs = append(allProductIDs, prodIDs...)
	}

	log.Println("Cleaning product IDs")
	allProductIDs = CleanUpIDs(allProductIDs)

	return allProductIDs, nil
}

// getProductsByCatID fetches products by category from an API or reads them from a file if they are already cached.
// If the products for a category are not yet cached, or if force fetch is true, this function fetches them from the API and writes them to a file.
// It returns a list of product IDs in both scenarios.
func getProductsByCatID(
	categoryID string,
	params params.FetchAndStoreParams,
	forceFetch bool,
) ([]string, error) {
	fileName := categoryID + params.StoreParams.FileNameExtension
	filePath := filepath.Join(params.StoreParams.FolderPath, fileName)

	// Check if products of category are already cached
	isCached, err := bfbio.IsFileExists(filePath)
	if err != nil {
		return nil, fmt.Errorf("error checking products of category in file cache: %w", err)
	}

	var prodContent []byte
	// Geting products from cache if they are already cached and forceFetch is false
	if isCached && !forceFetch {
		prodContent, err = bfbio.GetFile(filePath)
		if err != nil {
			log.Println("Error reading products from cache, trying fetching from URL", err)
		}
	}
	// Fetching products from URL if they are not cached or forceFetch is true
	if prodContent == nil {
		if forceFetch {
			log.Println("Force fetching products from URL")
		}
		prodContent, err = FetchProductsFromAPI(categoryID, params.FetchParams)
		if err != nil {
			return nil, fmt.Errorf("error reading or fetching products: %w", err)
		}
	}

	productsIDs, err := ExtractProductIds(prodContent)
	if err != nil {
		return nil, fmt.Errorf("error extracting product IDs: %w", err)
	}

	if forceFetch || !isCached {
		err = bfbio.WriteFile(params.StoreParams.FolderPath, fileName, prodContent)
		if err != nil {
			return nil, fmt.Errorf("error writing products file: %w", err)
		}

		log.Println("Category ID: ", categoryID, " Products fetched and written to file: ", filepath.Join(params.StoreParams.FolderPath, fileName))
	}

	return productsIDs, nil
}

func CleanUpIDs(ids []string) []string {
	return filters.ApplyAllCleaner(ids)
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
	return productsIDs, nil
}

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

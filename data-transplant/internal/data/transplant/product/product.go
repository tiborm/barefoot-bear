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

	"github.com/tiborm/barefoot-bear/constants"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/bfb_io"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/searchtemplate"
	"github.com/tiborm/barefoot-bear/internal/model"
)

type (
	ProductsContent []byte
)

// GetProducts fetches products by category from an API or reads them from a file if they are already cached.
// If the products for a category are not yet cached, or if force fetch is true, this function fetches them from the API and writes them to a file.
// It returns a list of product IDs in both scenarios.
func GetProducts(categoryID string, forceFetch bool, fetchSleepTime float64) ([]string, error) {
	// FIXME you could use a generic fecth or read function for both categories and products, inventory too
	// function signature??: fetchOrRead(filePath string, fetchURL string, forceFetch bool) ([]byte, error)
	fileName := categoryID + constants.ProductsFileExtension
	filePath := filepath.Join(constants.ProductsFolderPath, fileName)

	pc, isCached, err := getProducts(categoryID, filePath, forceFetch, fetchSleepTime)
	if err != nil {
		return nil, err
	}

	// FIXME could we use a generic func to extract ids? a parameter function could return the id (path is a variable)
	// and the wraping function could be responsible for the interation or the recursion
	productsIDs, err := extractProductIds(pc)
	if err != nil {
		return nil, fmt.Errorf("error extracting product IDs: %w", err)
	}

	if !isCached || forceFetch {
		err = bfb_io.WriteFile(constants.ProductsFolderPath, fileName, pc)
		if err != nil {
			return nil, fmt.Errorf("error writing products file: %w", err)
		}

		log.Println("Category ID: ", categoryID, " Products fetched and written to file: ", filepath.Join(constants.ProductsFolderPath, fileName))
	}

	return productsIDs, nil
}

func extractProductIds(pc ProductsContent) ([]string, error) {
	var productsResponse model.ProductResponse
	err := json.Unmarshal(pc, &productsResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling products: %w", err)
	}

	products := productsResponse.Results[0].Items
	productsIDs := make([]string, len(products))
	for i, product := range products {
		productsIDs[i] = product.Product.Id
	}
	return productsIDs, nil
}

func getProducts(categoryId string, filePath string, forceFetch bool, fetchSleepTime float64) (ProductsContent, bool, error) {
	// Check if products of category are already cached
	isCached, err := bfb_io.IsFileExists(filePath)
	if err != nil {
		return nil, isCached, fmt.Errorf("error checking products of category in file cache: %w", err)
	}

	var prodContent ProductsContent
	// Geting products from cache if they are already cached and forceFetch is false
	if isCached && !forceFetch {
		prodContent, err = bfb_io.GetFile(filePath)
		if err != nil {
			fmt.Println("Error reading products from cache, trying fetching from URL", err)
		}
	}
	// Fetching products from URL if they are not cached or forceFetch is true
	if prodContent == nil {
		if forceFetch {
			fmt.Println("Force fetching products from URL")
		}
		prodContent, err = fetchProductsFromAPI(categoryId, fetchSleepTime)
	}

	if err != nil {
		return nil, isCached, fmt.Errorf("error reading or fetching products: %w", err)
	} else {
		return prodContent, isCached, nil
	}
}

func fetchProductsFromAPI(categoryId string, fetchSleepTime float64) (ProductsContent, error) {
	var searchJsonMap map[string]interface{}
	json.Unmarshal(searchtemplate.SearchJSONTemplate, &searchJsonMap)

	searchJsonMap["searchParameters"].(map[string]interface{})["input"] = categoryId

	payload, _ := json.Marshal(searchJsonMap)

	res, err := http.Post(
		constants.ProductSearchUrl,
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching products by category: %w", err)
	}

	time.Sleep(time.Duration(fetchSleepTime) * time.Second)

	prodContent, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading fetched products: %w", err)
	}

	return ProductsContent(prodContent), nil
}

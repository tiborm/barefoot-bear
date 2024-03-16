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
	"github.com/tiborm/barefoot-bear/internal/data/transplant/searchtemplate"
	"github.com/tiborm/barefoot-bear/internal/filters"
	"github.com/tiborm/barefoot-bear/internal/model"
)

type (
	productBytes []byte
	Product         struct{}
)

func GetAllProducts(catIds []string, outputDirectory string, fileExtension string, url string, forceFetch bool, fetchSleepTime float64) ([]string, error) {
	allProductIDs := make([]string, 0)
	for _, catId := range catIds {
		prodIDs, err := getProductsByCatID(catId, outputDirectory, fileExtension, url, forceFetch, fetchSleepTime)
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
func getProductsByCatID(categoryID string, outputDirectory string, fileExtension string, url string, forceFetch bool, fetchSleepTime float64) ([]string, error) {
	// FIXME you could use a generic fecth or read function for both categories and products, inventory too
	// function signature??: fetchOrRead(filePath string, fetchURL string, forceFetch bool) ([]byte, error)
	fileName := categoryID + fileExtension
	filePath := filepath.Join(outputDirectory, fileName)

	pc, isCached, err := getProducts(categoryID, filePath, url, forceFetch, fetchSleepTime)
	if err != nil {
		return nil, fmt.Errorf("error getting products by category: %w", err)
	}

	// FIXME could we use a generic func to extract ids? a parameter function could return the id (path is a variable)
	// and the wraping function could be responsible for the interation or the recursion
	productsIDs, err := extractProductIds(pc)
	if err != nil {
		return nil, fmt.Errorf("error extracting product IDs: %w", err)
	}

	if !isCached || forceFetch {
		err = bfbio.WriteFile(outputDirectory, fileName, pc)
		if err != nil {
			return nil, fmt.Errorf("error writing products file: %w", err)
		}

		log.Println("Category ID: ", categoryID, " Products fetched and written to file: ", filepath.Join(outputDirectory, fileName))
	}

	return productsIDs, nil
}

func CleanUpIDs(ids []string) []string {
	return filters.ApplyAllCleaner(ids)
}

func extractProductIds(pc productBytes) ([]string, error) {
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

func getProducts(categoryId string, filePath string, url string, forceFetch bool, fetchSleepTime float64) (productBytes, bool, error) {
	// Check if products of category are already cached
	isCached, err := bfbio.IsFileExists(filePath)
	if err != nil {
		return nil, isCached, fmt.Errorf("error checking products of category in file cache: %w", err)
	}

	var prodContent productBytes
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
		prodContent, err = fetchProductsFromAPI(categoryId, url, fetchSleepTime)
		if err != nil {
			return nil, isCached, fmt.Errorf("error reading or fetching products: %w", err)
		}
	}

	return prodContent, isCached, nil
}

func fetchProductsFromAPI(categoryId string, url string, fetchSleepTime float64) (productBytes, error) {
	var searchJsonMap map[string]interface{}
	json.Unmarshal(searchtemplate.SearchJSONTemplate, &searchJsonMap)

	searchJsonMap["searchParameters"].(map[string]interface{})["input"] = categoryId

	payload, _ := json.Marshal(searchJsonMap)

	res, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching products by category: %w", err)
	}
	defer res.Body.Close()

	time.Sleep(time.Duration(fetchSleepTime) * time.Second)

	prodContent, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading fetched products: %w", err)
	}

	return productBytes(prodContent), nil
}

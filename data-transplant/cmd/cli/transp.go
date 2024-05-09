package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/tiborm/barefoot-bear/constants"
	"github.com/tiborm/barefoot-bear/internal/data/transplant"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/category"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/inventory"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/product"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/searchtemplate"
	"github.com/tiborm/barefoot-bear/internal/params"
	"github.com/tiborm/barefoot-bear/internal/utils/config"
)

// fetchSleepTime and forceFetch are configurable parameters for the transplant operation.
var (
	categoryURL      string
	productSearchURL string
	inventoryURL     string
	forceFetch       bool
	fetchSleepTime   float64
	clientToken      string
)

type Transp struct {
	CategoryURL      string
	ProductSearchURL string
	InventoryURL     string
	ForceFetch       bool
	FetchSleepTime   float64
	ClientToken      string
}

// init initializes the categoryURL, productSearchURL, inventoryURL, forceFetch, fetchSleepTime and clientToken variables.
// It also sets the default values. The default values are read from the environment variables.
// If the environment variables are not set, the default values are used.
//
// The function also get the flag values if they are provided.
//
// The default values are:
// - categoryURL: CATEGORY_URL environment variable or "http://localhost:8080/categories"
// - productSearchURL: PRODUCT_SEARCH_URL environment variable or "http://localhost:8080/products/search"
// - inventoryURL: INVENTORY_URL environment variable or "http://localhost:8080/inventory/"
// - forceFetch: FORCE_FETCH environment variable or false
// - fetchSleepTime: FETCH_SLEEP_TIME environment variable or 5
// - clientToken: CLIENT_TOKEN environment variable or "client_token"
func InitTransp() *Transp {
	categoryURL = config.GetEnvAsString("CATEGORY_URL")
	productSearchURL = config.GetEnvAsString("PRODUCT_SEARCH_URL")
	inventoryURL = config.GetEnvAsString("INVENTORY_URL")
	forceFetchBool := config.GetEnvAsBool("FORCE_FETCH", constants.ForceFetch)
	sleepTimeInt := config.GetEnvAsFloat64("FETCH_SLEEP_TIME", constants.FetchSleepTime)
	clientToken = config.GetEnvAsString("CLIENT_TOKEN")

	flag.StringVar(&categoryURL, "categoryURL", categoryURL, "The URL to fetch the category data from. Environment variable: CATEGORY_URL")
	flag.StringVar(&productSearchURL, "productSearchURL", productSearchURL, "The URL to fetch the product data from. Environment variable: PRODUCT_SEARCH_URL")
	flag.StringVar(&inventoryURL, "inventoryURL", inventoryURL, "The URL to fetch the inventory data from. Environment variable: INVENTORY_URL")
	flag.StringVar(&clientToken, "clientToken", clientToken, "The client token to use for fetching inventory data. Environment variable: CLIENT_TOKEN")
	flag.BoolVar(&forceFetch, "forceFetch", forceFetchBool, "Whether to force fetch or not. Environment variable: FORCE_FETCH")
	flag.Float64Var(&fetchSleepTime, "fetchSleepTime", float64(sleepTimeInt), "The sleep time between fetches. Environment variable: FETCH_SLEEP_TIME")

	flag.Parse()

	return &Transp{
		CategoryURL:      categoryURL,
		ProductSearchURL: productSearchURL,
		InventoryURL:     inventoryURL,
		ForceFetch:       forceFetchBool,
		FetchSleepTime:   fetchSleepTime,
		ClientToken:      clientToken,
	}
}

func (t Transp) Validate() {
	// Validate URLs
	_, err := url.ParseRequestURI(t.CategoryURL)
	if err != nil {
		log.Fatalf("Invalid category URL: %v", err)
	}

	_, err = url.ParseRequestURI(t.ProductSearchURL)
	if err != nil {
		log.Fatalf("Invalid product search URL: %v", err)
	}

	_, err = url.ParseRequestURI(t.InventoryURL)
	if err != nil {
		log.Fatalf("Invalid inventory URL: %v", err)
	}

	// Validate client token
	if t.ClientToken == "" {
		log.Fatal("Client token is empty")
	}

	// Validate fetch sleep time
	if t.FetchSleepTime < 0 {
		log.Fatal("Fetch sleep time cannot be negative")
	}
}

func (t Transp) Run() {
	categoryParams := params.FetchAndStoreParams{
		StoreParams: params.StoreParams{
			FolderPath:        constants.CategoryFolderPath,
			FileNameExtension: constants.CategoryFileName,
		},
		FetchParams: params.FetchParams{
			URL:            categoryURL,
			ForceFetch:     forceFetch,
			FetchSleepTime: fetchSleepTime,
		},
		FetchFn:       category.FetchCategoriesFromAPI,
		IDExtractorFn: category.GetCategoryIDs,
	}

	ids, err := transplant.FetchAndStore(nil, categoryParams)
	if err != nil {
		log.Fatalf("Error during category transplant operation: %v", err)
	}
	log.Printf("Fetched %v categories.", len(ids))

	productsParams := params.FetchAndStoreParams{
		StoreParams: params.StoreParams{
			FolderPath:        constants.ProductsFolderPath,
			FileNameExtension: constants.ProductsFileExtension,
		},
		FetchParams: params.FetchParams{
			URL:            productSearchURL,
			PostPayload:    searchtemplate.SearchJSONTemplate,
			ForceFetch:     forceFetch,
			FetchSleepTime: fetchSleepTime,
		},
		FetchFn:       product.FetchProductsFromAPI,
		IDExtractorFn: product.ExtractProductIds,
	}

	ids, err = transplant.FetchAndStore(ids, productsParams)
	if err != nil {
		log.Fatalf("Error during product transplant operation: %v", err)
	}
	log.Printf("Fetched %v product data record", len(ids))

	inventoryParams := params.FetchAndStoreParams{
		StoreParams: params.StoreParams{
			FolderPath:        constants.InventoryFolderPath,
			FileNameExtension: constants.InventoryFileExtension,
		},
		FetchParams: params.FetchParams{
			URL:            inventoryURL,
			QueryParams:    constants.InventoryQueryParams,
			ForceFetch:     forceFetch,
			FetchSleepTime: fetchSleepTime,
			ClientToken:    clientToken,
		},
		FetchFn: inventory.FetchInventoriesFromAPI,
	}

	_, err = transplant.FetchAndStore(ids, inventoryParams)
	if err != nil {
		log.Fatalf("Error during inventory transplant operation: %v", err)
	}
	log.Printf("Fetched %v inventory data record.", len(ids))

	log.Println("Transplant operation completed successfully.")
}

func main() {
	transp := InitTransp()
	transp.Validate()
	transp.Run()
}

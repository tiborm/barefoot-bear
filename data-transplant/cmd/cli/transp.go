package main

import (
	"flag"
	"log"

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

// init initializes the fetchSleepTime and forceFetch variables from environment variables or command-line flags.
func init() {
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
}

func init() {
	flag.Parse()
}

func main() {
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

	ids, err := transplant.FetchAndStore(categoryParams, nil)

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

	ids, err = transplant.FetchAndStore(productsParams, ids)

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
		},
		FetchFn: inventory.FetchInventoriesFromAPI,
	}

	_, err = transplant.FetchAndStore(inventoryParams, ids)
	if err != nil {
		log.Fatalf("Error during transplant operation: %v", err)
	}

	log.Println("Transplant operation completed successfully.")
}

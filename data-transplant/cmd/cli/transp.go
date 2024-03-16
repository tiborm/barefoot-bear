package main

import (
	"flag"
	"log"

	"github.com/tiborm/barefoot-bear/constants"
	"github.com/tiborm/barefoot-bear/internal/data/transplant"
	"github.com/tiborm/barefoot-bear/internal/utils/config"
)

// fetchSleepTime and forceFetch are configurable parameters for the transplant operation.
var (
	categoryURL    string
	productSearchURL string
	inventoryURL   string
	forceFetch     bool
	fetchSleepTime float64
)

// init initializes the fetchSleepTime and forceFetch variables from environment variables or command-line flags.
func init() {
	forceFetchBool := config.GetEnvAsBool("FORCE_FETCH", constants.ForceFetch)
	sleepTimeInt := config.GetEnvAsFloat64("FETCH_SLEEP_TIME", constants.FetchSleepTime)
	categoryURL = config.GetEnvAsString("CATEGORY_URL")
	productSearchURL = config.GetEnvAsString("PRODUCT_SEARCH_URL")
	inventoryURL = config.GetEnvAsString("INVENTORY_URL")

	flag.StringVar(&categoryURL, "categoryURL", categoryURL, "The URL to fetch the category data from. Environment variable: CATEGORY_URL")
	flag.StringVar(&productSearchURL, "productSearchURL", productSearchURL, "The URL to fetch the product data from. Environment variable: PRODUCT_SEARCH_URL")
	flag.StringVar(&inventoryURL, "inventoryURL", inventoryURL, "The URL to fetch the inventory data from. Environment variable: INVENTORY_URL")
	flag.BoolVar(&forceFetch, "forceFetch", forceFetchBool, "Whether to force fetch or not. Environment variable: FORCE_FETCH")
	flag.Float64Var(&fetchSleepTime, "fetchSleepTime", float64(sleepTimeInt), "The sleep time between fetches. Environment variable: FETCH_SLEEP_TIME")
}

func main() {
	flag.Parse()

	err := transplant.StartDataTransplant(
		categoryURL,
		constants.CategoryFolderPath,
		constants.CategoryFileName,
		productSearchURL,
		constants.ProductsFolderPath,
		constants.ProductsFileExtension,
		inventoryURL,
		constants.InventoryFolderPath,
		constants.InventoryFileExtension,
		constants.InventoryQueryParams,
		forceFetch,
		fetchSleepTime,
	)
	if err != nil {
		log.Fatalf("Error during transplant operation: %v", err)
	}

	log.Println("Transplant operation completed successfully.")
}

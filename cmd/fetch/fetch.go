package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/tiborm/barefoot-bear/pkg/config"
	"github.com/tiborm/barefoot-bear/pkg/constants"
	"github.com/tiborm/barefoot-bear/internal/fetch"
)

// fetchSleepTime and forceFetch are configurable parameters for the fetch operation.
var (
	categoryURL      string
	productSearchURL string
	inventoryURL     string
	forceFetch       bool
	fetchSleepTime   float64
	clientToken      string
)

type Fetch struct {
	CategoryURL      string
	ProductSearchURL string
	InventoryURL     string
	ForceFetch       bool
	FetchSleepTime   float64
	ClientToken      string
}

type sleeper struct{}

func (s sleeper) Sleep(time.Duration) {
	time.Sleep(time.Duration(fetchSleepTime) * time.Second)
}

type poster struct{}

func (p poster) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return http.Post(url, contentType, body)
}

type getter struct{}

func (g getter) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
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
func InitFetch() *Fetch {
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

	return &Fetch{
		CategoryURL:      categoryURL,
		ProductSearchURL: productSearchURL,
		InventoryURL:     inventoryURL,
		ForceFetch:       forceFetchBool,
		FetchSleepTime:   fetchSleepTime,
		ClientToken:      clientToken,
	}
}

func (t Fetch) Validate() {
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

func (t Fetch) Run() {
	categoryParams := fetch.FetchAndStoreParams{
		StoreParams: fetch.StoreParams{
			FolderPath:        constants.CategoryFolderPath,
			FileNameExtension: constants.CategoryFileName,
		},
		FetchParams: fetch.FetchParams{
			URL:            categoryURL,
			ForceFetch:     forceFetch,
			FetchSleepTime: fetchSleepTime,
		},
		Fetcher: fetch.NewCategoryFetcher(getter{}),
	}

	fileHandler := fetch.FileHandlerService{}

	fetchService := fetch.NewFetchService(fileHandler)

	ids, err := fetchService.FetchAndStore(nil, categoryParams)
	if err != nil {
		log.Fatalf("Error during category fetch operation: %v", err)
	}
	log.Printf("Fetched %v categories.", len(ids))

	productsParams := fetch.FetchAndStoreParams{
		StoreParams: fetch.StoreParams{
			FolderPath:        constants.ProductsFolderPath,
			FileNameExtension: constants.ProductsFileExtension,
		},
		FetchParams: fetch.FetchParams{
			URL:            productSearchURL,
			PostPayload:    fetch.SearchJSONTemplate,
			ForceFetch:     forceFetch,
			FetchSleepTime: fetchSleepTime,
		},
		Fetcher: fetch.NewProductFetcher(sleeper{}, poster{}),
	}

	ids, err = fetchService.FetchAndStore(ids, productsParams)
	if err != nil {
		log.Fatalf("Error during product fetch operation: %v", err)
	}
	log.Printf("Fetched %v product data record", len(ids))

	inventoryParams := fetch.FetchAndStoreParams{
		StoreParams: fetch.StoreParams{
			FolderPath:        constants.InventoryFolderPath,
			FileNameExtension: constants.InventoryFileExtension,
		},
		FetchParams: fetch.FetchParams{
			URL:            inventoryURL,
			QueryParams:    constants.InventoryQueryParams,
			ForceFetch:     forceFetch,
			FetchSleepTime: fetchSleepTime,
			ClientToken:    clientToken,
		},
		Fetcher: fetch.NewInventoryFetcher(http.Client{}),
	}

	_, err = fetchService.FetchAndStore(ids, inventoryParams)
	if err != nil {
		log.Fatalf("Error during inventory fetch operation: %v", err)
	}
	log.Printf("Fetched %v inventory data record.", len(ids))

	log.Println("fetch operation completed successfully.")
}

func main() {
	transp := InitFetch()
	transp.Validate()
	transp.Run()
}

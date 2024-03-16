package transplant

import (
	"fmt"

	"github.com/tiborm/barefoot-bear/internal/data/transplant/category"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/inventory"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/product"
)

// StartDataTransplant orchestrates the fetching of the data from an unamed API.
// It fetches the categories, products and inventory data.
// If forceFetch is true, it will remove the file cache directory and fetch the data again.
// If forceFetch is false, it will fetch the data only if it is not already cached.
// fetchSleepTime is the time to wait between each fetch request.
// It returns an error if any of the fetching fails.
// It returns nil if the fetching is successful.
func StartDataTransplant(
	categoryURL string, 
	categoryFolderPath string,
	CategoryFileName string, 
	productSearchURL string,
	productsFolderPath string,
	productsFileExtension string,
	inventoryURL string,
	inventoryFolderPath string,
	inventoryFileExtension string,
	inventoryQueryParams string,
	forceFetch bool, 
	fetchSleepTime float64,
) error {
	// FIXME if forceFetch is true, empty the cache directory
	catIds,  err := category.GetCategories(
		categoryURL,
		categoryFolderPath,
		CategoryFileName,
		forceFetch,
	)
	if err != nil {
		return fmt.Errorf("failed to get categories: %w", err)
	}

	allProductIDs, err := product.GetAllProducts(
		catIds, 
		productsFolderPath, 
		productsFileExtension, 
		productSearchURL, 
		forceFetch, 
		fetchSleepTime,
	)
	if err != nil {
		return fmt.Errorf("failed to get all products: %w", err)
	}

	err = inventory.GetAllInventoryData(
		allProductIDs,
		inventoryFolderPath,
		inventoryFileExtension,
		inventoryURL,
		inventoryQueryParams,
		forceFetch,
		fetchSleepTime,
	)
	if err != nil {
		return fmt.Errorf("failed to get all inventory data: %w", err)
	}

	return nil
}

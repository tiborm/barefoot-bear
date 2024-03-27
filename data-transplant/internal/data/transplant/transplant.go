package transplant

import (
	"fmt"

	"github.com/tiborm/barefoot-bear/internal/data/transplant/category"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/inventory"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/product"
	"github.com/tiborm/barefoot-bear/internal/params"
)

// StartDataTransplant orchestrates the fetching of the data from an unamed API.
// It fetches the categories, products and inventory data.
// If forceFetch is true, it will remove the file cache directory and fetch the data again.
// If forceFetch is false, it will fetch the data only if it is not already cached.
// fetchSleepTime is the time to wait between each fetch request.
// It returns an error if any of the fetching fails.
// It returns nil if the fetching is successful.
func StartDataTransplant(params params.DataTransplantConfig, clientToken string) error {
	// FIXME if forceFetch is true, empty the cache directory
	catIds, err := category.GetCategories(
		params.Categories,
		params.ForceFetch,
		params.FetchSleepTime,
	)
	if err != nil {
		return fmt.Errorf("failed to get categories: %w", err)
	}

	allProductIDs, err := product.GetAllProducts(
		catIds,
		params.Products,
		params.ForceFetch,
		params.FetchSleepTime,
	)
	if err != nil {
		return fmt.Errorf("failed to get all products: %w", err)
	}

	err = inventory.GetAllInventoryData(
		allProductIDs,
		params.Inventory,
		clientToken,
		params.ForceFetch,
		params.FetchSleepTime,
	)
	if err != nil {
		return fmt.Errorf("failed to get all inventory data: %w", err)
	}

	return nil
}

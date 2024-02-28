package transplant

import (
	"time"

	"github.com/tiborm/barefoot-bear/internal/data/transplant/category"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/products"
)

func Transplant() {
	// TODO: move url to config
	catIds := category.FetchCategoriesFromURL("https://www.ikea.com/at/en/meta-data/navigation/catalog-products-slim.json")
	catIds = category.ApplyAllCleaner(catIds)

	for _, catId := range *catIds {
		// TODO: make the sleep time configurable
		time.Sleep(6 * time.Second) // Sleep for 6 seconds to avoid rate limiting
		fetchproducts.FetchProducsByCategory(catId)
	}

	// TODO give some feedback to the user about the progress of the fetching process
	// like "Fetching products for category 1 of 100"
}

package main

import (
	"time"

	"github.com/tiborm/barefoot-bear/internal/data/transplant/categories"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/products"
)



func main() {
	// TODO: move url to config
	catIds := categories.FetchCategoriesFromURL("https://www.ikea.com/at/en/meta-data/navigation/catalog-products-slim.json")

	// fmt.Println(len(*catIds), " categories found")
	// unique.Strings(catIds)
	// fmt.Println(len(*catIds), " unique categories found")
	
	// fmt.Println(len(*catIds), " categories found")
	
	// unique.Strings(catIds)
	// Out of 1504 categories, 23 are unique? That's not right

	// TODO filter ids with / character

	for _, catId := range *catIds {
		// TODO: make the sleep time configurable
		time.Sleep(6 * time.Second) // Sleep for 6 seconds to avoid rate limiting
		fetchproducts.FetchProducsByCategory(catId)
	}

	// TODO give some feedback to the user about the progress of the fetching process
	// like "Fetching products for category 1 of 100"
}

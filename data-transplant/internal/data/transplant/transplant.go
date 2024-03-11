package transplant

import (
	"fmt"
	"log"
	"os"

	"github.com/tiborm/barefoot-bear/internal/data/transplant/category"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/inventory"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/product"
	"github.com/tiborm/barefoot-bear/internal/filters"
)

func Transplant(fetchSleepTime float64, forceFetch bool) {
	// FIXME if forceFetch is true, remove file cache directory
	catIds, err := category.GetCategories(forceFetch)
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Cleaning category IDs")
	catIds = filters.ApplyAllCleaner(catIds)

	allProductIDs := make([]string, 0)
	for _, catId := range catIds {
		prodIDs, err := product.GetProducts(catId, forceFetch, fetchSleepTime)
		if err != nil {
			log.Println("Error: ", err)
			os.Exit(1)
		}
		allProductIDs = append(allProductIDs, prodIDs...)
	}

	fmt.Println("Cleaning product IDs")
	allProductIDs = filters.ApplyAllCleaner(allProductIDs)

	for _, prodId := range allProductIDs {
		err := inventory.FetchInventoryByProductID(prodId, fetchSleepTime)
		if err != nil {
			log.Println("Error: ", err)
			os.Exit(1)
		}
	}
	// TODO give some feedback to the user about the progress of the fetching process
	// like "Fetching products for category 1 of 100"
}

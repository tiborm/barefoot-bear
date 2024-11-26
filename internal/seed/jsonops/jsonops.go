package jsonops

import "github.com/tiborm/barefoot-bear/pkg/model"

type DataModel interface {
	model.Category | model.Product | model.Inventory
}

type JsonResponseStruct interface {
	model.CategoryJsonResponse | model.ProductJsonResponse | model.InventoryJsonResponse
}

type JsonFolderToCollection[J JsonResponseStruct, D DataModel] struct {
	JsonFolder         string
	Collection         string
	Extractor          func(*J) []*D
	IndexField         string
	JsonResponseStruct *J
	DataModel          *[]D
}

func ExtractCategories(categoriesJSON *model.CategoryJsonResponse) []*model.Category {
	var extract []*model.Category

	for _, c := range *categoriesJSON {
		extract = append(extract, &c)
	}

	return extract
}

func ExtractProducts(productsByCategoryJSONs *model.ProductJsonResponse) []*model.Product {
	var extract []*model.Product

	rs := productsByCategoryJSONs.Results
	for _, r := range rs {
		is := r.Items
		for _, pi := range is {
			extract = append(extract, &pi.Product)
		}
	}

	return extract
}

func ExtractInventory(inventoryByProductJSON *model.InventoryJsonResponse) []*model.Inventory {
	var inventoryData []*model.Inventory

	for _, data := range inventoryByProductJSON.Data {
		inventoryData = append(inventoryData, &data)
	}

	return inventoryData
}
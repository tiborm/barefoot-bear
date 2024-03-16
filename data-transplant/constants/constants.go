package constants

const (
	CategoryURL = "https://www.ikea.com/at/en/meta-data/navigation/catalog-products-slim.json"
	CategoryFolderPath = "./json-cache/"
	CategoryFileName = "categories.json"
	
	ProductSearchUrl = "https://sik.search.blue.cdtapps.com/at/en/search?c=listaf&v=20240110"
	ProductsFolderPath = "./json-cache/products/"
	ProductsFileExtension = ".products.json"

	InventoryFolderPath = "./json-cache/inventory/"
	InventoryFileExtension = ".inventory.json"
	InventoryURL = "https://api.ingka.ikea.com/cia/availabilities/ru/at?itemNos="
	InventoryQueryParams = "&expand=StoresList,Restocks,SalesLocations,DisplayLocations"
	
	FetchSleepTime = 1.0
	ForceFetch = false
)
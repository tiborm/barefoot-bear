package constants

const (
	CategoryFolderPath = "./json-cache/"
	CategoryFileName = "categories.json"
	
	ProductsFolderPath = "./json-cache/products/"
	ProductsFileExtension = ".products.json"

	InventoryFolderPath = "./json-cache/inventory/"
	InventoryFileExtension = ".inventory.json"
	InventoryQueryParams = "&expand=StoresList,Restocks,SalesLocations,DisplayLocations"
	
	FetchSleepTime = 1.0
	ForceFetch = false
)
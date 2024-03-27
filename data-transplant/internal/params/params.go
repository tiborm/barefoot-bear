package params

type FetchAndStoreConfig struct {
	FolderPath        string
	FileNameExtension string
	FetchURL          string
	QueryParams       string
}

type DataTransplantConfig struct {
	Products       FetchAndStoreConfig
	Categories     FetchAndStoreConfig
	Inventory      FetchAndStoreConfig
	ForceFetch     bool
	FetchSleepTime float64
}

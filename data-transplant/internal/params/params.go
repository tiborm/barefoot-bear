package params

type FetchAndStoreParams struct {
	FolderPath        string
	FileNameExtension string
	FetchURL          string
	PostPayload       []byte
	QueryParams       string
}

type DataTransplantParams struct {
	Categories     FetchAndStoreParams
	Products       FetchAndStoreParams
	Inventory      FetchAndStoreParams
	ForceFetch     bool
	FetchSleepTime float64
}

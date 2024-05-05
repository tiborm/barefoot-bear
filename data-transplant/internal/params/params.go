package params

type FetchParams struct {
	URL            string
	PostPayload    []byte
	QueryParams    string
	ClientToken    string
	ForceFetch     bool
	FetchSleepTime float64
}

type StoreParams struct {
	FolderPath        string
	FileNameExtension string
}
type FetchAndStoreParams struct {
	FetchParams   FetchParams
	StoreParams   StoreParams
	IDExtractorFn func([]byte) ([]string, error)
	FetchFn       func(id string, params FetchParams) ([]byte, error)
}

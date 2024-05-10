package transplant

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

type Fetcher interface {
	Fetch(id string, params FetchParams) ([]byte, error)
	GetIDs(rawJson []byte) ([]string, error)
}
type FetchAndStoreParams struct {
	FetchParams FetchParams
	StoreParams StoreParams
	Fetcher     Fetcher
}

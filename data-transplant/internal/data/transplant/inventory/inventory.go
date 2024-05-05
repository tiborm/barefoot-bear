package inventory

import (
	"io"
	"net/http"
	"time"

	"github.com/tiborm/barefoot-bear/internal/params"
)

func FetchInventoriesFromAPI(id string, params params.FetchParams) ([]byte, error) {
	fetchURL := params.URL + id + params.QueryParams
	req, err := http.NewRequest("GET", fetchURL, nil)
	if err != nil {
		return nil, err
	}

	// TODO: where can I get this header from? Move it to env var
	req.Header.Add("X-Client-Id", params.ClientToken)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	time.Sleep(time.Duration(params.FetchSleepTime) * time.Second)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

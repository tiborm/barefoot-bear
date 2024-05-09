package inventory

import (
	"fmt"
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

	req.Header.Add("X-Client-Id", params.ClientToken)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var body []byte
	switch {
		case response.StatusCode == http.StatusOK:
			// Sleep for a while before fetching the next inventory
			time.Sleep(time.Duration(params.FetchSleepTime) * time.Second)
			body, err = io.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}
		case response.StatusCode == http.StatusNotFound:
			return nil, fmt.Errorf("inventory not found for ID: %s", id)
		case response.StatusCode == http.StatusUnauthorized:
			return nil, fmt.Errorf("unauthorized access to inventory for ID: %s\nClient token expired or missing", id)
	}
	return body, nil
}

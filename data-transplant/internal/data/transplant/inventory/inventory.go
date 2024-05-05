package inventory

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/tiborm/barefoot-bear/internal/data/transplant/bfbio"
	"github.com/tiborm/barefoot-bear/internal/params"
)

func GetAllInventoryData(
	allProductIDs []string,
	params params.FetchAndStoreParams,
) error {
	for _, prodId := range allProductIDs {
		err := getInventoryByProductID(
			prodId,
			params,
		)
		if err != nil {
			log.Println("failed to get inventory data for prduct: %w, %w", prodId, err)
			return err
		}
		log.Println("Inventory data fetched and cached for product: ", prodId)
	}

	return nil
}

func getInventoryByProductID(
	productID string,
	params params.FetchAndStoreParams,
) error {
	fileName := productID + params.StoreParams.FileNameExtension
	filePath := filepath.Join(params.StoreParams.FolderPath, fileName)
	var inventoryBytes []byte

	isCached, err := bfbio.IsFileExists(filePath)
	if err != nil {
		return fmt.Errorf("failed to verify inventory file in cache: %w", err)
	}

	if params.FetchParams.ForceFetch || !isCached {
		inventoryBytes, err = FetchInventoriesFromAPI(productID, params.FetchParams)
		if err != nil {
			return fmt.Errorf("failed to fetch inventory data: %w", err)
		}
	}

	if len(inventoryBytes) == 0 {
		inventoryBytes, err = readInventoryFromFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read cached inventory file: %w", err)
		}
	}

	return bfbio.WriteFile(params.StoreParams.FolderPath, fileName, inventoryBytes)
}

func readInventoryFromFile(file string) ([]byte, error) {
	inventoryByteArray, err := bfbio.GetFile(file)
	if err != nil {
		return nil, err
	}

	return inventoryByteArray, nil
}

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

package inventory

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/tiborm/barefoot-bear/internal/data/transplant/bfbio"
)

func GetAllInventoryData(
	allProductIDs []string,
	outputDirectory string,
	fileExtension string,
	url string,
	queryParams string,
	forceFetch bool,
	fetchSleepTime float64,
) error {
	for _, prodId := range allProductIDs {
		err := fetchInventoryByProductID(
			prodId,
			url,
			fileExtension,
			outputDirectory,
			queryParams,
			forceFetch,
			fetchSleepTime,
		)
		if err != nil {
			return fmt.Errorf("failed to get inventory data for prduct: %s, %w", prodId, err)
		}
		log.Println("Inventory data fetched and cached for product: ", prodId)
	}

	return nil
}

func fetchInventoryByProductID(
	productID string,
	url string,
	outputDirectory string,
	fileExtension string,
	queryParams string,
	forceFetch bool,
	fetchSleepTime float64,
) error {
	fileName := productID + fileExtension
	filePath := filepath.Join(outputDirectory, fileName)
	fetchURL := url + productID + queryParams
	var inventoryBytes []byte

	isCached, err := bfbio.IsFileExists(filePath)
	if err != nil {
		return fmt.Errorf("failed to verify inventory file in cache: %w", err)
	}

	if forceFetch && !isCached {
		inventoryBytes, err = fetchInventoryDataByURL(fetchURL, fetchSleepTime)
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

	return bfbio.WriteFile(outputDirectory, fileName, inventoryBytes)
}

func readInventoryFromFile(file string) ([]byte, error) {
	inventoryByteArray, err := bfbio.GetFile(file)
	if err != nil {
		return nil, err
	}

	return inventoryByteArray, nil
}

func fetchInventoryDataByURL(fetchURL string, fetchSleepTime float64) ([]byte, error) {
	req, err := http.NewRequest("GET", fetchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// TODO: where can I get this header from?
	req.Header.Add("X-Client-Id", "b6c117e5-ae61-4ef5-b4cc-e0b1e37f0631")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch inventory: %w", err)
	}
	defer response.Body.Close()

	time.Sleep(time.Duration(fetchSleepTime) * time.Second)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read fetched inventory data: %w", err)
	}
	return body, nil
}

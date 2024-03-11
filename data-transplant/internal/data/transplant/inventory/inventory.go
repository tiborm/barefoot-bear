package inventory

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/tiborm/barefoot-bear/constants"
	"github.com/tiborm/barefoot-bear/internal/data/transplant/bfb_io"
)

// FIXME !!! Fetch doesn't work bc, of request key problem
func FetchInventoryByProductID(productID string, fetchSleepTime float64) error {
	fileName := productID + constants.InventoryFileExtension
	filePath := filepath.Join(constants.InventoryFolderPath, fileName)
	fetchUrl := constants.InventoryUrl + productID + constants.InventoryExpand

	response, err := http.Get(fetchUrl)
	if err != nil {
		return fmt.Errorf("error fetching inventory: %w", err)
	}
	time.Sleep(time.Duration(fetchSleepTime) * time.Second)

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return fmt.Errorf("error reading fetched inventory: %w", err)
	}

	// TODO Fetch only if file yet not exists (state sync is not a concern), force synd from config
	// TODO separate fetching and writing to file
	
	err = bfb_io.WriteFile(filePath, fileName, body)
	if err != nil {
		return fmt.Errorf("error writing inventory file: %w", err)
	}

	log.Println("Fetched and cached inventory data for product: ", productID)

	return nil
}

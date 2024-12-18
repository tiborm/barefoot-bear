package fetch

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/tiborm/barefoot-bear/pkg/model"
)

type Getter interface {
	Get(url string) (resp *http.Response, err error)
}

type CategoryFetcher struct {
	getter Getter
}

func NewCategoryFetcher(getter Getter) *CategoryFetcher {
	return &CategoryFetcher{getter: getter}
}

func (cf CategoryFetcher) Fetch(id string, params FetchParams) ([]byte, error) {
	response, err := http.Get(params.URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	categoriesByteArray, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return categoriesByteArray, nil
}

func (cf CategoryFetcher) GetIDs(rawJson []byte) ([]string, error) {
	var categories []model.Category
	if err := json.Unmarshal(rawJson, &categories); err != nil {
		return nil, err
	}

	log.Printf("Fetched %d main categories", len(categories))

	allCategories := *getSubIDsInDepth(categories, &[]string{})

	log.Printf("Fetched %d categories in total, including sub-categories", len(allCategories))
	return allCategories, nil
}

// getSubIDsInDepth is a helper function to extract all sub-category ID
func getSubIDsInDepth(categories []model.Category, ids *[]string) *[]string {
	for _, category := range categories {
		*ids = append(*ids, category.ID)
		if category.Subs != nil {
			getSubIDsInDepth(category.Subs, ids)
		}
	}

	return ids
}

package main

type CollectionJSON struct {
	ID   string           `json:"id"`
	Name string           `json:"name"`
	URL  string           `json:"url"`
	Im   *string          `json:"im,omitempty"` // Optional field
	Subs []CollectionJSON `json:"subs"`
}

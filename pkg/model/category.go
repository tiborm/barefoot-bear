package model

type Category struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    URL  string `json:"url"`
    IM   string `json:"im,omitempty"`
    Subs []Category `json:"subs,omitempty"`
}

type CategoryJsonResponse []Category
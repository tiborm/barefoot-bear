package model

type Category struct {
	Id   string     `json:"id"`
	Name string     `json:"name"`
	Subs []Category `json:"subs"`
}
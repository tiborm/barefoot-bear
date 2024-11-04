package main

type Price struct {
	CurrencyCode string  `json:"currencyCode"`
	Numeral      float64 `json:"numeral"`
	Current      struct {
		Prefix            string `json:"prefix"`
		WholeNumber       string `json:"wholeNumber"`
		Separator         string `json:"separator"`
		Decimals          string `json:"decimals"`
		Suffix            string `json:"suffix"`
		IsRegularCurrency bool   `json:"isRegularCurrency"`
	} `json:"current"`
	IsBreathTaking bool   `json:"isBreathTaking"`
	Discount       string `json:"discount"`
	Tag            string `json:"tag"`
	TagText        string `json:"tagText"`
	PriceText      string `json:"priceText"`
}

type Variant struct {
	ID                       string        `json:"id"`
	PipURL                   string        `json:"pipUrl"`
	ImageURL                 string        `json:"imageUrl"`
	ImageAlt                 string        `json:"imageAlt"`
	QuickFacts               []interface{} `json:"quickFacts"`
	Availability             []interface{} `json:"availability"`
	RatingValue              float64       `json:"ratingValue"`
	RatingCount              int           `json:"ratingCount"`
	Name                     string        `json:"name"`
	TypeName                 string        `json:"typeName"`
	ItemMeasureReferenceText string        `json:"itemMeasureReferenceText"`
	MainImageURL             string        `json:"mainImageUrl"`
	MainImageAlt             string        `json:"mainImageAlt"`
	ContextualImageURL       string        `json:"contextualImageUrl"`
	ItemNoGlobal             string        `json:"itemNoGlobal"`
	OnlineSellable           bool          `json:"onlineSellable"`
	LastChance               bool          `json:"lastChance"`
	ItemNo                   string        `json:"itemNo"`
	ItemType                 string        `json:"itemType"`
	SalesPrice               Price         `json:"salesPrice"`
}

type GPRDescription struct {
	NumberOfVariants int       `json:"numberOfVariants"`
	Variants         []Variant `json:"variants"`
}

type Product struct {
	Name                     string         `json:"name"`
	TypeName                 string         `json:"typeName"`
	ItemMeasureReferenceText string         `json:"itemMeasureReferenceText"`
	MainImageURL             string         `json:"mainImageUrl"`
	PipURL                   string         `json:"pipUrl"`
	ID                       string         `json:"id"`
	ItemNoGlobal             string         `json:"itemNoGlobal"`
	OnlineSellable           bool           `json:"onlineSellable"`
	LastChance               bool           `json:"lastChance"`
	GPRDescription           GPRDescription `json:"gprDescription"`
}

type ProductItem struct {
	Metadata string  `json:"metadata"`
	Product  Product `json:"product"`
}

type Component struct {
	Component string        `json:"component"`
	ViewMode  string        `json:"viewMode"`
	Filters   []interface{} `json:"filters"`
	Items     []ProductItem `json:"items"`
}

type ProductsJsonResult struct {
	UserGroup string      `json:"usergroup"`
	Results   []Component `json:"results"`
}

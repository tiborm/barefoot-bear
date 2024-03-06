package model

type Product struct {
    Metadata string `json:"metadata"`
    Product ProductDetails `json:"product"`
    Type string `json:"type"`
    Label string `json:"label"`
    ActionTokens []interface{} `json:"actionTokens"`
    IsBreakout bool `json:"isBreakout"`
}

type ProductDetails struct {
    Name string `json:"name"`
    TypeName string `json:"typeName"`
    ItemMeasureReferenceText string `json:"itemMeasureReferenceText"`
    MainImageUrl string `json:"mainImageUrl"`
    PipUrl string `json:"pipUrl"`
    Id string `json:"id"`
    ItemNoGlobal string `json:"itemNoGlobal"`
    OnlineSellable bool `json:"onlineSellable"`
    LastChance bool `json:"lastChance"`
    GprDescription GprDescription `json:"gprDescription"`
    Colors []Color `json:"colors"`
    Badge Badge `json:"badge"`
    Tag string `json:"tag"`
    TagText string `json:"tagText"`
    QuickFacts []interface{} `json:"quickFacts"`
    Features []interface{} `json:"features"`
    Availability []interface{} `json:"availability"`
    RatingValue float64 `json:"ratingValue"`
    RatingCount int `json:"ratingCount"`
    ItemNo string `json:"itemNo"`
    ItemType string `json:"itemType"`
    SalesPrice SalesPrice `json:"salesPrice"`
    ContextualImageUrl string `json:"contextualImageUrl"`
    MainImageAlt string `json:"mainImageAlt"`
    BusinessStructure BusinessStructure `json:"businessStructure"`
    CategoryPath []CategoryPath `json:"categoryPath"`
    HeroBackoffData struct {
    } `json:"heroBackoffData"`
    OptimizelyAttributes OptimizelyAttributes `json:"optimizelyAttributes"`
}

type GprDescription struct {
    NumberOfVariants int `json:"numberOfVariants"`
    Variants []Variant `json:"variants"`
}

type Variant struct {
    Id string `json:"id"`
    PipUrl string `json:"pipUrl"`
    ImageUrl string `json:"imageUrl"`
    ImageAlt string `json:"imageAlt"`
    QuickFacts []interface{} `json:"quickFacts"`
    Availability []interface{} `json:"availability"`
    RatingValue float64 `json:"ratingValue"`
    RatingCount int `json:"ratingCount"`
    Name string `json:"name"`
    TypeName string `json:"typeName"`
    ItemMeasureReferenceText string `json:"itemMeasureReferenceText"`
    MainImageUrl string `json:"mainImageUrl"`
    MainImageAlt string `json:"mainImageAlt"`
    ContextualImageUrl string `json:"contextualImageUrl"`
    ItemNoGlobal string `json:"itemNoGlobal"`
    OnlineSellable bool `json:"onlineSellable"`
    LastChance bool `json:"lastChance"`
    ItemNo string `json:"itemNo"`
    ItemType string `json:"itemType"`
    SalesPrice SalesPrice `json:"salesPrice"`
    OptimizelyAttributes OptimizelyAttributes `json:"optimizelyAttributes"`
}

type Color struct {
    Name string `json:"name"`
    Id string `json:"id"`
    Hex string `json:"hex"`
}

type Badge struct {
    Type string `json:"type"`
    Text string `json:"text"`
}

type SalesPrice struct {
    CurrencyCode string `json:"currencyCode"`
    Numeral float64 `json:"numeral"`
    Current Current `json:"current"`
    Previous Previous `json:"previous"`
    IsBreathTaking bool `json:"isBreathTaking"`
    Discount string `json:"discount"`
    PrevPriceLabel string `json:"prevPriceLabel"`
    ValidFrom string `json:"validFrom"`
    ValidTo string `json:"validTo"`
    Tag string `json:"tag"`
    TagText string `json:"tagText"`
    PriceText string `json:"priceText"`
}

type Current struct {
    Prefix string `json:"prefix"`
    WholeNumber string `json:"wholeNumber"`
    Separator string `json:"separator"`
    Decimals string `json:"decimals"`
    Suffix string `json:"suffix"`
    IsRegularCurrency bool `json:"isRegularCurrency"`
}

type Previous struct {
    Prefix string `json:"prefix"`
    WholeNumber string `json:"wholeNumber"`
    Separator string `json:"separator"`
    Decimals string `json:"decimals"`
    Suffix string `json:"suffix"`
    IsRegularCurrency bool `json:"isRegularCurrency"`
}

type BusinessStructure struct {
    HomeFurnishingBusinessName string `json:"homeFurnishingBusinessName"`
    HomeFurnishingBusinessNo string `json:"homeFurnishingBusinessNo"`
    ProductAreaName string `json:"productAreaName"`
    ProductAreaNo string `json:"productAreaNo"`
    ProductRangeAreaName string `json:"productRangeAreaName"`
    ProductRangeAreaNo string `json:"productRangeAreaNo"`
}

type CategoryPath struct {
    Name string `json:"name"`
    Key string `json:"key"`
}

type OptimizelyAttributes struct {
    PRODUCTTYPE string `json:"PRODUCT_TYPE"`
}
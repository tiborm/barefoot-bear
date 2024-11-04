package main

import "time"

type Colour struct {
	RGBDec string `json:"rgbDec"`
	RGBHex string `json:"rgbHex"`
	Token  string `json:"token"`
}

type Communication struct {
	Colour      Colour `json:"colour"`
	MessageType string `json:"messageType"`
}

type Probability struct {
	Type           string        `json:"type"`
	UpdateDateTime time.Time     `json:"updateDateTime"`
	Communication  Communication `json:"communication"`
}

type AvailableStock struct {
	Type           string        `json:"type"`
	Quantity       int           `json:"quantity"`
	UpdateDateTime time.Time     `json:"updateDateTime"`
	Probabilities  []Probability `json:"probabilities"`
}

type ClassUnitKey struct {
	ClassUnitCode string `json:"classUnitCode"`
	ClassUnitType string `json:"classUnitType"`
}

type ItemKey struct {
	ItemNo   string `json:"itemNo"`
	ItemType string `json:"itemType"`
}

type InventoryData struct {
	IsInCashAndCarryRange bool             `json:"isInCashAndCarryRange"`
	IsInHomeDeliveryRange bool             `json:"isInHomeDeliveryRange"`
	AvailableStocks       []AvailableStock `json:"availableStocks"`
	ClassUnitKey          ClassUnitKey     `json:"classUnitKey"`
	ItemKey               ItemKey          `json:"itemKey"`
}

type InventoryJsonResponse struct {
	Availabilities interface{}     `json:"availabilities,omitempty"`
	Data           []InventoryData `json:"data"`
}

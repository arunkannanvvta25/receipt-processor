package model

type Receipt struct {
	ID           string        `json:"id"`
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Total        string        `json:"total"`
	Items        []ReceiptItem `json:"items"`
}

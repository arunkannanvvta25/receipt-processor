package tests

import (
	"receipt-processor-challenge/api/data"
	"receipt-processor-challenge/api/model"
	"testing"
)

func TestCreateReceipt(t *testing.T) {
	// Initialize the DataAccessLayer
	dal := data.NewDataAccessLayer()

	// Create a sample receipt
	receipt := model.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []model.ReceiptItem{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		},
		Total: "18.74",
	}

	// Call CreateReceipt
	id, err := dal.CreateReceipt(receipt)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check that the receipt was stored
	storedReceipt, err := dal.GetReceiptByID(id)
	if err != nil {
		t.Errorf("receipt not found")
	}
	// Verify the stored receipt matches the original
	if storedReceipt.ID != id {
		t.Errorf("stored receipt does not match original: got %+v, want %+v", storedReceipt, receipt)
	}
}

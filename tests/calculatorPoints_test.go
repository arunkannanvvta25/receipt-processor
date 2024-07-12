package tests

import (
	"receipt-processor-challenge/api/model"
	"receipt-processor-challenge/api/service"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	strategy := service.PointCalculatorStrategy_one{}
	receipt := model.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []model.ReceiptItem{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
		Total: "35.35",
	}

	expectedPoints := int64(28)

	points, err := strategy.CalculatePoints(receipt)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if points != expectedPoints {
		t.Errorf("expected %d points, got %d", expectedPoints, points)
	}
}

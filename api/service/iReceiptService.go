package service

import (
	"receipt-processor-challenge/api/model"
)

type ReceiptServiceInterface interface {
	CreateReceipt(receipt model.Receipt) (string, error)
	GetPointsForReceipt(id string) (int64, error)
}

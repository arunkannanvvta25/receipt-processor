package data

import "receipt-processor-challenge/api/model"

type DataAccessLayerInterface interface {
	CreateReceipt(receipt model.Receipt) (string, error)
	GetReceiptByID(id string) (model.Receipt, error)
}

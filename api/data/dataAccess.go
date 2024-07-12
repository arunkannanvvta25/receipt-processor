package data

import (
	"errors"
	"receipt-processor-challenge/api/model"

	"github.com/google/uuid"
)

type DataAccessLayerImpl struct {
	receipts map[string]model.Receipt // Map to store receipts by ID
}

func NewDataAccessLayer() *DataAccessLayerImpl {
	return &DataAccessLayerImpl{
		receipts: make(map[string]model.Receipt),
	}
}

func (dal *DataAccessLayerImpl) CreateReceipt(receipt model.Receipt) (string, error) {
	receipt.ID = dal.generateID()
	dal.receipts[receipt.ID] = receipt
	return receipt.ID, nil
}

func (dal *DataAccessLayerImpl) GetReceiptByID(id string) (model.Receipt, error) {
	receipt, ok := dal.receipts[id]
	if !ok {
		return model.Receipt{}, errors.New("receipt not found")
	}
	return receipt, nil
}

func (dal *DataAccessLayerImpl) generateID() string {
	u := uuid.New()
	return u.String()
}

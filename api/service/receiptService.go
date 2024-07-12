package service

import (
	"receipt-processor-challenge/api/data"
	"receipt-processor-challenge/api/model"
)

// ReceiptServiceImpl implements ReceiptService.
type ReceiptServiceImpl struct {
	dal      data.DataAccessLayerInterface
	strategy PointsCalculationStrategyInterface
}

func NewReceiptService(strategy PointsCalculationStrategyInterface, dal data.DataAccessLayerInterface) *ReceiptServiceImpl {
	return &ReceiptServiceImpl{
		dal:      dal,
		strategy: strategy,
	}
}

func (rs *ReceiptServiceImpl) CreateReceipt(receipt model.Receipt) (string, error) {
	return rs.dal.CreateReceipt(receipt)
}

func (rs *ReceiptServiceImpl) GetPointsForReceipt(id string) (int64, error) {
	receipt, err := rs.dal.GetReceiptByID(id)
	if err != nil {
		return 0, err
	}

	points, err := rs.strategy.CalculatePoints(receipt)
	if err != nil {
		return 0, err
	}

	return points, nil
}

package service

import (
	"receipt-processor-challenge/api/model"
)

type PointsCalculationStrategyInterface interface {
	CalculatePoints(receipt model.Receipt) (int64, error)
}

package main

import (
	"receipt-processor-challenge/api/controller"
	"receipt-processor-challenge/api/data"
	"receipt-processor-challenge/api/router"
	"receipt-processor-challenge/api/service"
)

func main() {
	dal := data.NewDataAccessLayer()                   // Instantiating the DataAccessLayer implementation
	strategy := &service.PointCalculatorStrategy_one{} // Instantiating the PointsCalculationStrategy implementation

	// Create ReceiptService with injected dependencies
	receiptService := service.NewReceiptService(strategy, dal)

	// Create ReceiptController with injected ReceiptService
	receiptController := controller.NewReceiptController(receiptService)

	// Setup router and run your application
	r := router.SetupRouter(receiptController)
	r.Run(":8080")
}

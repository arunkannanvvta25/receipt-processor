package router

import (
	"receipt-processor-challenge/api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(receiptController *controller.ReceiptController) *gin.Engine {
	r := gin.Default()

	r.POST("/receipts/process", receiptController.PostReceiptForProcessing)
	r.GET("/receipts/:id/points", receiptController.GetPointsForReceipt)

	return r
}

package controller

import (
	"errors"
	"net/http"
	"receipt-processor-challenge/api/model"
	"receipt-processor-challenge/api/service"

	"github.com/gin-gonic/gin"
)

type ReceiptController struct {
	Service service.ReceiptServiceInterface
}

func NewReceiptController(service service.ReceiptServiceInterface) *ReceiptController {
	return &ReceiptController{
		Service: service,
	}
}

func (rc *ReceiptController) PostReceiptForProcessing(c *gin.Context) {
	var newReceipt model.Receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := rc.validateReceipt(&newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receiptID, err := rc.Service.CreateReceipt(newReceipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": receiptID})
}

func (rc *ReceiptController) GetPointsForReceipt(c *gin.Context) {
	id := c.Param("id")
	points, err := rc.Service.GetPointsForReceipt(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "receipt not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"points": points})
}

func (rc *ReceiptController) validateReceipt(receipt *model.Receipt) error {
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || len(receipt.Items) == 0 || receipt.Total == "" {
		return errors.New("missing required fields in receipt")
	}
	return nil
}

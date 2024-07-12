package service

import (
	"errors"
	"math"
	"receipt-processor-challenge/api/model"
	"strconv"
	"strings"
	"time"
)

type PointCalculatorStrategy_one struct{}

func (s *PointCalculatorStrategy_one) CalculatePoints(receipt model.Receipt) (int64, error) {
	points := int64(0)
	// Rule 1: One point for every alphanumeric character in the retailer name.
	retailerName := receipt.Retailer
	var count int64
	for _, char := range retailerName {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			count++
		}
	}
	points += count

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	totalFloat, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, errors.New("invalid total amount")
	}
	if totalFloat == math.Floor(totalFloat) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if math.Mod(totalFloat, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	numItems := len(receipt.Items)
	points += int64(numItems / 2 * 5)

	// Rule 5: If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			priceFloat, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, errors.New("invalid item price")
			}
			points += int64(math.Ceil(priceFloat * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd.
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return 0, errors.New("invalid purchase date format")
	}
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, errors.New("invalid purchase time format")
	}
	startTime, _ := time.Parse("15:04", "14:00")
	endTime, _ := time.Parse("15:04", "16:00")
	if purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
		points += 10
	}

	return points, nil
}

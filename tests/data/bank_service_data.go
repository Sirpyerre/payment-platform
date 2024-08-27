package data

import "github.com/Sirpyerre/payment-platform/internal/models"

var ResponsePaymentData = &models.ResponsePayment{
	TransactionID: "12345657890",
	Status:        "success",
	Message:       "Payment processed successfully",
}

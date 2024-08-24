package models

type SuccessPayment struct {
	Status        string `json:"status"`
	TransactionID string `json:"transaction_id"`
}

type ResponsePayment struct {
	Status        string `json:"status"`
	TransactionID string `json:"transaction_id"`
	Message       string `json:"message"`
}

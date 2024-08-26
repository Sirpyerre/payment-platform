package models

import "time"

type Refund struct {
	ID            int
	TransactionID int
	Amount        float64
	Status        string
	CreatedAt     time.Time
}

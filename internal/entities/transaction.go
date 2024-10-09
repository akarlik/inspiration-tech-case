package entities

import "time"

type Transaction struct {
	TransactionId string    `json:"transactionId"`
	MessageType   string    `json:"messageType"`
	Origin        string    `json:"origin"`
	Amount        float64   `json:"amount"`
	Commission    float64   `json:"commission"`
	Timestamp     time.Time `json:"timestamp"`
}

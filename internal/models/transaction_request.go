package models

type TransactionRequest struct {
	MessageType   string  `json:"messageType"`   // PAYMENT | ADJUSTMENT
	TransactionId string  `json:"transactionId"` // GUID
	AccountId     int64   `json:"accountId"`
	Origin        string  `json:"origin"`        // VISA | MASTER
	Amount        float64 `json:"amount"`
}


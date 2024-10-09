package models

type AccountResponse struct {
	AccountId   int64         `json:"accountId"`
	Balance     float64       `json:"balance"`
}
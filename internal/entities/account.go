package entities

type Account struct {
	AccountId    int64         `json:"accountId"`
	Balance      float64       `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

func (account *Account) ValidateDuplicateTransaction(transactionId string) bool {
	for _, transaction := range account.Transactions {
		if transaction.MessageType == "PAYMENT" && transaction.TransactionId == transactionId {
			return false
		}
	}
	return true
}

func (account *Account) ValidateAdjustmentTransactionAmount(transactionId string, newAmount float64) bool {
	paymentAmount := sumAmountsByMessageType(account.Transactions, transactionId, "PAYMENT")
	adjusmentAmount := sumAmountsByMessageType(account.Transactions, transactionId, "ADJUSTMENT")
	remain := paymentAmount - adjusmentAmount
	return remain > newAmount
}

func sumAmountsByMessageType(transactions []Transaction, transactionId string, messageType string) float64 {
	var totalAmount float64
	for _, transaction := range transactions {
		if transaction.TransactionId == transactionId && transaction.MessageType == messageType {
			totalAmount += (transaction.Amount + transaction.Commission)
		}
	}
	return totalAmount
}

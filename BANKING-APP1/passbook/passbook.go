package passbook

import "time"

var transactionId int = 0

type Transaction struct {
	TransactionID int
	Timestamp     time.Time
	Type          string
	Amount        float32
	Balance       float32
	Note          string
}

func NewTransaction(tType string, amount float32, balance float32, note string) Transaction {

	transactionId++

	return Transaction{
		TransactionID: transactionId,
		Timestamp:     time.Now(),
		Type:          tType,
		Amount:        amount,
		Balance:       balance,
		Note:          note,
	}
}

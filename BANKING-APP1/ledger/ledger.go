package ledger

import "errors"

var transactionCounter int = 0

type BankTransaction struct {
	BankTransactionID int
	SenderBankID      int
	ReceiverBankID    int
	Amount            float32
}

func NewBankTransaction(senderBankID int, recieverBankID int, amount float32) (*BankTransaction, error) {

	if senderBankID < 0 {
		return nil, errors.New("sender bank ID cannot be negative")
	}
	if recieverBankID < 0 {
		return nil, errors.New("reciever bank ID cannot be negative")
	}
	if amount < 0 {
		return nil, errors.New("amount cannot be negative")
	}

	return &BankTransaction{
		BankTransactionID: transactionCounter,
		SenderBankID:      senderBankID,
		ReceiverBankID:    recieverBankID,
		Amount:            amount,
	}, nil
}

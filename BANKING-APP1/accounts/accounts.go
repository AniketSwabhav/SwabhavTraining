package accounts

import (
	"errors"
)

var accountNo int = 100
var AccountsMap = make(map[int]*Accounts)

type Accounts struct {
	AccountNo int
	Balance   float32
	BankId    int
	UserID    int
	Passbook  []string
}

func NewAccount(bankId int, userId int) (*Accounts, error) {
	if bankId <= 0 {
		return nil, errors.New("invalid bank ID")
	}

	accountNo++

	newAcc := &Accounts{
		AccountNo: accountNo,
		Balance:   1000,
		BankId:    bankId,
		UserID:    userId,
		Passbook:  []string{"Account created with initial balance Rs.1000"},
	}
	AccountsMap[newAcc.AccountNo] = newAcc
	return newAcc, nil
}

func GetReceiverAccountById(accountNo int) (*Accounts, error) {
	if acc, exists := AccountsMap[accountNo]; exists {
		return acc, nil
	}
	return nil, errors.New("receiver account not found")
}

func (acc *Accounts) Transfer(amount float32, targetAccNo int) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	if acc.Balance < amount {
		return errors.New("insufficient balance for transfer")
	}

	targetAcc, err := GetReceiverAccountById(targetAccNo)
	if err != nil {
		return err
	}

	acc.Balance -= amount
	targetAcc.Balance += amount
	return nil
}

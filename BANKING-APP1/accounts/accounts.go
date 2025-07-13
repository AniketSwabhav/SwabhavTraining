package accounts

import (
	"banking_app/passbook"
	"banking_app/util"
	"errors"
	"fmt"
	"time"
)

var accountNo int = 100
var AccountsMap = make(map[int]*Accounts)

type Accounts struct {
	AccountNo int
	Balance   float32
	BankId    int
	UserID    int
	Passbook  []passbook.Transaction
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
		Passbook: []passbook.Transaction{
			{
				Timestamp: time.Now(),
				Type:      "AccountCreation",
				Amount:    1000,
				Balance:   1000,
				Note:      "Account created with initial balance Rs.1000",
			},
		},
	}
	AccountsMap[newAcc.AccountNo] = newAcc
	return newAcc, nil
}

func GetReceiverAccountById(accountNo int) (*Accounts, error) {

	defer util.HandlePanic()

	if acc, exists := AccountsMap[accountNo]; exists {
		return acc, nil
	}
	return nil, errors.New("receiver account not found")
}

func (acc *Accounts) SelfTransfer(amount float32, toacc *Accounts) error {

	defer util.HandlePanic()

	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}
	if acc.Balance < amount {
		return errors.New("insufficient balance for transfer")
	}

	acc.Balance -= amount
	toacc.Balance += amount

	debitTransaction := passbook.NewTransaction("SelfTransferDebit", amount, acc.Balance, fmt.Sprintf("Transferred Rs.%.2f to Account #%d", amount, toacc.AccountNo))
	acc.Passbook = append(acc.Passbook, debitTransaction)

	creditTransaction := passbook.NewTransaction("SelfTransferCredit", amount, toacc.Balance, fmt.Sprintf("Received Rs.%.2f from Account #%d", amount, acc.AccountNo))
	toacc.Passbook = append(toacc.Passbook, creditTransaction)
	return nil

}

func (acc *Accounts) BankTransfer(amount float32, targetAccNo int) error {

	defer util.HandlePanic()

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

	debitTransaction := passbook.NewTransaction("BankTransferDebit", amount, acc.Balance, fmt.Sprintf("Transferred Rs.%.2f to Account #%d", amount, targetAccNo))
	acc.Passbook = append(acc.Passbook, debitTransaction)

	creditTransaction := passbook.NewTransaction("BankTransferCredit", amount, targetAcc.Balance, fmt.Sprintf("Received Rs.%.2f from Account #%d", amount, acc.AccountNo))
	targetAcc.Passbook = append(targetAcc.Passbook, creditTransaction)

	return nil
}

func (acc *Accounts) Withdraw(amount float32) error {

	defer util.HandlePanic()

	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if acc.Balance < amount {
		return errors.New("insufficient balance for withdrawal")
	}
	acc.Balance -= amount

	withdrawTransaction := passbook.NewTransaction("Withdrawal", amount, acc.Balance, fmt.Sprintf("Withdrawn Rs.%.2f", amount))
	acc.Passbook = append(acc.Passbook, withdrawTransaction)
	return nil
}

func (acc *Accounts) Deposit(amount float32) error {

	defer util.HandlePanic()

	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}

	acc.Balance += amount

	depositTransaction := passbook.NewTransaction("Deposite", amount, acc.Balance, fmt.Sprintf("Deposited Rs.%.2f", amount))
	acc.Passbook = append(acc.Passbook, depositTransaction)
	return nil
}

func (acc *Accounts) GetPassbook(page, pageSize int) ([]passbook.Transaction, error) {

	defer util.HandlePanic()

	if page <= 0 || pageSize <= 0 {
		return nil, errors.New("page and pageSize must be positive integers")
	}

	start := (page - 1) * pageSize

	if start >= len(acc.Passbook) {
		return []passbook.Transaction{}, nil
	}

	end := start + pageSize
	if end > len(acc.Passbook) {
		end = len(acc.Passbook)
	}

	return acc.Passbook[start:end], nil

}

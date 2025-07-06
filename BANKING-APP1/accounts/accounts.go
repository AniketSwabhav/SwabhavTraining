package accounts

import (
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
	Passbook  []PassbookEntry
}

type PassbookEntry struct {
	Timestamp time.Time
	Type      string
	Amount    float32
	Balance   float32
	Note      string
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
		Passbook: []PassbookEntry{
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
	if acc, exists := AccountsMap[accountNo]; exists {
		return acc, nil
	}
	return nil, errors.New("receiver account not found")
}

func (acc *Accounts) SelfTransfer(amount float32, toacc *Accounts) error {

	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}
	if acc.Balance < amount {
		return errors.New("insufficient balance for transfer")
	}

	acc.Balance -= amount
	toacc.Balance += amount

	acc.Passbook = append(acc.Passbook, PassbookEntry{
		Timestamp: time.Now(),
		Type:      "SelfTransferDebit",
		Amount:    amount,
		Balance:   acc.Balance,
		Note:      fmt.Sprintf("Transferred Rs.%.2f to Account #%d", amount, toacc.AccountNo),
	})

	toacc.Passbook = append(toacc.Passbook, PassbookEntry{
		Timestamp: time.Now(),
		Type:      "SelfTransferCredit",
		Amount:    amount,
		Balance:   toacc.Balance,
		Note:      fmt.Sprintf("Received Rs.%.2f from Account #%d", amount, acc.AccountNo),
	})

	return nil
}

func (acc *Accounts) BankTransfer(amount float32, targetAccNo int) error {

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

	acc.Passbook = append(acc.Passbook, PassbookEntry{
		Timestamp: time.Now(),
		Type:      "BankTransferDebit",
		Amount:    amount,
		Balance:   acc.Balance,
		Note:      fmt.Sprintf("Transferred Rs.%.2f to Account #%d", amount, targetAccNo),
	})

	targetAcc.Passbook = append(targetAcc.Passbook, PassbookEntry{
		Timestamp: time.Now(),
		Type:      "BankTransferCredit",
		Amount:    amount,
		Balance:   targetAcc.Balance,
		Note:      fmt.Sprintf("Received Rs.%.2f from Account #%d", amount, acc.AccountNo),
	})

	return nil
}

func (acc *Accounts) Withdraw(amount float32) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if acc.Balance < amount {
		return errors.New("insufficient balance for withdrawal")
	}
	acc.Balance -= amount

	acc.Passbook = append(acc.Passbook, PassbookEntry{
		Timestamp: time.Now(),
		Type:      "Withdrawal",
		Amount:    amount,
		Balance:   acc.Balance,
		Note:      fmt.Sprintf("Withdrawn Rs.%.2f", amount),
	})

	return nil
}

func (acc *Accounts) Deposit(amount float32) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	acc.Balance += amount

	acc.Passbook = append(acc.Passbook, PassbookEntry{
		Timestamp: time.Now(),
		Type:      "Deposit",
		Amount:    amount,
		Balance:   acc.Balance,
		Note:      fmt.Sprintf("Deposited Rs.%.2f", amount),
	})

	return nil
}

func (acc *Accounts) GetPassbook() []PassbookEntry {
	return acc.Passbook
}

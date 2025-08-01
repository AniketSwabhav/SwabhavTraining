package bank

import (
	"banking_app/accounts"
	"banking_app/ledger"
	"banking_app/util"
	"errors"
	"fmt"
	"strings"
)

var bankId = 0
var allBanks = make(map[int]*Bank)

type Bank struct {
	BankID           int
	FullName         string
	Abbreviation     string
	Accounts         []accounts.Accounts
	BankTransactions []ledger.BankTransaction
}

func NewBank(fullName string) (*Bank, error) {

	if fullName == "" {
		return nil, errors.New("bank Name cannot be empty")
	}
	bankId++
	bank := &Bank{
		BankID:       bankId,
		FullName:     fullName,
		Abbreviation: getAbbreviation(fullName),
	}

	allBanks[bankId] = bank
	return bank, nil
}

func getAbbreviation(input string) string {
	words := strings.Fields(input)
	var firstLetters []string

	for _, word := range words {
		if len(word) > 0 {
			firstLetters = append(firstLetters, string(word[0]))
		}
	}
	return strings.Join(firstLetters, "")
}

func GetBank(bankId int) (*Bank, error) {

	bank, exists := allBanks[bankId]
	if !exists {
		return nil, errors.New("bank not found")
	}
	return bank, nil
}

func GetAllBanks() []Bank {
	totalBanks := []Bank{}
	for _, bank := range allBanks {
		totalBanks = append(totalBanks, *bank)
	}
	return totalBanks
}

func (b *Bank) UpdateBank(param string, value interface{}) error {

	defer util.HandlePanic()

	if param == "" {
		return errors.New("parameter cannot be empty")
	}
	switch param {
	case "FullName":
		b.updateBankFullName(value)
	default:
		return errors.New("provide valid parameters")
	}
	return nil
}

func (b *Bank) updateBankFullName(value interface{}) error {

	defer util.HandlePanic()

	strVal, ok := value.(string)
	if !ok || strVal == "" {
		return errors.New("value is empty, provide valid value")
	}
	b.FullName = strVal
	b.Abbreviation = getAbbreviation(strVal)
	fmt.Println("Bank name updated successfully")
	return nil
}

func DeleteBank(bankId int) error {

	defer util.HandlePanic()

	b, exists := allBanks[bankId]
	if !exists {
		return errors.New("bank not found")
	}

	if len(b.Accounts) > 0 {
		return errors.New("cannot delete bank with active accounts")
	}

	delete(allBanks, bankId)
	fmt.Println("Bank deleted successfully")
	return nil
}

func (b *Bank) CreateNewBankTransaction(senderBankID int, recieverBankID int, amount float32) *ledger.BankTransaction {

	newBankTransaction, err := ledger.NewBankTransaction(senderBankID, recieverBankID, amount)
	if err != nil {
		return nil
	}

	b.BankTransactions = append(b.BankTransactions, *newBankTransaction)
	return newBankTransaction
}

func (b *Bank) GetBankTransactionAmount(bankID int) (float32, error) {

	defer util.HandlePanic()

	if bankID < 0 {
		return 0, errors.New("bank ID cannot be negative")
	}

	var totalAmount float32 = 0.0

	for _, transactions := range b.BankTransactions {

		if (transactions.SenderBankID == b.BankID) && (transactions.ReceiverBankID == bankID) {
			totalAmount -= transactions.Amount
		} else if (transactions.SenderBankID == bankID) && (transactions.ReceiverBankID == b.BankID) {
			totalAmount += transactions.Amount
		}
	}

	return totalAmount, nil
}

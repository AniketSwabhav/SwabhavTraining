package bank

import (
	"banking_app/accounts"
	"errors"
	"fmt"
	"strings"
)

var bankId = 0
var allBanks = make(map[int]*Bank)

type Bank struct {
	BankID       int
	FullName     string
	Abbreviation string
	Accounts     []accounts.Accounts
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

func (b *Bank) UpdateBank(param string, value interface{}) error {
	if param == "" {
		return errors.New("parameter cannot be empty")
	}
	switch param {
	case "FullName":
		return b.updateBankFullName(value)
	default:
		return errors.New("provide valid parameters")
	}
}

func (b *Bank) updateBankFullName(value interface{}) error {
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

func GetAccount(accountNo int) {

}

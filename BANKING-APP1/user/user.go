package user

import (
	"banking_app/accounts"
	"banking_app/bank"
	"banking_app/passbook"
	"banking_app/util"
	"errors"
	"fmt"
)

var userId int = 0
var userMap = make(map[int]*User)

type User struct {
	UserID       int
	FirstName    string
	LastName     string
	TotalBalance float32
	Accounts     []*accounts.Accounts
	Banks        []*bank.Bank
	IsAdmin      bool
	IsActive     bool
}

func NewUser(firstName, lastName string, isAdmin bool) (*User, error) {
	if firstName == "" {
		return nil, errors.New("first name cannot be empty")
	}
	if lastName == "" {
		return nil, errors.New("last name cannot be empty")
	}

	userId++

	user := &User{
		UserID:    userId,
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   isAdmin,
		IsActive:  true,
	}

	userMap[userId] = user

	return user, nil
}

func (u *User) CreateAccount(bankId int) {

	defer util.HandlePanic()

	if u.IsAdmin {
		panic("only customers can create bank accounts")
	}
	if !u.IsActive {
		panic("customer needs to be active to get  banks")
	}

	targetBank := bank.GetBank(bankId)

	newAccount, err := accounts.NewAccount(bankId, u.UserID)
	if err != nil {
		panic(err)
	}

	u.Accounts = append(u.Accounts, newAccount)
	u.TotalBalance += newAccount.Balance

	targetBank.Accounts = append(targetBank.Accounts, *newAccount)

	fmt.Printf("Account created successfully for %s %s | Account No: %d | Balance: %.2f | Bank: %s\n",
		u.FirstName, u.LastName, newAccount.AccountNo, newAccount.Balance, targetBank.FullName)
}

func (u *User) CalculateTotalBalance() float32 {

	defer util.HandlePanic()

	if u.IsAdmin {
		panic("only customers can see Total Balance")
	}
	if !u.IsActive {
		panic("customer needs to be active to get  calculate balance")
	}
	total := float32(0)
	for _, acc := range u.Accounts {
		total += acc.Balance
	}
	u.TotalBalance = total
	return total
}

func (u *User) GetMyAccountBlance(accountNo int) float32 {

	defer util.HandlePanic()

	if u.IsAdmin {
		panic("only valid customer can see account balance")
	}
	if !u.IsActive {
		panic("customer needs to be active to get balance")
	}
	for i := range u.Accounts {
		if u.Accounts[i].AccountNo == accountNo {
			return u.Accounts[i].Balance
		}
	}
	panic("provided account number is not related to requesting user")
}

func (u *User) GetSelfAccountById(accountNo int) *accounts.Accounts {

	defer util.HandlePanic()

	if u.IsAdmin {
		panic("only customer can get own account")
	}
	if !u.IsActive {
		panic("customer needs to be active to get account")
	}

	for i := range u.Accounts {
		if u.Accounts[i].AccountNo == accountNo {
			return u.Accounts[i]
		}
	}
	panic("account not found for use ")
}

func (u *User) TransferBetweenSelfAccounts(fromAccNo, toAccNo int, amount float32) {

	defer util.HandlePanic()

	if u.IsAdmin {
		panic("admin cannot perform transfers")
	}
	if !u.IsActive {
		panic("customer needs to be active to transefer money")
	}

	fromAcc := u.GetSelfAccountById(fromAccNo)

	toAcc := u.GetSelfAccountById(toAccNo)

	if fromAcc.AccountNo == toAcc.AccountNo {
		panic("cannot transfer to the same account")
	}

	fromAcc.SelfTransfer(amount, toAcc)

	fmt.Printf("Successfully transferred Rs.%.2f from Acc#%d to Acc#%d\n", amount, fromAccNo, toAccNo)
}

func (u *User) TransferToOtherUser(fromAccNo, targetAccNo int, amount float32) {

	defer util.HandlePanic()

	if u.IsAdmin {
		panic("admin cannot perform transfers")
	}
	if !u.IsActive {
		panic("customer needs to be active to transefer money")
	}

	fromAcc := u.GetSelfAccountById(fromAccNo)

	if fromAcc == nil {
		panic("source account not found for user")
	}

	targetAcc := accounts.GetReceiverAccountById(targetAccNo)

	senderBankID := fromAcc.BankId
	receiverBankID := targetAcc.BankId

	senderBank := bank.GetBank(senderBankID)

	senderBank.CreateNewBankTransaction(senderBankID, receiverBankID, amount)

	fromAcc.BankTransfer(amount, targetAccNo)

	fmt.Printf("Successfully transferred Rs.%.2f from your Acc#%d to target Acc#%d\n", amount, fromAccNo, targetAccNo)
}

func (u *User) WithdrawFromAccount(accountNo int, amount float32) {

	defer util.HandlePanic()

	if u.IsAdmin {
		panic("admin cannot perform withdrawals")
	}
	if !u.IsActive {
		panic("customer needs to be active to withdraw money")
	}

	for _, acc := range u.Accounts {
		if acc.AccountNo == accountNo {
			acc.Withdraw(amount)
		}
	}

	panic("account not found for this user")
}

func (u *User) DepositToAccount(accountNo int, amount float32) error {

	defer util.HandlePanic()

	if u.IsAdmin {
		return errors.New("admin cannot perform deposits")
	}
	if !u.IsActive {
		panic("customer needs to be active to deposite money")
	}

	for _, acc := range u.Accounts {
		if acc.AccountNo == accountNo {
			acc.Deposit(amount)
		}
	}

	return errors.New("account not found for this user")
}

func (u *User) ViewMyPassbook(accountNo, page, pageSize int) []passbook.Transaction {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only customer can see his own passbook")
	}
	if !u.IsActive {
		panic("customer needs to be active to view passbook")
	}

	acc := u.GetSelfAccountById(accountNo)

	return acc.GetPassbook(page, pageSize)
}

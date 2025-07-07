package user

import (
	"banking_app/accounts"
	"banking_app/bank"
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
		IsAdmin:   true,
	}

	userMap[userId] = user

	return user, nil
}

func NewAdmin(firstName, lastName string) (*User, error) {

	newAdmin, err := NewUser(firstName, lastName, true)
	if err != nil {
		return nil, err
	}
	return newAdmin, nil
}

// ======================================================================Bank Related Methods=======================================================================

func (u *User) AddBank(fullName string) (*bank.Bank, error) {
	if !u.IsAdmin {
		return nil, errors.New("only admin can add banks")
	}

	newBank, err := bank.NewBank(fullName)
	if err != nil {
		return nil, err
	}
	return newBank, nil
}

func (u *User) GetBankById(bankId int) (*bank.Bank, error) {
	if !u.IsAdmin {
		return nil, errors.New("only admin can add customers")
	}

	// return bank.GetBank(bankId)
	bank, err := bank.GetBank(bankId)
	if err != nil {
		return nil, err
	}

	return bank, nil
}

func (u *User) UpdateBankById(bankId int, param string, value interface{}) error {
	if !u.IsAdmin {
		return errors.New("only admin can add banks")
	}

	if len(u.Banks) == 0 {
		return errors.New("no banks associated with this admin")
	}
	bankToBeUpdated, err := u.GetBankById(bankId)
	if err != nil {
		return err
	}
	err = bankToBeUpdated.UpdateBank(param, value)
	if err != nil {
		return err
	}
	return err
}

func (u *User) DeleteBankById(bankId int) error {

	if !u.IsAdmin {
		return errors.New("only admin can add customers")
	}
	err := bank.DeleteBank(bankId)
	if err != nil {
		return err
	}

	index := -1
	for i, b := range u.Banks {
		if b.BankID == bankId {
			index = i
			break
		}
	}
	if index != -1 {
		u.Banks = append(u.Banks[:index], u.Banks[index+1:]...)
	}

	return nil
}

// =====================================================================Customer Related Methods========================================================================

func (u *User) NewCustomer(firstName, lastName string) (*User, error) {
	if !u.IsAdmin {
		return nil, errors.New("only admin can add customers")
	}
	// createdStaff = NewUser(fName, lName, false)
	newCustomer, err := NewUser(firstName, lastName, false)
	if err != nil {
		return nil, err
	}
	return newCustomer, nil
}

func (u *User) GetCustomerById(customerId int) (*User, error) {
	if !u.IsAdmin {
		return nil, errors.New("only admin can Perform CRUD on customers")
	}
	customer, exists := userMap[customerId]
	if !exists {
		return nil, errors.New("customer not found with the given ID")
	}
	return customer, nil
}

func (u *User) UpdateCustomerById(customerId int, param string, value interface{}) error {

	if !u.IsAdmin {
		return errors.New("only admin can Perform CRUD on customers")
	}
	if customerId < 0 {
		return errors.New("customerId cannot be negative")
	}
	targetCustomer, err := u.GetCustomerById(customerId)
	if err != nil {
		return err
	}
	if param == "" {
		return errors.New("parameter cannot be empty")
	}

	switch param {
	case "FirstName":
		return targetCustomer.updateFirstName(value)
	case "LastName":
		return targetCustomer.updateLastName(value)
	default:
		return errors.New("invalid parameter for update")
	}
}

func (target *User) updateFirstName(value interface{}) error {
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		return errors.New("value is empty, provide valid value")
	}
	target.FirstName = strVal
	fmt.Println("First name updated successfully")
	return nil
}

func (target *User) updateLastName(value interface{}) error {
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		return errors.New("value is empty, provide valid value")
	}
	target.LastName = strVal
	fmt.Println("Last name updated successfully")
	return nil
}

func (u *User) DeleteCustomerById(customerId int) error {
	if !u.IsAdmin {
		return errors.New("only admin can Perform CRUD on customers")
	}
	if customerId < 0 {
		return errors.New("customerId cannot be negative")
	}

	_, exists := userMap[customerId]
	if !exists {
		return errors.New("customer not found with the given ID")
	}

	delete(userMap, customerId)
	return nil
}

// ======================================================================Transaction related Methods===================================================================

func (u *User) CreateAccount(bankId int) error {
	if u.IsAdmin {
		return errors.New("only customers can create bank accounts")
	}
	targetBank, err := bank.GetBank(bankId)
	if err != nil {
		return err
	}
	newAccount, err := accounts.NewAccount(bankId, u.UserID)
	if err != nil {
		return err
	}

	u.Accounts = append(u.Accounts, newAccount)
	u.TotalBalance += newAccount.Balance

	targetBank.Accounts = append(targetBank.Accounts, *newAccount)

	fmt.Printf("Account created successfully for %s %s | Account No: %d | Balance: %.2f | Bank: %s\n",
		u.FirstName, u.LastName, newAccount.AccountNo, newAccount.Balance, targetBank.FullName)
	return nil
}

func (u *User) CalculateTotalBalance() (float32, error) {
	if u.IsAdmin {
		return 0, errors.New("only customers can see Total Balance")
	}
	total := float32(0)
	for _, acc := range u.Accounts {
		total += acc.Balance
	}
	u.TotalBalance = total
	return total, nil
}

func (u *User) GetMyAccountBlance(accountNo int) (float32, error) {
	if u.IsAdmin {
		return 0, errors.New("only valid customer can see account balance")
	}
	for i := range u.Accounts {
		if u.Accounts[i].AccountNo == accountNo {
			return u.Accounts[i].Balance, nil
		}
	}
	return 0, errors.New("provided account number is not related to requesting user")
}

func (u *User) GetSelfAccountById(accountNo int) (*accounts.Accounts, error) {
	if u.IsAdmin {
		return nil, nil
	}

	for i := range u.Accounts {
		if u.Accounts[i].AccountNo == accountNo {
			return u.Accounts[i], nil
		}
	}
	return nil, fmt.Errorf("account %d not found for user %s %s", accountNo, u.FirstName, u.LastName)
}

func (u *User) TransferBetweenSelfAccounts(fromAccNo, toAccNo int, amount float32) error {
	if u.IsAdmin {
		return errors.New("admin cannot perform transfers")
	}

	fromAcc, err := u.GetSelfAccountById(fromAccNo)
	if err != nil {
		return fmt.Errorf("source account retrieval failed: %v", err)
	}

	toAcc, err := u.GetSelfAccountById(toAccNo)
	if err != nil {
		return fmt.Errorf("target account retrieval failed: %v", err)
	}

	if fromAcc.AccountNo == toAcc.AccountNo {
		return errors.New("cannot transfer to the same account")
	}

	if err := fromAcc.SelfTransfer(amount, toAcc); err != nil {
		return fmt.Errorf("transfer between accounts failed: %v", err)
	}

	fmt.Printf("Successfully transferred Rs.%.2f from Acc#%d to Acc#%d\n", amount, fromAccNo, toAccNo)
	return nil
}

func (u *User) TransferToOtherUser(fromAccNo, targetAccNo int, amount float32) error {
	if u.IsAdmin {
		return errors.New("admin cannot perform transfers")
	}

	fromAcc, err := u.GetSelfAccountById(fromAccNo)
	if err != nil {
		return fmt.Errorf("source account retrieval failed: %v", err)
	}

	if fromAcc == nil {
		return fmt.Errorf("source account %d not found for user %s %s", fromAccNo, u.FirstName, u.LastName)
	}

	if err := fromAcc.BankTransfer(amount, targetAccNo); err != nil {
		return fmt.Errorf("transfer to other user failed: %v", err)
	}

	fmt.Printf("Successfully transferred Rs.%.2f from your Acc#%d to target Acc#%d\n", amount, fromAccNo, targetAccNo)
	return nil
}

func (u *User) WithdrawFromAccount(accountNo int, amount float32) error {
	if u.IsAdmin {
		return errors.New("admin cannot perform withdrawals")
	}

	for _, acc := range u.Accounts {
		if acc.AccountNo == accountNo {
			return acc.Withdraw(amount)
		}
	}

	return errors.New("account not found for this user")
}

func (u *User) DepositToAccount(accountNo int, amount float32) error {
	if u.IsAdmin {
		return errors.New("admin cannot perform deposits")
	}

	for _, acc := range u.Accounts {
		if acc.AccountNo == accountNo {
			return acc.Deposit(amount)
		}
	}

	return errors.New("account not found for this user")
}

// ====================================================================== Passbook entries with pagination ===================================================================

func (u *User) ViewMyPassbook(accountNo, page, pageSize int) ([]accounts.PassbookEntry, error) {
	if u.IsAdmin {
		return nil, errors.New("admin cannot view passbook as customer")
	}

	acc, err := u.GetSelfAccountById(accountNo)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve account: %v", err)
	}

	return acc.GetPassbook(page, pageSize)
}

func (u *User) ViewAccountSpecificPassbook(accountNo int, page, pageSize int) ([]accounts.PassbookEntry, error) {
	if !u.IsAdmin {
		return nil, errors.New("only admin can view other users' passbooks")
	}

	targetAcc, err := accounts.GetReceiverAccountById(accountNo)
	if err != nil {
		return nil, fmt.Errorf("account retrieval failed: %v", err)
	}

	return targetAcc.GetPassbook(page, pageSize)
}

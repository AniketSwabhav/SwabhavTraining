package user

import (
	"banking_app/accounts"
	"banking_app/bank"
	"banking_app/passbook"
	"banking_app/util"
	"fmt"
)

func NewAdmin(firstName, lastName string) *User {

	defer util.HandlePanic()

	newAdmin, err := NewUser(firstName, lastName, true)
	if err != nil {
		panic(err)
	}
	return newAdmin
}

// ======================================================================Bank Related Methods=======================================================================

func (u *User) AddBank(fullName string) *bank.Bank {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admin can create a bank")
	}
	if !u.IsActive {
		panic("admin needs to be active to create a bank")
	}
	newBank, err := bank.NewBank(fullName)
	if err != nil {
		panic(err)
	}
	return newBank
}

func (u *User) GetBankById(bankId int) *bank.Bank {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admin can create a bank")
	}
	if !u.IsActive {
		panic("admin needs to be active to get  banks")
	}
	bank, err := bank.GetBank(bankId)
	if err != nil {
		panic(err)
	}

	return bank
}

func (u *User) GetAllBanks() []bank.Bank {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admins can get all banks")
	}
	if !u.IsActive {
		panic("admin needs to be active to get all banks")
	}
	allbank := bank.GetAllBanks()

	return allbank
}

func (u *User) UpdateBankById(bankId int, param string, value interface{}) {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admin can add banks")
	}
	if !u.IsActive {
		panic("admin needs to be active to update banks")
	}
	if len(u.Banks) == 0 {
		panic("no banks associated with this admin")
	}
	bankToBeUpdated := u.GetBankById(bankId)

	bankToBeUpdated.UpdateBank(param, value)
}

func (u *User) DeleteBankById(bankId int) {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admin can add customers")
	}
	if !u.IsActive {
		panic("admin needs to be active to delete a banks")
	}
	bank.DeleteBank(bankId)

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
}

// =====================================================================Customer Related Methods========================================================================

func (u *User) NewCustomer(firstName, lastName string) *User {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admin can add customers")
	}
	if !u.IsActive {
		panic("admin needs to be active to add customer")
	}
	newCustomer, err := NewUser(firstName, lastName, false)
	if err != nil {
		panic(err)
	}
	return newCustomer
}

func (u *User) GetCustomerById(customerId int) *User {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admin can Perform CRUD on customers")
	}
	if !u.IsActive {
		panic("admin needs to be active to get customer")
	}
	customer, exists := userMap[customerId]
	if !exists {
		panic("customer not found with the given ID")
	}
	return customer
}

func (u *User) GetAllUsers() []User {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admin can see all banks")
	}
	if !u.IsActive {
		panic("inactive Admin cannot see all banks")
	}
	totalUsers := []User{}
	for _, user := range userMap {
		totalUsers = append(totalUsers, *user)
	}
	return totalUsers
}

func (u *User) UpdateCustomerById(customerId int, param string, value interface{}) {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admin can Perform CRUD on customers")
	}
	if !u.IsActive {
		panic("inactive Admin cannot update banks")
	}
	if customerId < 0 {
		panic("customerId cannot be negative")
	}
	targetCustomer := u.GetCustomerById(customerId)

	if param == "" {
		panic("parameter cannot be empty")
	}

	switch param {
	case "FirstName":
		targetCustomer.updateFirstName(value)
	case "LastName":
		targetCustomer.updateLastName(value)
	default:
		panic("invalid parameter for update")
	}
}

func (target *User) updateFirstName(value interface{}) {

	defer util.HandlePanic()

	strVal, ok := value.(string)
	if !ok || strVal == "" {
		panic("value is empty, provide valid value")
	}
	target.FirstName = strVal
	fmt.Println("First name updated successfully")
}

func (target *User) updateLastName(value interface{}) {

	defer util.HandlePanic()

	strVal, ok := value.(string)
	if !ok || strVal == "" {
		panic("value is empty, provide valid value")
	}
	target.LastName = strVal
	fmt.Println("Last name updated successfully")
}

func (u *User) DeleteCustomerById(customerId int) {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admin can Perform CRUD on customers")
	}
	if !u.IsActive {
		panic("inactive Admin cannot delete customer")
	}
	if customerId < 0 {
		panic("customerId cannot be negative")
	}

	_, exists := userMap[customerId]
	if !exists {
		panic("customer not found with the given ID")
	}

	delete(userMap, customerId)
}

func (u *User) ViewAccountSpecificPassbook(accountNo int, page, pageSize int) []passbook.Transaction {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admin can view other users' passbooks")
	}
	if !u.IsActive {
		panic("admin needs to be active to View Account Specific Passbook")
	}

	targetAcc, err := accounts.GetReceiverAccountById(accountNo)
	if err != nil {
		panic(err)
	}
	passbook, err := targetAcc.GetPassbook(page, pageSize)
	if err != nil {
		panic(err)
	}

	return passbook
}

func (u *User) GetBankTransactionAmount(bankAId, bankBId int) float32 {

	defer util.HandlePanic()

	if !u.IsAdmin {
		panic("only admin can do bank settlement")
	}
	if !u.IsActive {
		panic("admin needs to be active to Get Bank Transaction Amount")
	}

	bankA := u.GetBankById(bankAId)

	amount, err := bankA.GetBankTransactionAmount(bankBId)
	if err != nil {
		panic(err)
	}

	return amount
}

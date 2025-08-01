package main

import (
	"banking_app/user"
	"fmt"
)

func main() {

	admin := user.NewAdmin("Super", "Admin")
	customer := admin.NewCustomer("Aniket", "Pardeshi")
	fmt.Println(*admin)
	fmt.Println(*customer)

	customer2 := admin.NewCustomer("ankush", "Sondal")
	fmt.Println(*customer2)

	customer3 := admin.NewCustomer("Brijesh", "Mavani")
	fmt.Println(*customer3)

	bank1 := admin.AddBank("Bank of Baroda")
	fmt.Println(*bank1)
	bank2 := admin.AddBank("Kotak Mahindra")
	fmt.Println(*bank2)
	bank3 := admin.AddBank("Punjab National Bank")
	fmt.Println(*bank3)

	customer.CreateAccount(1)
	customer.CreateAccount(1)
	customer.CreateAccount(1)
	customer2.CreateAccount(2)
	customer3.CreateAccount(3)

	customer.TransferBetweenSelfAccounts(101, 103, 300)
	customer.TransferBetweenSelfAccounts(102, 103, 450)
	customer.TransferBetweenSelfAccounts(103, 101, 300)
	customer.TransferBetweenSelfAccounts(101, 103, 200)

	// fmt.Println(customer.CalculateTotalBalance())
	// fmt.Println(customer.GetMyAccountBlance(101))

	customer.TransferToOtherUser(101, 104, 300)
	customer.TransferToOtherUser(102, 104, 200)
	customer2.TransferToOtherUser(104, 101, 250)
	customer.TransferToOtherUser(101, 105, 379)

	// fmt.Println(customer.CalculateTotalBalance())
	// fmt.Println(customer2.CalculateTotalBalance())

	// fmt.Println(customer.GetMyAccountBlance(101))

	customer.DepositToAccount(101, 500)

	// fmt.Println(customer.GetMyAccountBlance(101))

	customer.DepositToAccount(101, 600)
	customer.DepositToAccount(101, 10)
	customer.WithdrawFromAccount(101, 400)
	// fmt.Println(customer.GetMyAccountBlance(101))

	customer.DepositToAccount(101, 450)
	customer.WithdrawFromAccount(101, 300)

	// fmt.Println(customer.GetMyAccountBlance(101))

	customer.WithdrawFromAccount(101, 200)
	// fmt.Println(customer.GetMyAccountBlance(101))

	entries := customer.ViewMyPassbook(101, 1, 5)
	for _, entry := range entries {
		fmt.Printf("%s | %s | Rs.%.2f | Balance: Rs.%.2f | Note: %s\n",
			entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Type, entry.Amount, entry.Balance, entry.Note)
	}

	fmt.Println("=====================================================================================================")

	entries2 := admin.ViewAccountSpecificPassbook(102, 1, 5)
	for _, entry := range entries2 {
		fmt.Printf("%s | %s | Rs.%.2f | Balance: Rs.%.2f | Note: %s\n",
			entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Type, entry.Amount, entry.Balance, entry.Note)
	}

	fmt.Println("=====================================================================================================")

	// fmt.Println(admin.GetBankById(1))
	// fmt.Println(admin.GetBankById(2))

	amount := admin.GetBankTransactionAmount(1, 3)
	fmt.Println(amount)

	allbanks := admin.GetAllBanks()
	for _, bank := range allbanks {
		fmt.Printf("Bank Id : %d\n", bank.BankID)
		fmt.Printf("FullName: %s\n", bank.FullName)
		fmt.Printf("Abbreviation %s\n", bank.Abbreviation)
		fmt.Println("---------------------------------------------------------")
	}

	fmt.Println("=====================================================================================================")

	users := admin.GetAllUsers()
	for _, user := range users {
		fmt.Printf("User Id : %d\n", user.UserID)
		fmt.Printf("First Name: %s\n", user.FirstName)
		fmt.Printf("Last Name: %s\n", user.LastName)
		fmt.Printf("TotalBalance: %.2f\n", user.TotalBalance)
		fmt.Println("---------------------------------------------------------")
	}

}

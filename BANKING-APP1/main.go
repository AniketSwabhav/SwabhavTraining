package main

import (
	"banking_app/user"
	"fmt"
)

func main() {

	admin, _ := user.NewAdmin("Super", "Admin")
	customer, _ := admin.NewCustomer("Aniket", "Pardeshi")
	fmt.Println(*admin)
	fmt.Println(*customer)

	customer2, _ := admin.NewCustomer("ankush", "Sondal")
	fmt.Println(*customer2)

	bank1, err := admin.AddBank("Bank of Baroda")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*bank1)
	bank2, err := admin.AddBank("Kotak Mahindra")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*bank2)

	customer.CreateAccount(1)
	customer.CreateAccount(1)
	customer.CreateAccount(1)
	customer2.CreateAccount(2)

	customer.TransferBetweenSelfAccounts(101, 103, 300)
	customer.TransferBetweenSelfAccounts(102, 103, 450)
	customer.TransferBetweenSelfAccounts(103, 101, 300)
	customer.TransferBetweenSelfAccounts(101, 103, 200)

	// fmt.Println(customer.CalculateTotalBalance())
	// fmt.Println(customer.GetMyAccountBlance(101))

	customer.TransferToOtherUser(101, 104, 300)
	customer.TransferToOtherUser(102, 104, 200)
	customer2.TransferToOtherUser(104, 101, 250)

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

	entries, err := customer.ViewMyPassbook(101, 1, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, entry := range entries {
			fmt.Printf("%s | %s | Rs.%.2f | Balance: Rs.%.2f | Note: %s\n",
				entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Type, entry.Amount, entry.Balance, entry.Note)
		}
	}

	fmt.Println("=====================================================================================================")

	entries2, err := admin.ViewAccountSpecificPassbook(102, 1, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, entry := range entries2 {
			fmt.Printf("%s | %s | Rs.%.2f | Balance: Rs.%.2f | Note: %s\n",
				entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Type, entry.Amount, entry.Balance, entry.Note)
		}
	}
}

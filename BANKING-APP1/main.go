package main

import (
	"banking_app/user"
	"fmt"
)

func main() {

	admin, _ := user.NewAdmin("Super", "Admin", true)
	customer, _ := admin.NewCustomer("Aniket", "Pardeshi", false)
	fmt.Println(*customer)

	customer2, err := admin.NewCustomer("ankush", "Sondal", false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*customer2)

	bank1, err := admin.AddBank("Bank of Baroda")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*bank1)
	bank2, err := admin.AddBank("Bank of Baroda")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*bank2)

	customer.CreateAccount(1)
	customer.CreateAccount(1)
	customer.CreateAccount(1)
	customer2.CreateAccount(1)

	customer.TransferBetweenSelfAccounts(101, 102, 500)
	fmt.Println(customer.CalculateTotalBalance())

	customer.TransferToOtherUser(101, 104, 300)
	fmt.Println(customer.CalculateTotalBalance())
	fmt.Println(customer2.CalculateTotalBalance())

	fmt.Println(customer.GetMyAccountBlance(101))

	customer.DepositToAccount(101, 500)
	fmt.Println(customer.GetMyAccountBlance(101))

	customer.WithdrawFromAccount(101, 200)
	fmt.Println(customer.GetMyAccountBlance(101))
}

package main

import (
	helper "Contact_app/Helper"
	"fmt"
)

func main() {
	admin := helper.NewAdmin("Aniket", "Pardeshi")
	fmt.Println(*admin)

	staff := admin.NewStaff("Anii", "Pardeshi")
	ap := staff.CreateContact("Ankush", "Sondal")
	hh := ap.CreateContactDetail("dumm", "mmy")

	staff2 := admin.NewStaff("Yash", "Shan")
	ys := staff2.CreateContact("Yash2", "Shah")
	gg := ys.CreateContactDetail("hhh", "hhh")

	staff3 := admin.NewStaff("adhah", "Shan")
	ya := staff3.CreateContact("dcscs", "Shcsdcsdah")
	jj := ya.CreateContactDetail("hssshh", "hhsssh")

	// admin.UpdateUserById(2, "FName", "Aniket")
	// admin.UpdateUserById(2, "LName", "Pardeshi")
	// fmt.Println(*staff)
	// admin.GetUserById(2)
	// admin.deleteUserById(2)
	// admin.getAllUsers()

	// staff.updateContactById(1, "FName", "Prathamesh")

	fmt.Println(ap)
	fmt.Println(*hh)

	fmt.Println(*ys)
	fmt.Println(*gg)

	fmt.Println(*ya)
	fmt.Println(*jj)
}

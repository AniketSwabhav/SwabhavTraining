package main

import (
	helper "Contact_app/Helper"
	"fmt"
)

func main() {
	admin := helper.NewAdmin("Aniket", "hjjhghj")
	fmt.Println(*admin)

	staff := admin.NewStaff("Anii", "Pardeshi")
	staff2 := admin.NewStaff("Anii2", "Legened")

	// admin.UpdateUserById(1, "FName", "Aniket")
	// admin.UpdateUserById(1, "LName", "Pardeshi")
	// admin.UpdateUserById(1, "IsAdmin", true)
	// admin.UpdateUserById(1, "IsActive", false)
	// admin.UpdateUserById(1, "IsActive", true) //Extra

	// fmt.Println(admin.GetUserById(1))
	// admin.DeleteUserById(1)
	// fmt.Println(admin.GetUserById(2))
	allUsers := admin.GetAllUsers()
	fmt.Println("===== All Users =====")
	for _, user := range allUsers {
		fmt.Printf("UserID: %d, Name: %s %s, IsAdmin: %t, IsActive: %t\n",
			user.UserID, user.FName, user.LName, user.IsAdmin, user.IsActive)
	}
	fmt.Println("=====================")
	staff.CreateContact("Anku", "Son")
	staff.CreateContact("Yashodip", "Mahajan")
	staff.CreateContact("Prathamesh", "Lokare")
	staff2.CreateContact("Anupam", "Singh")
	staff2.CreateContact("Bhushan", "Gavand")

	// staff.UpdateContactById(1, "FName", "Ankush")
	// staff.UpdateContactById(1, "LName", "Sondal")
	// staff.UpdateContactById(1, "IsActive", false) //Extra
	// fmt.Println(staff.GetContactById(1))
	// admin.DeleteContactById(1)
	// fmt.Println(staff.GetContactById(1))
	// fmt.Println("------------------------------------")
	// fmt.Println(admin.GetAllUserContacts(2))
	allContacts := staff.GetAllContactsOfAllUsers()
	fmt.Println("===== All Contacts of All Users =====")
	if len(allContacts) == 0 {
		fmt.Println("No contacts found.")
	} else {
		for _, c := range allContacts {
			fmt.Printf("ContactID: %d, Name: %s %s, IsActive: %t\n", c.ContactID, c.FName, c.LName, c.IsActive)
		}
	}
	fmt.Println("=================================")

	staff.CreateContactDetailByContactId(1, "Gmail", "aniket@gmail.com")
	staff.CreateContactDetailByContactId(1, "no.", "9834985338")
	staff.CreateContactDetailByContactId(2, "no.", "9834985338")
	staff.CreateContactDetailByContactId(2, "no.", "9834985338")

	staff2.CreateContactDetailByContactId(1, "Gmail", "aniket@swabhavtechlabs.com")
	staff2.CreateContactDetailByContactId(2, "Login no.", 102)
	staff2.CreateContactDetailByContactId(2, "no.", "9834985338")

	// staff.UpdateContactDetailById(1, 2, "Type", "Mobile no.")
	// staff2.UpdateContactDetailById(1, 1, "Value", "anii@gmail.com`")

	// staff.UpdateContactDetailById(2, 1, "Type", "Phone no.")
	// staff.UpdateContactDetailById(2, 2, "Type", "Mobile no.")
	// staff.UpdateContactDetailById(2, 2, "Value", "88282828282")

	// staff.DeleteContactDetailsById(1, 1)
	// fmt.Println(staff.GetContactDetailsById(1, 1))

	contactDetails := staff.GetAllContactDetailsOfAllUsers()
	fmt.Println("===== All Contact Details of All Users =====")
	for _, detail := range contactDetails {
		fmt.Printf("DetailID: %d, Type: %s, Value: %v\n",
			detail.ContactDetailID, detail.Type, detail.Value)
	}

}

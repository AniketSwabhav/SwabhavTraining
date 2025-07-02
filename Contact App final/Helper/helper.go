package helper

import user "Contact_app/User"

// Creating Admin User
func NewAdmin(fName, lName string) *user.User {
	newAdmin := user.NewUser(fName, lName, true)
	return newAdmin
}

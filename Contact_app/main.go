package main

import (
	"fmt"
)

var UserId int = 1
var ContactId int = 0
var ContactDetailId int = 0
var userMap = make(map[int]*User)
var contactMap = make(map[int]*Contact)
var contactDetailMap = make(map[int]*ContactDetail)

// structs

type User struct {
	UserID   int
	FName    string
	LName    string
	IsAdmin  bool
	IsActive bool
	Contacts []*Contact
}

type Contact struct {
	ContactID      int
	FName          string
	LName          string
	IsActive       bool
	ContactDetails []*ContactDetail
}

type ContactDetail struct {
	ContactDetailID int
	Type            string
	Value           interface{}
}

// Factory Methods

func newUser(fName, lName string, isAdmin, isActive bool) *User {
	if fName == "" && lName == "" {
		return nil
	}

	UserId = UserId + 1
	user := &User{
		UserID:   UserId,
		FName:    fName,
		LName:    lName,
		IsAdmin:  isAdmin,
		IsActive: isActive,
		Contacts: []*Contact{},
	}

	userMap[UserId] = user

	role := "Staff"
	if user.IsAdmin {
		role = "Admin"
	}
	fmt.Printf("User Created with UserId : %d (%s)\n", UserId, role)
	return user
}

func newContact(user *User, fName, lName string, isActive bool) *Contact {
	if fName == "" && lName == "" {
		return nil
	}

	ContactId = ContactId + 1
	// ContactId := len(u.Contacts) + 1
	contact := &Contact{
		ContactID: ContactId,
		FName:     fName,
		LName:     lName,
		IsActive:  isActive,
	}

	contactMap[ContactId] = contact
	fmt.Printf("user: %d created the contact with Id: %d\n", user.UserID, ContactId)
	return contact
}

func newContactDetail(contact *Contact, detailType string, value interface{}) *ContactDetail {
	if contact == nil {
		return nil
	}

	ContactDetailId = ContactDetailId + 1

	contactDetail := &ContactDetail{
		ContactDetailID: ContactDetailId,
		Type:            detailType,
		Value:           value,
	}

	contactDetailMap[ContactDetailId] = contactDetail
	fmt.Printf("ContactDetailId: %d inserted in contactId : %d\n", ContactDetailId, contact.ContactID)
	return contactDetail
}

// Switch Case methods for calling update methods.
func (u *User) updateUser(target *User, param string, value interface{}) {
	switch param {
	case "FName":
		u.updateUserFirstName(target, value)
	case "LName":
		u.updateUserLastName(target, value)
	case "IsAdmin":
		u.updateIsAdminStatus(target, value)
	case "IsActive":
		u.updateUserIsActiveStatus(target, value)
	}
}

func (u *User) updateContact(target *Contact, param string, value interface{}) {
	switch param {
	case "FName":
		u.updateContactFirstName(target, value)
	case "LName":
		u.updateContactLastName(target, value)
	case "IsActive":
		u.updateContactIsActiveStatus(target, value)
	}
}

func (u *User) updateContactDetail(target *ContactDetail, param string, value interface{}) {
	switch param {
	case "Type":
		u.updateContactDetailType(target, value)
	case "Value":
		u.updateContactDetailValue(target, value)
	}
}

// User Methods

// IsAdmin and IsActive status check method for user
func (u *User) checkForisAdminAndIsActiveForUser() bool {
	if !u.IsAdmin {
		fmt.Println("Only Admin can perform CRUD on Users.")
		return false
	}
	if !u.IsActive {
		fmt.Println("Inactive Admin cannot perform CRUD on Users.")
		return false
	}
	return true
}

// Create Method for User Entity
func (u *User) createUser(fName, lName string, isAdmin, isActive bool) *User {
	if !u.checkForisAdminAndIsActiveForUser() {
		return nil
	}

	role := "Staff"
	if u.IsAdmin {
		role = "Admin"
	}
	fmt.Printf("userId: %d (%s) is creating a user\n", u.UserID, role)
	return newUser(fName, lName, isAdmin, isActive)
}

// Update methods for User Entity
func (u *User) updateUserFirstName(target *User, value interface{}) {
	if !u.checkForisAdminAndIsActiveForUser() {
		return
	}
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateFirstName: invalid string")
		return
	}
	target.FName = strVal
	fmt.Println("First name updated successfully.")
}

func (u *User) updateUserLastName(target *User, value interface{}) {
	if !u.checkForisAdminAndIsActiveForUser() {
		return
	}
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateFirstName: invalid string")
		return
	}
	target.LName = strVal
	fmt.Println("Last name updated successfully.")
}

func (u *User) updateIsAdminStatus(target *User, value interface{}) {
	if !u.checkForisAdminAndIsActiveForUser() {
		return
	}
	status, ok := value.(bool)
	if !ok {
		fmt.Println("updateIsAdminStatus: invalid status")
		return
	}
	target.IsAdmin = status
	fmt.Printf("IsAdmin status changed to: %t\n", status)
}

func (u *User) updateUserIsActiveStatus(target *User, value interface{}) {
	if !u.checkForisAdminAndIsActiveForUser() {
		return
	}
	status, ok := value.(bool)
	if !ok {
		fmt.Println("updateIsActiveStatus: invalid status")
		return
	}
	target.IsActive = status
	fmt.Printf("IsActive status changed to: %t\n", status)
}

// Getter Methods for User Entity
func (u *User) getAllUsers() []*User {
	if !u.checkForisAdminAndIsActiveForUser() {
		return nil
	}
	var users []*User
	for _, user := range userMap {
		// if user.IsActive {
		// 	users = append(users, user)
		// }
		users = append(users, user)
	}
	return users
}

func (u *User) getUserById(userId int) *User {
	if !u.checkForisAdminAndIsActiveForUser() {
		return nil
	}
	if singleUser, exists := userMap[userId]; exists {
		fmt.Printf("User data fetched with id: %d\n", userId)
		return singleUser
	}
	fmt.Printf("User Not Present with id: %d\n", userId)
	return nil
}

// Delete method for User Entity
func (u *User) deleteUserById(userId int) bool {
	if !u.checkForisAdminAndIsActiveForUser() {
		return false
	}
	if user, exists := userMap[userId]; exists {
		user.IsActive = false
		fmt.Printf("User with userId %d marked as inactive.\n", userId)
		return true
	}
	fmt.Printf("User not present with userId : %d\n", userId)
	return false
}

// Contact Methods
// IsAdmin and IsActive status check method for Contact
func (u *User) checkForisAdminAndIsActiveForContact() bool {
	if u.IsAdmin {
		fmt.Println("Only Staff can perform CRUD on Contact")
		return false
	}
	if !u.IsActive {
		fmt.Println("Inactive staff cannot perform CRUD on contacts.")
		return false
	}
	return true
}

// Create method for Contact Entity
func (u *User) createContact(user *User, fName, lName string, isActive bool) *Contact {
	if !u.checkForisAdminAndIsActiveForContact() {
		return nil
	}

	return newContact(user, fName, lName, isActive)
}

// Update methods for Contact Entity
func (u *User) updateContactFirstName(target *Contact, value interface{}) {
	if !u.checkForisAdminAndIsActiveForContact() {
		return
	}
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateContactFirstName: invalid string")
		return
	}
	target.FName = strVal
	fmt.Println("First name updated successfully.")
}

func (u *User) updateContactLastName(target *Contact, value interface{}) {
	if !u.checkForisAdminAndIsActiveForContact() {
		return
	}
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateContactLastName: invalid string")
		return
	}
	target.LName = strVal
	fmt.Println("Last name updated successfully.")
}

func (u *User) updateContactIsActiveStatus(target *Contact, value interface{}) {
	if !u.checkForisAdminAndIsActiveForContact() {
		return
	}
	status, ok := value.(bool)
	if !ok {
		fmt.Println("updateUserIsActiveStatus: invalid status")
		return
	}
	target.IsActive = status
	fmt.Printf("IsActive status changed to: %t\n", status)
}

// Getter Methods for Contact Entity
func (u *User) getAllContacts() []*Contact {
	if !u.checkForisAdminAndIsActiveForContact() {
		return nil
	}
	var contacts []*Contact
	for _, contact := range contactMap {
		// if contact.IsActive {
		// 	contacts = append(contacts, contact)
		// }
		contacts = append(contacts, contact)
	}
	return contacts
}

func (u *User) getContactById(contactId int) *Contact {
	if !u.checkForisAdminAndIsActiveForContact() {
		return nil
	}
	if singleContact, exists := contactMap[contactId]; exists {
		fmt.Printf("Contact data fetched with ContactId: %d\n", contactId)
		return singleContact
	}
	fmt.Printf("contact Not Present with ContactId: %d\n", contactId)
	return nil
}

// Delete method for Contact Entity
func (u *User) deleteContactById(contactId int) bool {
	if !u.checkForisAdminAndIsActiveForContact() {
		return false
	}
	if contact, exists := contactMap[contactId]; exists {
		contact.IsActive = false
		fmt.Printf("Contact with contactId %d marked as inactive.\n", contactId)
		return true
	}
	fmt.Printf("Contact not present with contactId : %d\n", contactId)
	return false
}

// Contact Details Methods
// IsAdmin and IsActive status check method for Contact Details Entity
func (u *User) checkForisAdminAndIsActiveForContactDetail() bool {
	if u.IsAdmin {
		fmt.Println("Only Staff can perform CRUD on Contact")
		return false
	}
	if !u.IsActive {
		fmt.Println("Inactive staff cannot perform CRUD on contacts.")
		return false
	}
	return true
}

// Create method for Contact Details Entity
func (u *User) createContactDetail(contact *Contact, detailType string, value interface{}) *ContactDetail {
	if !u.checkForisAdminAndIsActiveForContactDetail() {
		return nil
	}

	return newContactDetail(contact, detailType, value)
}

// Update methods for Contact Details Entity
func (u *User) updateContactDetailType(target *ContactDetail, value interface{}) {
	if !u.checkForisAdminAndIsActiveForContactDetail() {
		return
	}
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateContactDetailType: invalid string")
		return
	}
	target.Type = strVal
	fmt.Println("Contact Detail Type updated successfully.")
}

func (u *User) updateContactDetailValue(target *ContactDetail, value interface{}) {
	if !u.checkForisAdminAndIsActiveForContactDetail() {
		return
	}
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateContactDetailValue: invalid string")
		return
	}
	target.Value = strVal
	fmt.Println("Contact Detail Value updated successfully.")
}

// Getter mthod for Contact Detail Entity
func (u *User) getAllContactDetails() *[]ContactDetail {
	if !u.checkForisAdminAndIsActiveForContactDetail() {
		return nil
	}
	var contactDetails []ContactDetail
	for _, contactDetail := range contactDetailMap {
		contactDetails = append(contactDetails, *contactDetail)
	}
	return &contactDetails
}

func (u *User) getContactDetailById(contactDetailId int) *ContactDetail {
	if !u.checkForisAdminAndIsActiveForContactDetail() {
		return nil
	}
	if singleContactDetail, exists := contactDetailMap[contactDetailId]; exists {
		fmt.Printf("Contact Detail fetched with contactDetailId: %d\n", contactDetailId)
		return singleContactDetail
	}
	fmt.Printf("contact Detail Not Present with contactDetailId: %d\n", contactDetailId)
	return nil
}

// Delete method for Contact Detail
func (u *User) deleteContactDetailById(contactDetailId int) bool {
	if !u.checkForisAdminAndIsActiveForContactDetail() {
		return false
	}
	if _, exists := contactDetailMap[contactDetailId]; exists {
		// Mark as deleted or remove from map as per your design.
		delete(contactDetailMap, contactDetailId)
		fmt.Printf("ContactDetail deleted with contactDetailId: %d\n", contactDetailId)
		return true
	}
	fmt.Printf("ContactDetail not present with contactDetailId: %d\n", contactDetailId)
	return false
}

func main() {

	admin := &User{
		UserID:   1,
		FName:    "Admin",
		LName:    "User",
		IsAdmin:  true,
		IsActive: true,
	}

	// adminUser := admin.createUser("Admin", "User", true, true)
	// fmt.Println(*adminUser)
	// fmt.Println("---------------------------------------------------")

	// adminUser.updateUser(adminUser, "FName", "Super")
	// fmt.Println(*adminUser)

	// adminUser.updateUser(adminUser, "LName", "Admin")
	// fmt.Println(*adminUser)

	// adminUser.updateUser("IsAdmin", false)
	// fmt.Println(*adminUser)

	// adminUser.updateUser("IsActive", false)
	// fmt.Println(*adminUser)

	staffUser := admin.createUser("Staff", "User", false, true)
	fmt.Println(*staffUser)
	fmt.Println("---------------------------------------------------")

	// adminUser.updateUser(staffUser, "FName", "Aniket")
	// fmt.Println(*staffUser)

	// adminUser.updateUser(staffUser, "LName", "Pardeshi")
	// fmt.Println(*staffUser)

	// adminUser.updateUser(staffUser, "IsAdmin", true)
	// fmt.Println(*staffUser)

	// staffUser.updateUser(staffUser, "FName", "Abhishek")
	// fmt.Println(*staffUser)
	// staffUser.updateUser(staffUser, "LName", "Pandey")
	// fmt.Println(*staffUser)

	// staffUser2 := admin.createUser("Staff2", "User", false, true)
	// fmt.Println(*staffUser2)
	// fmt.Println("---------------------------------------------------")

	staffUserContact1 := staffUser.createContact(staffUser, "Contact1", "StaffUser", true)
	fmt.Println(*staffUserContact1)
	fmt.Println("---------------------------------------------------")

	// staffUser.updateContact(staffUserContact1, "FName", "Anurag")
	// fmt.Println(*staffUserContact1)

	// staffUser.updateContact(staffUserContact1, "LName", "Kashyap")
	// fmt.Println(*staffUserContact1)

	// staffUser.updateContact(staffUserContact1, "IsActive", false)
	// fmt.Println(*staffUserContact1)

	// staffUser2Contact1 := staffUser2.createContact(staffUser2, "Contact1", "StaffUser2", true)
	// fmt.Println(*staffUser2Contact1)
	// fmt.Println("---------------------------------------------------")

	// staffContact := staffUser.createContact(staffUser, "Contact1", "StaffUser", true)
	// fmt.Println(*staffContact)
	// fmt.Println("---------------------------------------------------")

	contactDetail1 := staffUser.createContactDetail(staffUserContact1, "Mobile", "1234567890")
	fmt.Println(*contactDetail1)
	contactDetail2 := staffUser.createContactDetail(staffUserContact1, "Email", "anii@gmail.com")
	fmt.Println(*contactDetail2)

	// adminContactDetail := admin.createContactDetail(staffContact, "Email", "test@example.com")
	// fmt.Println(adminContactDetail)

	staffUser.updateContactDetail(contactDetail1, "Type", "Mobile No.")
	staffUser.updateContactDetail(contactDetail1, "Value", "9834985338")
	// fmt.Println(*contactDetail1)

	// allUsers := admin.getAllUsers()
	// for _, user := range allUsers {
	// 	fmt.Printf("UserId: %d\n", user.UserID)
	// 	fmt.Printf("First Name: %s\n", user.FName)
	// 	fmt.Printf("Last Name: %s\n", user.LName)
	// 	fmt.Printf("IsAdmin Status: %t\n", user.IsAdmin)
	// 	fmt.Printf("IsActive Status: %t\n", user.IsActive)
	// 	fmt.Println("--------------------------------------")
	// }

	// admin.deleteUserById(2)

	// SingleUserData := admin.getUserById(4)
	// fmt.Println(*SingleUserData)

	// staffUser.deleteContactById(2)

	// allContacts := staffUser.getAllContacts()
	// for _, contact := range allContacts {
	// 	// fmt.Println("userId: %d\n", contact.userID)
	// 	fmt.Printf("contactId: %d\n", contact.ContactID)
	// 	fmt.Printf("First Name: %s\n", contact.FName)
	// 	fmt.Printf("Last Name: %s\n", contact.LName)
	// 	fmt.Printf("IsActive Status: %t\n", contact.IsActive)
	// 	fmt.Println("--------------------------------------")
	// }

	// staffUser.updateContact(staffUser2Contact1, "IsActive", true)

	// singleContact := staffUser.getContactById(2)
	// fmt.Println(*singleContact)

	// allContactDetails := staffUserContact1.getAllContactDetails()

	staffUser.deleteContactDetailById(1)

	allContactDetails := staffUser.getAllContactDetails()
	for _, contactDetail := range *allContactDetails {
		fmt.Printf("ContactDetailId: %d\n", contactDetail.ContactDetailID)
		fmt.Printf("Type: %s\n", contactDetail.Type)
		fmt.Printf("Value: %v\n", contactDetail.Value)
		fmt.Println("--------------------------------------")
	}

	singleContactDetail := staffUser.getContactDetailById(1)
	fmt.Println(*singleContactDetail)

}

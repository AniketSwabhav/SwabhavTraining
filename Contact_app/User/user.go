package user

import (
	contact "Contact_app/Contact"
	contact_detail "Contact_app/Contact_Detail"
	"fmt"
)

var UserId int = 0
var userMap = make(map[int]*User)

type User struct {
	UserID   int
	FName    string
	LName    string
	IsAdmin  bool
	IsActive bool
	Contacts []*contact.Contact
}

// New Factory for User
func NewUser(fName, lName string, isAdmin bool) *User {
	if fName == "" || lName == "" {
		return nil
	}

	UserId = UserId + 1
	user := &User{
		UserID:   UserId,
		FName:    fName,
		LName:    lName,
		IsAdmin:  isAdmin,
		IsActive: true,
		Contacts: []*contact.Contact{},
	}

	userMap[UserId] = user

	role := "Staff"
	if user.IsAdmin {
		role = "Admin"
	}
	fmt.Printf("User Created with UserId : %d (%s)\n", UserId, role)
	return user
}

// Creating Staff User
func (u *User) NewStaff(fName, lName string) *User {
	if !u.CheckForisAdminAndIsActiveForUser() {
		return nil
	}
	newStaff := NewUser(fName, lName, false)
	return newStaff
}

// IsAdmin and IsActive status check method for user
func (u *User) CheckForisAdminAndIsActiveForUser() bool {
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

// New Switch Case method for calling update methods.
func (u *User) UpdateUserById(userId int, param string, value interface{}) {
	if !u.CheckForisAdminAndIsActiveForUser() {
		return
	}

	userToBeUpdated := u.GetUserById(userId)

	switch param {
	case "FName":
		userToBeUpdated.updateUserFirstName(value)
	case "LName":
		userToBeUpdated.updateUserLastName(value)
	case "IsAdmin":
		userToBeUpdated.updateIsAdminStatus(value)
	case "IsActive":
		userToBeUpdated.updateUserIsActiveStatus(value)
	}
}

// New Update methods for User Entity
func (targetUser *User) updateUserFirstName(value interface{}) {
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateFirstName: invalid string")
		return
	}
	targetUser.FName = strVal
	fmt.Println("First name updated successfully.")
}

func (targetUser *User) updateUserLastName(value interface{}) {
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateFirstName: invalid string")
		return
	}
	targetUser.LName = strVal
	fmt.Println("Last name updated successfully.")
}

func (targetUser *User) updateIsAdminStatus(value interface{}) {
	status, ok := value.(bool)
	if !ok {
		fmt.Println("updateIsAdminStatus: invalid status")
		return
	}
	targetUser.IsAdmin = status
	fmt.Printf("IsAdmin status changed to: %t\n", status)
}

func (targetUser *User) updateUserIsActiveStatus(value interface{}) {
	status, ok := value.(bool)
	if !ok {
		fmt.Println("updateIsActiveStatus: invalid status")
		return
	}
	targetUser.IsActive = status
	fmt.Printf("IsActive status changed to: %t\n", status)
}

func (u *User) GetUserById(userId int) *User {
	if !u.CheckForisAdminAndIsActiveForUser() {
		return nil
	}
	if userId < 0 || userId > len(userMap) {
		return nil
	}
	if singleUser, exists := userMap[userId]; exists {
		return singleUser
	}
	fmt.Printf("User Not Present with id: %d\n", userId)
	return nil
}

// New Getter Methods for User Entity
func (u *User) GetAllUsers() []*User {
	if !u.CheckForisAdminAndIsActiveForUser() {
		return nil
	}
	var users []*User
	for _, user := range userMap {
		users = append(users, user)
	}
	return users
}

// Delete method for User Entity
func (u *User) DeleteUserById(userId int) bool {
	if !u.CheckForisAdminAndIsActiveForUser() {
		return false
	}
	if userId < 0 || userId > len(userMap) {
		return false
	}

	userToBeDeleted := u.GetUserById(userId)

	if userToBeDeleted == nil || !userToBeDeleted.IsActive {
		fmt.Printf("User not present with userId : %d\n", userId)
		return false
	}
	if user, exists := userMap[userId]; exists {
		user.IsActive = false
		fmt.Printf("User with userId %d marked as inactive.\n", userId)
		return true
	}
	return false
}

// IsAdmin and IsActive status check method for Contact
func (u *User) CheckForisAdminAndIsActiveForContact() bool {
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

// New Create method for Contact Entity
func (u *User) CreateContact(fName, lName string) *contact.Contact {
	if !u.CheckForisAdminAndIsActiveForContact() {
		return nil
	}
	ContactId := len(u.Contacts) + 1
	newContact := contact.NewContact(fName, lName, ContactId)
	if newContact == nil {
		return nil
	}
	u.Contacts = append(u.Contacts, newContact)
	return newContact
}

func (u *User) UpdateContactById(Contactid int, param string, value interface{}) {
	if !u.CheckForisAdminAndIsActiveForContact() {
		return
	}
	contactToBeUpdated := u.GetContactById(Contactid)
	contactToBeUpdated.UpdateContact(param, value)
}

func (u *User) GetContactById(contactId int) *contact.Contact {
	if !u.CheckForisAdminAndIsActiveForContact() {
		return nil
	}
	checkId := u.ValidateContactId(contactId)
	if !checkId {
		fmt.Println("Provide a valid contactId")
		return nil
	}
	if contactId < 0 {
		return nil
	}
	for i := 0; i < len(u.Contacts); i++ {
		if contactId == u.Contacts[i].GetId() {
			return u.Contacts[i]
		}
	}
	fmt.Printf("Contact Not Present with ContactId: %d\n", contactId)
	return nil
}

func (U *User) ValidateContactId(contactId int) bool {
	if contactId < 0 || contactId > len(U.Contacts) {
		return false
	}
	return true
}

func (u *User) GetAllUserContacts() []contact.Contact {
	if !u.CheckForisAdminAndIsActiveForContact() {
		return nil
	}
	copyOfUserContacts := []contact.Contact{}
	for _, userContact := range u.Contacts {
		copyOfUserContacts = append(copyOfUserContacts, *userContact)
	}
	return copyOfUserContacts
}

func (u *User) GetAllContactsOfAllUsers() []*contact.Contact {
	if !u.CheckForisAdminAndIsActiveForContact() {
		return nil
	}
	var allContacts []*contact.Contact

	for _, user := range userMap {
		for _, contact := range user.Contacts {
			if contact.IsActive {
				allContacts = append(allContacts, contact)
			}
		}
	}
	return allContacts
}

// Delete method for Contact Entity
func (u *User) DeleteContactById(contactId int) bool {
	if !u.CheckForisAdminAndIsActiveForContact() {
		return false
	}
	if contactId < 0 || contactId >= len(u.Contacts) {
		return false
	}

	contactToBeDeleted := u.GetContactById(contactId)

	if contactToBeDeleted == nil || !contactToBeDeleted.IsActive {
		fmt.Printf("Contact not present with ContactId : %d\n", contactId)
		return false
	}
	contactToBeDeleted.IsActive = false
	fmt.Printf("Contact with contactId %d marked as inactive.\n", contactId)
	return true
}

//Contact Details method

func (u *User) CreateContactDetailByContactId(contactId int, paramType string, value interface{}) {
	if !u.CheckForisAdminAndIsActiveForContact() {
		return
	}
	targetContactForContactDetail := u.GetContactById(contactId)

	targetContactForContactDetail.CreateContactDetail(paramType, value)
}

func (u *User) UpdateContactDetailById(contactId, contactDetailId int, param string, value interface{}) {
	if !u.CheckForisAdminAndIsActiveForContact() {
		return
	}
	if contactId < 0 || contactId >= len(u.Contacts) {
		return
	}
	targetContact := u.GetContactById(contactId)
	targetContact.UpdateContactDetail(contactDetailId, param, value)
}

func (u *User) GetContactDetailsById(contactId, contactDetailId int) *contact_detail.ContactDetail {
	if !u.CheckForisAdminAndIsActiveForContact() {
		return nil
	}
	if contactId < 0 || contactId >= len(u.Contacts) {
		return nil
	}

	targetContact := u.GetContactById(contactId)
	if targetContact == nil {
		return nil
	}

	if !targetContact.ValidateContactDetailsId(contactDetailId) {
		return nil
	}
	resultContactDetail := targetContact.GetContactDetailById(contactDetailId)
	return resultContactDetail
}

// func (u *User) GetAllContactDeatailsOfUser() []contact_detail.ContactDetail {
// 	if !u.CheckForisAdminAndIsActiveForContact() {
// 		return nil
// 	}

// 	var allContactDetails []contact_detail.ContactDetail

// 	for i := 1; i < len(u.Contacts); i++ {
// 		targetContact := u.GetContactById(i)
// 		for i := 1; i < len(targetContact.ContactDetails); i++ {
// 			targetContactDetails := targetContact.GetContactDetailById(i)
// 			allContactDetails = append(allContactDetails, *targetContactDetails)
// 		}
// 	}
// 	return allContactDetails
// }

func (u *User) GetAllContactDetailsOfAllUsers() []contact_detail.ContactDetail {
	if !u.CheckForisAdminAndIsActiveForContact() {
		return nil
	}
	var allContactDetails []contact_detail.ContactDetail

	for _, user := range userMap {
		for _, contact := range user.Contacts {
			if contact.IsActive {
				for _, contactDetail := range contact.ContactDetails {
					allContactDetails = append(allContactDetails, *contactDetail)
				}
			}
		}
	}
	return allContactDetails
}

func (u *User) DeleteContactDetailsById(contactid, contactDDetailId int) error {
	if !u.CheckForisAdminAndIsActiveForContact() {
		return nil
	}
	ContacthavingDetailsSlice := u.GetContactById(contactid)
	ContacthavingDetailsSlice.DeleteContactDetailById(contactDDetailId)
	return nil
}

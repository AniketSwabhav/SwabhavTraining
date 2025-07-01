package contact

import (
	contact_detail "Contact_app/Contact_Detail"
	"fmt"
)

type Contact struct {
	ContactID      int
	FName          string
	LName          string
	IsActive       bool
	ContactDetails []*contact_detail.ContactDetail
}

func NewContact(fName, lName string, contactId int) *Contact {
	if fName == "" && lName == "" {
		return nil
	}

	contact := &Contact{
		ContactID:      contactId,
		FName:          fName,
		LName:          lName,
		IsActive:       true,
		ContactDetails: []*contact_detail.ContactDetail{},
	}

	fmt.Printf("Contact created with ContactId: %d\n", contactId)
	return contact
}

func (c *Contact) UpdateContact(param string, value interface{}) {
	if c.IsActive {
		fmt.Println("Cannot perform CRUD on InActive contacts.")
		return
	}
	switch param {
	case "FName":
		c.UpdateContactFirstName(value)
	case "LName":
		c.UpdateContactLastName(value)
	case "IsActive":
		c.UpdateContactIsActiveStatus(value)
	}
}

// Update methods for Contact Entity
func (targetContact *Contact) UpdateContactFirstName(value interface{}) {
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateContactFirstName: invalid string")
		return
	}
	targetContact.FName = strVal
	fmt.Println("First name updated successfully.")
}

func (targetContact *Contact) UpdateContactLastName(value interface{}) {
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateContactLastName: invalid string")
		return
	}
	targetContact.LName = strVal
	fmt.Println("Last name updated successfully.")
}

func (targetContact *Contact) UpdateContactIsActiveStatus(value interface{}) {
	status, ok := value.(bool)
	if !ok {
		fmt.Println("updateUserIsActiveStatus: invalid status")
		return
	}
	targetContact.IsActive = status
	fmt.Printf("IsActive status changed to: %t\n", status)
}

// Getter Methods for Contact Entity

func (c *Contact) GetId() int {
	if !c.IsActive {
		fmt.Println("Inactive contact cannot perform operations.")
		return -1
	}
	return c.ContactID
}

// Contact Details Methods

// New Create method for Contact Details Entity
func (c *Contact) CreateContactDetail(detailType string, value interface{}) *contact_detail.ContactDetail {
	if !c.IsActive {
		fmt.Println("Cannot add Contact Details in inActive contact")
		return nil
	}
	ContactDetailId := len(c.ContactDetails) + 1
	newContactDetail := contact_detail.NewContactDetail(ContactDetailId, detailType, value)
	if newContactDetail == nil {
		return nil
	}
	c.ContactDetails = append(c.ContactDetails, newContactDetail)
	return newContactDetail
}

func (c *Contact) UpdateContactDetail(contactDetalId int, param string, value interface{}) {
	if !c.IsActive {
		fmt.Println("Cannot add Contact Details in inActive contact")
		return
	}
	targetContactDetail := c.GetContactDetailById(contactDetalId)
	if targetContactDetail == nil {
		return
	}
	targetContactDetail.UpdateContactDetail(param, value)
}

func (c *Contact) GetContactDetailById(contactDetailId int) *contact_detail.ContactDetail {
	if !c.IsActive {
		fmt.Println("Inactive staff cannot perform CRUD on contact Detail.")
		return nil
	}
	if contactDetailId < 0 {
		return nil
	}
	for i := 0; i < len(c.ContactDetails); i++ {
		if contactDetailId == c.ContactDetails[i].GetId() {
			return c.ContactDetails[i]
		}
	}
	fmt.Printf("Contact Detail Not Present with contactDetailId: %d\n", contactDetailId)
	return nil
}

// Delete method for Contact Detail
func (c *Contact) DeleteContactDetailById(contactDetailId int) bool {
	if !c.IsActive {
		fmt.Println("Inactive contact cannot perform CRUD on contact details.")
		return false
	}

	contactDetailToBeDeleted := c.GetContactDetailById(contactDetailId)

	if contactDetailToBeDeleted == nil {
		fmt.Printf("ContactDetail not present with contactDetailId: %d\n", contactDetailId)
		return false
	}

	for i, detail := range c.ContactDetails {
		if detail.ContactDetailID == contactDetailId {
			c.ContactDetails = append(c.ContactDetails[:i], c.ContactDetails[i+1:]...)
			fmt.Printf("ContactDetail with contactDetailId %d deleted successfully.\n", contactDetailId)
			return true
		}
	}

	fmt.Printf("ContactDetail not present with contactDetailId: %d\n", contactDetailId)
	return false
}

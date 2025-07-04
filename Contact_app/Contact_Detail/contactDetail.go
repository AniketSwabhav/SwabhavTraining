package contact_detail

import "fmt"

type ContactDetail struct {
	ContactDetailID int
	Type            string
	Value           interface{}
}

func NewContactDetail(contactDetailId int, detailType string, value interface{}) *ContactDetail {

	if detailType == "" || value == "" {
		return nil
	}
	contactDetail := &ContactDetail{
		ContactDetailID: contactDetailId,
		Type:            detailType,
		Value:           value,
	}
	fmt.Printf("Contact Detail created with ContactDetailId %d\n", contactDetailId)
	return contactDetail
}

// Switch case method for contact Detail
func (contactDetailToBeUpdated *ContactDetail) UpdateContactDetail(param string, value interface{}) {
	switch param {
	case "Type":
		contactDetailToBeUpdated.UpdateContactDetailType(value)
	case "Value":
		contactDetailToBeUpdated.UpdateContactDetailValue(value)
	}
}

// Update methods for Contact Details Entity
func (target *ContactDetail) UpdateContactDetailType(value interface{}) {
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateContactDetailType: invalid string")
		return
	}
	target.Type = strVal
	fmt.Printf("Contact Detail Type updated to : %s\n", strVal)
}

func (target *ContactDetail) UpdateContactDetailValue(value interface{}) {
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("updateContactDetailValue: invalid string")
		return
	}
	target.Value = strVal
	fmt.Printf("Contact Detail Value updated to : %s\n", strVal)
}

func (cd *ContactDetail) GetId() int {
	if cd == nil {
		fmt.Println("ContactDetail is nil")
		return 0
	}
	return cd.ContactDetailID
}

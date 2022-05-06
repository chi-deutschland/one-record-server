package model

type Branch struct {
	ID         			string   	`firestore:"id" json:"id,omitempty"`
	Company 			*Company   	`firestore:"company,omitempty" json:"company,omitempty"`
	ContactPerson   	*Person 		`firestore:"contactPerson,omitempty" json:"contactPerson,omitempty"`
	Location 			*Location  	`firestore:"location,omitempty" json:"location,omitempty"`
	OtherIdentifier 	string   	`firestore:"otherIdentifier,omitempty" json:"otherIdentifier,omitempty"`
}

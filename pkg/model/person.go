package model

type Person struct {
	ID         			string   				`firestore:"id" json:"id,omitempty"`
	associatedBranch 	CompanyBranch   		`firestore:"associatedBranch,omitempty" json:"associatedBranch,omitempty"`
	contact   			string 					`firestore:"contact,omitempty" json:"contact,omitempty"`
	documents 			[]ExternalReference 	`firestore:"documents,omitempty" json:"documents,omitempty"`
	contactType 		string 					`firestore:"contactType,omitempty" json:"contactType,omitempty"`
	department 			string 					`firestore:"department,omitempty" json:"department,omitempty"`
	employeeId 			string 					`firestore:"employeeId,omitempty" json:"employeeId,omitempty"`
	firstName 			string 					`firestore:"firstName,omitempty" json:"firstName,omitempty"`
	jobTitle 			string 					`firestore:"jobTitle,omitempty" json:"jobTitle,omitempty"`
	lastName 			string 					`firestore:"lastName,omitempty" json:"lastName,omitempty"`
	middleName 			string 					`firestore:"middleName,omitempty" json:"middleName,omitempty"`
	salutation 			string 					`firestore:"salutation,omitempty" json:"salutation,omitempty"`
}
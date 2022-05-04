package model

type Person struct {
	ID         			string   				`firestore:"id" json:"id,omitempty"`
	AssociatedBranch 	CompanyBranch   		`firestore:"associatedBranch,omitempty" json:"associatedBranch,omitempty"`
	Contact   			string 					`firestore:"contact,omitempty" json:"contact,omitempty"`
	Documents 			[]ExternalReference 	`firestore:"documents,omitempty" json:"documents,omitempty"`
	ContactType 		string 					`firestore:"contactType,omitempty" json:"contactType,omitempty"`
	Department 			string 					`firestore:"department,omitempty" json:"department,omitempty"`
	EmployeeId 			string 					`firestore:"employeeId,omitempty" json:"employeeId,omitempty"`
	FirstName 			string 					`firestore:"firstName,omitempty" json:"firstName,omitempty"`
	JobTitle 			string 					`firestore:"jobTitle,omitempty" json:"jobTitle,omitempty"`
	LastName 			string 					`firestore:"lastName,omitempty" json:"lastName,omitempty"`
	MiddleName 			string 					`firestore:"middleName,omitempty" json:"middleName,omitempty"`
	Salutation 			string 					`firestore:"salutation,omitempty" json:"salutation,omitempty"`
}
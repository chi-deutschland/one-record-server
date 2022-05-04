package model

type Company struct {
	ID     		string        	`firestore:"id" json:"id,omitempty"`
	Name   		string        	`firestore:"name,omitempty" json:"name,omitempty"`
	Type   		string        	`firestore:"type,omitempty" json:"type,omitempty"`
	Branch 		CompanyBranch 	`firestore:"branch,omitempty" json:"branch,omitempty"`
	Role   		string        	`firestore:"role,omitempty" json:"role,omitempty"`
}

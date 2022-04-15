package model

type Company struct {
	ID     string        `firestore:"id" json:"id,omitempty"`
	Name   string        `firestore:"name" json:"name"`
	Type   string        `firestore:"type" json:"type"`
	Branch CompanyBranch `firestore:"branch" json:"branch"`
	Role   string        `firestore:"role" json:"-"`
}

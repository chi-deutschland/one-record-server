package model

type CompanyBranch struct {
	ID         string   `firestore:"id" json:"id,omitempty"`
	BranchName string   `firestore:"name" json:"branch_name"`
	Location   Location `firestore:"location" json:"location"`
}

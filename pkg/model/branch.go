package model

type CompanyBranch struct {
	ID         string   `firestore:"id" json:"id,omitempty"`
	BranchName string   `firestore:"name,omitempty" json:"branch_name,omitempty"`
	Location   Location `firestore:"location,omitempty" json:"location,omitempty"`
}

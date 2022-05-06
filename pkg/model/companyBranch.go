package model

type CompanyBranch struct {
	ID         string   `firestore:"id" json:"id,omitempty"`
	BranchName string   `firestore:"branchName,omitempty" json:"branchName,omitempty"`
	Location   *Location `firestore:"location,omitempty" json:"location,omitempty"`
}

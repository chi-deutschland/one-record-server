package model

type RegulatedEntity struct {
	ID			string			`firestore:"id" json:"id,omitempty"`
	Entity		*Branch			`firestore:"entity,omitempty" json:"entity,omitempty"`
}
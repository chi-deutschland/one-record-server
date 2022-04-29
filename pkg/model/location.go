package model

type Location struct {
	ID      string  `firestore:"id" json:"id,omitempty"`
	Type    string  `firestore:"type,omitempty" json:"type,omitempty"`
	Name    string  `firestore:"name,omitempty" json:"name,omitempty"`
	Address Address `firestore:"address,omitempty" json:"address,omitempty"`
}

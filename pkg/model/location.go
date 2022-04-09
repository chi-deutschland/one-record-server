package model

type Location struct {
	ID      string  `firestore:"id" json:"id,omitempty"`
	Type    string  `firestore:"type" json:"type"`
	Name    string  `firestore:"name" json:"name"`
	Address Address `firestore:"address" json:"address"`
}

package model

type Address struct {
	ID         string `firestore:"id" json:"id,omitempty"`
	Street     string `firestore:"street" json:"street"`
	PostalCode string `firestore:"postal_code" json:"postal_code"`
	CityName   string `firestore:"city_name" json:"city_name"`
	Country    string `firestore:"country" json:"country"`
}

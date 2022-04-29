package model

type Address struct {
	ID         string `firestore:"id" json:"id,omitempty"`
	Street     string `firestore:"street,omitempty" json:"street,omitempty"`
	PostalCode string `firestore:"postal_code,omitempty" json:"postal_code,omitempty"`
	CityName   string `firestore:"city_name,omitempty" json:"city_name,omitempty"`
	Country    string `firestore:"country,omitempty" json:"country,omitempty"`
}

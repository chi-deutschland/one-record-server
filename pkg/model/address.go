package model

type Address struct {
    ID         string `firestore:"id"                       jsonld:"@id"`
    Type       string `firestore:"type"                     jsonld:"@type" default:"https://onerecord.iata.org/cargo#Address"`
    Street     string `firestore:"street,omitempty"         jsonld:"https://onerecord.iata.org/cargo#address#street"`
    PostalCode string `firestore:"postalCode,omitempty"     jsonld:"https://onerecord.iata.org/cargo#address#postalCode"`
    CityName   string `firestore:"cityName,omitempty"       jsonld:"https://onerecord.iata.org/cargo#address#cityName"`
    Country    string `firestore:"country,omitempty"        jsonld:"https://onerecord.iata.org/cargo#address#country"`
}

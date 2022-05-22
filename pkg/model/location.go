package model

type Location struct {
    ID              string      `firestore:"id"                         jsonld:"@id"`
    Type            string      `firestore:"type,omitempty"             jsonld:"@type" default:"https://onerecord.iata.org/cargo#Location"`
    LocationName    string      `firestore:"locationName,omitempty"     jsonld:"https://onerecord.iata.org/cargo#location#locationName"`
    Address         Address     `firestore:"address,omitempty"          jsonld:"https://onerecord.iata.org/cargo#location#address"`
}
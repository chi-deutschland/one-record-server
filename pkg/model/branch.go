package model

type Branch struct {
    ID              string   `firestore:"id"                         jsonld:"@id"`
    Type            string   `firestore:"type"                       jsonld:"@type" default:"https://onerecord.iata.org/cargo#Branch"`
    Company         Company  `firestore:"company,omitempty"          jsonld:"https://onerecord.iata.org/cargo#branch#company"`
    ContactPerson   Person   `firestore:"contactPerson,omitempty"    jsonld:"https://onerecord.iata.org/cargo#branch#contactPerson"`
    Location        Location `firestore:"location,omitempty"         jsonld:"https://onerecord.iata.org/cargo#branch#location"`
    OtherIdentifier string   `firestore:"otherIdentifier,omitempty"  jsonld:"https://onerecord.iata.org/cargo#branch#otherIdentifier"`
}

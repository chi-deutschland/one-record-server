package model

type RegulatedEntity struct {
    ID                                  string            `firestore:"id"                                       jsonld:"@id"`
    Type                                string            `firestore:"type"                                     jsonld:"@type" default:"https://onerecord.iata.org/cargo#RegulatedEntity"`
    Entity                              Branch            `firestore:"entity,omitempty"                         jsonld:"https://onerecord.iata.org/cargo#regulatedEntity#entity"`
    ExpiryDate                          string            `firestore:"expiryDate,omitempty"                     jsonld:"https://onerecord.iata.org/cargo#regulatedEntity#expiryDate"`
    RegulatedEntityCategory             string            `firestore:"regulatedEntityCategory,omitempty"        jsonld:"https://onerecord.iata.org/cargo#regulatedEntity#regulatedEntityCategory"`
    RegulatedEntityIdentifier           string            `firestore:"regulatedEntityIdentifier,omitempty"      jsonld:"https://onerecord.iata.org/cargo#regulatedEntity#regulatedEntityIdentifier"`
}
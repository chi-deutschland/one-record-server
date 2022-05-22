package model

type ExternalReference struct {
    ID                  string        `firestore:"id"                               jsonld:"@id"`
    Type                string        `firestore:"type"                             jsonld:"@type" default:"https://onerecord.iata.org/cargo#ExternalReference"`
    DocumentOriginator  Company       `firestore:"documentOriginator,omitempty"     jsonld:"https://onerecord.iata.org/cargo#externalReference#documentOriginator"`
    Location            Location      `firestore:"location,omitempty"               jsonld:"https://onerecord.iata.org/cargo#externalReference#location"`
    DocumentChecksum    string        `firestore:"documentChecksum,omitempty"       jsonld:"https://onerecord.iata.org/cargo#externalReference#documentChecksum"`
    DocumentId          string        `firestore:"documentId,omitempty"             jsonld:"https://onerecord.iata.org/cargo#externalReference#documentId"`
    DocumentLink        string        `firestore:"documentLink,omitempty"           jsonld:"https://onerecord.iata.org/cargo#externalReference#documentLink"`
    DocumentName        string        `firestore:"documentName,omitempty"           jsonld:"https://onerecord.iata.org/cargo#externalReference#documentName"`
    DocumentType        string        `firestore:"documentType,omitempty"           jsonld:"https://onerecord.iata.org/cargo#externalReference#documentType"`
    DocumentVersion     string        `firestore:"documentVersion,omitempty"        jsonld:"https://onerecord.iata.org/cargo#externalReference#documentVersion"`
    ExpiryDate          string        `firestore:"expiryDate,omitempty"             jsonld:"https://onerecord.iata.org/cargo#externalReference#expiryDate"`
    ValidFrom           string        `firestore:"validFrom,omitempty"              jsonld:"https://onerecord.iata.org/cargo#externalReference#validFrom"`
}   
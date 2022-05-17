package model
type Company struct {
    ID              string            `firestore:"id"                       jsonld:"@id"`
    Type            string            `firestore:"type"                     jsonld:"@type" default:"https://onerecord.iata.org/cargo#Company"`
    CompanyName     string            `firestore:"companyName,omitempty"    jsonld:"https://onerecord.iata.org/cargo#company#companyName"`
    Branch          CompanyBranch     `firestore:"branch,omitempty"         jsonld:"https://onerecord.iata.org/cargo#company#branch"`
    Role            string            `firestore:"role,omitempty"           jsonld:"https://onerecord.iata.org/cargo#company#role"`
}

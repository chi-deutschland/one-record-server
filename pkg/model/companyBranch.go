package model

type CompanyBranch struct {
    ID         string       `firestore:"id"                     jsonld:"@id"`
    Type       string       `firestore:"type"                   jsonld:"@type" default:"https://onerecord.iata.org/cargo#CompanyBranch"`
    BranchName string       `firestore:"branchName,omitempty"   jsonld:"https://onerecord.iata.org/cargo#companyBranch#branchName"`
    Location   Location     `firestore:"location,omitempty"     jsonld:"https://onerecord.iata.org/cargo#companyBranch#location"`
    Company    string      `firestore:"company,omitempty"     jsonld:"https://onerecord.iata.org/cargo#companyBranch#company"`
}

package model

type Person struct {
    ID                     string                     `firestore:"id"                           jsonld:"@id"`
    Type                   string                     `firestore:"type"                         jsonld:"@type" default:"https://onerecord.iata.org/cargo#Person"`
    AssociatedBranch       CompanyBranch              `firestore:"associatedBranch,omitempty"   jsonld:"https://onerecord.iata.org/cargo#person#associatedBranch"`
    Contact                string                     `firestore:"contact,omitempty"            jsonld:"https://onerecord.iata.org/cargo#person#contact"`
    Documents              []ExternalReference        `firestore:"documents,omitempty"          jsonld:"https://onerecord.iata.org/cargo#person#documents"`
    ContactType            string                     `firestore:"contactType,omitempty"        jsonld:"https://onerecord.iata.org/cargo#person#contactType"`
    Department             string                     `firestore:"department,omitempty"         jsonld:"https://onerecord.iata.org/cargo#person#department"`
    EmployeeId             string                     `firestore:"employeeId,omitempty"         jsonld:"https://onerecord.iata.org/cargo#person#employeeId"`
    FirstName              string                     `firestore:"firstName,omitempty"          jsonld:"https://onerecord.iata.org/cargo#person#firstName"`
    JobTitle               string                     `firestore:"jobTitle,omitempty"           jsonld:"https://onerecord.iata.org/cargo#person#jobTitle"`
    LastName               string                     `firestore:"lastName,omitempty"           jsonld:"https://onerecord.iata.org/cargo#person#lastName"`
    MiddleName             string                     `firestore:"middleName,omitempty"         jsonld:"https://onerecord.iata.org/cargo#person#middleName"`
    Salutation             string                     `firestore:"salutation,omitempty"         jsonld:"https://onerecord.iata.org/cargo#person#salutation"`
}
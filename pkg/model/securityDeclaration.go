package model

type SecurityDeclaration struct {
    ID                                      string                      `firestore:"id"                                      jsonld:"@id"`
    Type                                    string                      `firestore:"type"                                    jsonld:"@type" default:"https://onerecord.iata.org/cargo#SecurityDeclaration"`
    IssuedBy                                Person                      `firestore:"issuedBy,omitempty"                      jsonld:"https://onerecord.iata.org/cargo#securityDeclaration#issuedBy"`
    // OtherRegulatedEntity                    RegulatedEntity             `firestore:"otherRegulatedEntity,omitempty"          jsonld:"https://onerecord.iata.org/cargo#securityDeclaration#otherRegulatedEntity"`
    // Piece                                Piece                       `firestore:"piece,omitempty"                         jsonld:"https://onerecord.iata.org/cargo#securityDeclaration#piece"`
    // ReceivedFrom                            RegulatedEntity             `firestore:"receivedFrom,omitempty"                  jsonld:"https://onerecord.iata.org/cargo#securityDeclaration#receivedFrom"`
    // RegulatedEntityIssuer                   RegulatedEntity             `firestore:"regulatedEntityIssuer,omitempty"         jsonld:"https://onerecord.iata.org/cargo#securityDeclaration#regulatedEntityIssuer"`
    AdditionalSecurityInformation           string                      `firestore:"additionalSecurityInformation,omitempty" jsonld:"https://onerecord.iata.org/cargo#securityDeclaration#additionalSecurityInformation"`
    GroundsForExemption                     string                      `firestore:"groundsForExemption,omitempty"           jsonld:"https://onerecord.iata.org/cargo#securityDeclaration#groundsForExemption"`
    IssuedOn                                string                      `firestore:"issuedOn,omitempty"                      jsonld:"https://onerecord.iata.org/cargo#securityDeclaration#issuedOn"`
    OtherScreeningMethods                   []string                    `firestore:"otherScreeningMethods,omitempty"         jsonld:"https://onerecord.iata.org/cargo#securityDeclaration#otherScreeningMethods"`
    ScreeningMethod                         string                      `firestore:"screeningMethod,omitempty"               jsonld:"https://onerecord.iata.org/cargo#securityDeclaration#screeningMethod"`
    SecurityStatus                          string                      `firestore:"securityStatus,omitempty"                jsonld:"https://onerecord.iata.org/cargo#securityDeclaration#securityStatus"`
}
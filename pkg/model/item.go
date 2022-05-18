package model

type Item struct {
    ID                          string               `firestore:"id"                                           jsonld:"@id"`
    Type                        string               `firestore:"type"                                         jsonld:"@type" default:"https://onerecord.iata.org/cargo#Item"`
    IsInPiece                   Piece                `firestore:"isInPiece,omitempty"                          jsonld:"https://onerecord.iata.org/cargo#item#isInPiece"`
    OtherIdentifier             string               `firestore:"otherIdentifier,omitempty"                    jsonld:"https://onerecord.iata.org/cargo#item#otherIdentifier"`
    Product                     Product              `firestore:"product,omitempty"                            jsonld:"https://onerecord.iata.org/cargo#item#product"`
    ProductionCountry           string               `firestore:"productionCountry,omitempty"                  jsonld:"https://onerecord.iata.org/cargo#item#productionCountry"`
    Quantity                    int                  `firestore:"quantity,omitempty"                           jsonld:"https://onerecord.iata.org/cargo#item#quantity"`
    TargetCountry               string               `firestore:"targetCountry,omitempty"                      jsonld:"https://onerecord.iata.org/cargo#item#targetCountry"`
    UnitPrice                   string               `firestore:"unitPrice,omitempty"                          jsonld:"https://onerecord.iata.org/cargo#item#unitPrice"`
    Weight                      string               `firestore:"weight,omitempty"                             jsonld:"https://onerecord.iata.org/cargo#item#weight"`
    BatchNumber                 string               `firestore:"batchNumber,omitempty"                        jsonld:"https://onerecord.iata.org/cargo#item#batchNumber"`
    LotNumber                   string               `firestore:"lotNumber,omitempty"                          jsonld:"https://onerecord.iata.org/cargo#item#lotNumber"`
    ProductExpiryDate           string               `firestore:"productExpiryDate,omitempty"                  jsonld:"https://onerecord.iata.org/cargo#item#productExpiryDate"`
    ProductionDate              string               `firestore:"productionDate,omitempty"                     jsonld:"https://onerecord.iata.org/cargo#item#productionDate"`
    QuantityForUnitPrice        string               `firestore:"quantityForUnitPrice,omitempty"               jsonld:"https://onerecord.iata.org/cargo#item#quantityForUnitPrice"`
}
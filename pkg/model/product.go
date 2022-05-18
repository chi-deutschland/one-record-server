package model

type Product struct {
    ID                          string               `firestore:"id"                                           jsonld:"@id"`
    Type                        string               `firestore:"type"                                         jsonld:"@type" default:"https://onerecord.iata.org/cargo#Product"`
    // isInItems                   Item                 `firestore:"isInItems,omitempty"                          jsonld:"https://onerecord.iata.org/cargo#product#isInItems"`
    // isInPieces                  Piece                `firestore:"isInPieces,omitempty"                         jsonld:"https://onerecord.iata.org/cargo#product#isInPieces"`
    Manufacturer                Company              `firestore:"manufacturer,omitempty"                       jsonld:"https://onerecord.iata.org/cargo#product#manufacturer"`
    OtherIdentifier             string               `firestore:"otherIdentifier,omitempty"                    jsonld:"https://onerecord.iata.org/cargo#product#otherIdentifier"`
    CommodityItemNumber         string               `firestore:"commodityItemNumber,omitempty"                jsonld:"https://onerecord.iata.org/cargo#product#commodityItemNumber"`
    HsCode                      string               `firestore:"hsCode,omitempty"                             jsonld:"https://onerecord.iata.org/cargo#product#hsCode"`
    HsCommodityDescription      string               `firestore:"hsCommodityDescription,omitempty"             jsonld:"https://onerecord.iata.org/cargo#product#hsCommodityDescription"`
    HsCommodityName             string               `firestore:"hsCommodityName,omitempty"                    jsonld:"https://onerecord.iata.org/cargo#product#hsCommodityName"`
    HsType                      string               `firestore:"hsType,omitempty"                             jsonld:"https://onerecord.iata.org/cargo#product#hsType"`
    ProductDescription          string               `firestore:"productDescription,omitempty"                 jsonld:"https://onerecord.iata.org/cargo#product#productDescription"`
    ProductIdentifier           string               `firestore:"productIdentifier,omitempty"                  jsonld:"https://onerecord.iata.org/cargo#product#productIdentifier"`
}
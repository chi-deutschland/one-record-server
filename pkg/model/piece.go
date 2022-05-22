package model

type Piece struct {
    ID                                string                             `firestore:"id"                                 jsonld:"@id"`
    Type                              string                             `firestore:"type"                               jsonld:"@type" default:"https://onerecord.iata.org/cargo#Piece"`
    CompanyIdentifier                 string                             `firestore:"companyIdentifier,omitempty"        jsonld:"https://onerecord.iata.org/cargo#piece#companyIdentifier"`
    // Events                         []Event                            `firestore:"events,omitempty"                   jsonld:"https://onerecord.iata.org/cargo#piece#events"`
    // IotDevices                     []IotDevice
    // ContainedItems                    []Item
    ContainedPieces                   []Piece                            `firestore:"containedPieces,omitempty"          jsonld:"https://onerecord.iata.org/cargo#piece#containedPieces"`
    Product                           Product                            `firestore:"product,omitempty"                  jsonld:"https://onerecord.iata.org/cargo#piece#product"`
    // CustomsInfo                    CustomsInfo
    // Dimensions                     Dimensions
    // ExternalReferences             ExternalReference                  `firestore:"externalRefernces,omitempty"        jsonld:"https://onerecord.iata.org/cargo#piece#externalRefernces"`
    GrossWeight                       string                             `firestore:"grossWeight,omitempty"              jsonld:"https://onerecord.iata.org/cargo#piece#grossWeight"`
    // HandlingInstructions           HandlingInstructions
    // OtherIdentifiers               []OtherIdentifier
    OtherParty                        Company                            `firestore:"otherParty,omitempty"               jsonld:"https://onerecord.iata.org/cargo#piece#otherParty"`
    // PackagingType                  PackagingType
    // Parties                        []Party
    // ProductionCountry              ProductionCountry
    SecurityDeclaration               SecurityDeclaration                `firestore:"securityDeclaration,omitempty"      jsonld:"https://onerecord.iata.org/cargo#piece#securityDeclaration"`
    // SecurityStatus                 SecurityDeclaration                `firestore:"securityStatus,omitempty"           jsonld:"https://onerecord.iata.org/cargo#piece#securityStatus"`
    // ServiceRequest                 ServiceRequest
    // Shipment                          Shipment                           `firestore:"shipment,omitempty"                 jsonld:"https://onerecord.iata.org/cargo#piece#shipment"`
    Shipper                           Company                            `firestore:"shipper,omitempty"                  jsonld:"https://onerecord.iata.org/cargo#piece#shipper"`
    // SpecialHandling                SpecialHandling
    // TransportMovements             []TransportMovement                `firestore:"transportMovements,omitempty"       jsonld:"https://onerecord.iata.org/cargo#piece#transportMovements"`
    // TransportSegments              []TransportSegment
    // UldReference                   ULD
    // VolumetricWeight               VolumetricWeight
    Coload                            bool                               `firestore:"coload,omitempty"                   jsonld:"https://onerecord.iata.org/cargo#piece#coload"`
    DeclaredValueForCarriage          string                             `firestore:"declaredValueForCarriage,omitempty" jsonld:"https://onerecord.iata.org/cargo#piece#declaredValueForCarriage"`
    DeclaredValueForCustoms           string                             `firestore:"declaredValueForCustoms,omitempty"  jsonld:"https://onerecord.iata.org/cargo#piece#declaredValueForCustoms"`
    GoodsDescription                  string                             `firestore:"goodsDescription,omitempty"         jsonld:"https://onerecord.iata.org/cargo#piece#goodsDescription"`
    //                                LoadType                    
    NvdForCarriage                    bool                               `firestore:"nvdForCarriage,omitempty"           jsonld:"https://onerecord.iata.org/cargo#piece#nvdForCarriage"`
    NvdForCustoms                     bool                               `firestore:"nvdForCustoms,omitempty"            jsonld:"https://onerecord.iata.org/cargo#piece#nvdForCustoms"`
    //                                PackageMarkCoded            
    PackagedeIdentifier               string                             `firestore:"packagedeIdentifier,omitempty"      jsonld:"https://onerecord.iata.org/cargo#piece#packagedeIdentifier"`
    ShippingMarks                     string                             `firestore:"shippingMarks,omitempty"            jsonld:"https://onerecord.iata.org/cargo#piece#shippingMarks"`
    Slac                              int                                `firestore:"slac,omitempty"                     jsonld:"https://onerecord.iata.org/cargo#piece#slac"`
    Stackable                         bool                               `firestore:"stackable,omitempty"                jsonld:"https://onerecord.iata.org/cargo#piece#stackable"`
    Turnable                          bool                               `firestore:"turnable,omitempty"                 jsonld:"https://onerecord.iata.org/cargo#piece#turnable"`
    Upid                              string                             `firestore:"upid,omitempty"                     jsonld:"https://onerecord.iata.org/cargo#piece#upid"`
}

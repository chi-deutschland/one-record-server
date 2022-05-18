package model

type Shipment struct {
    ID                            string                         `firestore:"id"                               jsonld:"@id"`
    Type                          string                         `firestore:"type"                             jsonld:"@type" default:"https://onerecord.iata.org/cargo#Shipment"`
    ContainedPieces               []Piece                        `firestore:"containedPieces,omitempty"        jsonld:"https://onerecord.iata.org/cargo#shipment#containedPieces"`
    DeliveryLocation              Location                       `firestore:"deliveryLocation,omitempty"       jsonld:"https://onerecord.iata.org/cargo#shipment#deliveryLocation"`
    // Dimensions                 Dimensions
    ExternalReferences            []ExternalReference            `firestore:"externalReferences,omitempty"     jsonld:"https://onerecord.iata.org/cargo#shipment#externalReferences"`
    FreightForwarder              Company                        `firestore:"freightForwarder,omitempty"       jsonld:"https://onerecord.iata.org/cargo#shipment#freightForwarder"`
    // Insurance                  Insurance
    // Parties                    Party
    Shipper                       Company                        `firestore:"shipper,omitempty"                jsonld:"https://onerecord.iata.org/cargo#shipment#shipper"`
    // TotalGrossWeight           Value
    // VolumetricWeight           VolumetricWeight
    // WaybillNumber              Waybill
    DeliveryDate                  string                         `firestore:"deliveryDate,omitempty"           jsonld:"https://onerecord.iata.org/cargo#shipment#deliveryDate"`
    GoodsDescription              string                         `firestore:"goodsDescription,omitempty"       jsonld:"https://onerecord.iata.org/cargo#shipment#goodsDescription"`
    Incoterms                     string                         `firestore:"incoterms,omitempty"              jsonld:"https://onerecord.iata.org/cargo#shipment#incoterms"`
    //OtherChargesIndicator    
    TotalPieceCount               int                            `firestore:"totalPieceCount,omitempty"        jsonld:"https://onerecord.iata.org/cargo#shipment#totalPieceCount"`
    TotalSLAC                     int                            `firestore:"totalSLAC,omitempty"              jsonld:"https://onerecord.iata.org/cargo#shipment#totalSLAC"`
    // WeightValuationIndicator
}

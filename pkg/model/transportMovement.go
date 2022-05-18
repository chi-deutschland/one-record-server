package model

type TransportMovement struct {
    ID                                string                             `firestore:"id"                                      jsonld:"@id"`
    Type                              string                             `firestore:"type"                                    jsonld:"@type" default:"https://onerecord.iata.org/cargo#TransportMovement"`
    ArrivalLocation                   Location                           `firestore:"arrivalLocation,omitempty"               jsonld:"https://onerecord.iata.org/cargo#transportMovement#arrivalLocation"`
    // Co2CalculationMethod           CO2CalcMethod                 
    // Co2Emissions                   CO2Emissions                 
    DepartureLocation                 Location                           `firestore:"departureLocation,omitempty"             jsonld:"https://onerecord.iata.org/cargo#transportMovement#departureLocation"`
    // DistanceCalculated             Value                      
    // DistanceMeasured               Value                        
    ExternalReferences                []ExternalReference                `firestore:"externalReferences,omitempty"            jsonld:"https://onerecord.iata.org/cargo#transportMovement#externalReferences"`
    // FuelAmountCalculated           Value                        
    // FuelAmountMeasured             Value                      
    // MovementTimes                  MovementTimes               
    // Payload                        Value                         
    // TransportMeans                 TransportMeans             
    // TransportMeansOperators        TransportMeansOperators
    TransportedPieces                 []Piece                            `firestore:"transportedPieces,omitempty"             jsonld:"https://onerecord.iata.org/cargo#transportMovement#transportedPieces"`
    // TransportedUlds                ULD                           
    FuelType                          string                             `firestore:"fuelType,omitempty"                      jsonld:"https://onerecord.iata.org/cargo#transportMovement#fuelType"`
    // ModeCode                                                  
    // ModeQualifier                                             
    Seal                              string                             `firestore:"seal,omitempty"                          jsonld:"https://onerecord.iata.org/cargo#transportMovement#seal"`
    TransportIdentifier               string                             `firestore:"transportIdentifier,omitempty"           jsonld:"https://onerecord.iata.org/cargo#transportMovement#transportIdentifier"`
    UnplannedStop                     string                             `firestore:"unplannedStop,omitempty"                 jsonld:"https://onerecord.iata.org/cargo#transportMovement#unplannedStop"`
}
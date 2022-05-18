package model

type Event struct {
    ID                     string         `firestore:"id"                                               jsonld:"@id"`
    Type                   string         `firestore:"type"                                             jsonld:"@type" default:"https://onerecord.iata.org/cargo#Event"`
    LinkedPiece            Piece          `firestore:"linkedPiece,omitempty"                            jsonld:"https://onerecord.iata.org/cargo#event#linkedPiece"`
    Location               Location       `firestore:"location,omitempty"                               jsonld:"https://onerecord.iata.org/cargo#event#location"`
    PerformedBy            Company        `firestore:"performedBy,omitempty"                            jsonld:"https://onerecord.iata.org/cargo#event#performedBy"`
    PerformedByPerson      Person         `firestore:"performedByPerson,omitempty"                      jsonld:"https://onerecord.iata.org/cargo#event#performedByPerson"`
    DateTime               string         `firestore:"dateTime,omitempty"                               jsonld:"https://onerecord.iata.org/cargo#event#dateTime"`
    EventCode              string         `firestore:"eventCode,omitempty"                              jsonld:"https://onerecord.iata.org/cargo#event#eventCode"`
    EventName              string         `firestore:"eventName,omitempty"                              jsonld:"https://onerecord.iata.org/cargo#event#eventName"`
    EventTypeIndicator     string         `firestore:"eventTypeIndicator,omitempty"                     jsonld:"https://onerecord.iata.org/cargo#event#eventTypeIndicator"`
}                       

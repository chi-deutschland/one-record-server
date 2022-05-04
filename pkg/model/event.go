package model

type Event struct {
	ID         			string 		`firestore:"id" json:"id,omitempty"`
	LinkedPiece			Piece		`firestore:"linkedPiece,omitempty" json:"linkedPiece,omitempty"`
	Location			Location	`firestore:"location,omitempty" json:"location,omitempty"`
	PerformedBy			Company		`firestore:"performedBy,omitempty" json:"performedBy,omitempty"`
	PerformedByPerson	Person		`firestore:"performedByPerson,omitempty" json:"performedByPerson,omitempty"`
	DateTime			string		`firestore:"dateTime,omitempty" json:"dateTime,omitempty"`
	EventCode			string		`firestore:"eventCode,omitempty" json:"eventCode,omitempty"`
	EventName			string		`firestore:"eventName,omitempty" json:"eventName,omitempty"`
	EventTypeIndicator	string		`firestore:"eventTypeIndicator,omitempty" json:"eventTypeIndicator,omitempty"`
}

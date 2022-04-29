package model

import (
	"time"
)

type Event struct {
	ID         string `firestore:"id" json:"id,omitempty"`
	LinkedObject		LogisticsObject	`firestore:"linkedObject,omitempty" json:"linkedObject,omitempty"`
	// Location			Location
	PerformedBy			Company			`firestore:"performedBy,omitempty" json:"performedBy,omitempty"`
	// PerformedByPerson	Person
	DateTime			time.Time		`firestore:"dateTime,omitempty" json:"dateTime,omitempty"`
	EventCode			string			`firestore:"eventCode,omitempty" json:"eventCode,omitempty"`
	EventName			string			`firestore:"eventName,omitempty" json:"eventName,omitempty"`
	// EventTypeIndicator
}

package model

type LogisticsObject struct {
	ID     string        `firestore:"id" json:"id,omitempty"`
	CompanyIdentifier	string		`firestore:"companyIdentifier,omitempty" json:"companyIdentifier,omitempty"`
	Events				[]Event		`firestore:"events,omitempty" json:"events,omitempty"`
	// IotDevices			[]IotDevice
}

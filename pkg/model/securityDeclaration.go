package model

type SecurityDeclaration struct {
	ID									string					`firestore:"id" json:"id,omitempty"`
	IssuedBy							Person					`firestore:"issuedBy,omitempty" json:"issuedBy,omitempty"`
	OtherRegulatedEntity				RegulatedEntity			`firestore:"otherRegulatedEntity,omitempty" json:"otherRegulatedEntity,omitempty"`
	// Piece								Piece					`firestore:"piece,omitempty" json:"piece,omitempty"`
	ReceivedFrom						RegulatedEntity			`firestore:"receivedFrom,omitempty" json:"receivedFrom,omitempty"`
	RegulatedEntityIssuer				RegulatedEntity			`firestore:"regulatedEntityIssuer,omitempty" json:"regulatedEntityIssuer,omitempty"`
	AdditionalSecurityInformation		string					`firestore:"additionalSecurityInformation,omitempty" json:"additionalSecurityInformation,omitempty"`
	GroundsForExemption					string					`firestore:"groundsForExemption,omitempty" json:"groundsForExemption,omitempty"`
	IssuedOn							string					`firestore:"issuedOn,omitempty" json:"issuedOn,omitempty"`
	OtherScreeningMethods				[]string					`firestore:"otherScreeningMethods,omitempty" json:"otherScreeningMethods,omitempty"`
	ScreeningMethod						string					`firestore:"screeningMethod,omitempty" json:"screeningMethod,omitempty"`
	SecurityStatus						string					`firestore:"securityStatus,omitempty" json:"securityStatus,omitempty"`
}
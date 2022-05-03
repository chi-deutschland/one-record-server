package model

type ExternalReference struct {
	ID         			string 		`firestore:"id" json:"id,omitempty"`
	DocumentOriginator	Company		`firestore:"documentOriginator,omitempty" json:"documentOriginator,omitempty"`
	Location			Location	`firestore:"location,omitempty" json:"location,omitempty"`
	DocumentChecksum	string		`firestore:"documentChecksum,omitempty" json:"documentChecksum,omitempty"`
	DocumentId			string		`firestore:"documentId,omitempty" json:"documentId,omitempty"`
	DocumentLink		string		`firestore:"documentLink,omitempty" json:"documentLink,omitempty"`
	DocumentName		string		`firestore:"documentName,omitempty" json:"documentName,omitempty"`
	DocumentType		string		`firestore:"documentType,omitempty" json:"documentType,omitempty"`
	DocumentVersion		string		`firestore:"documentVersion,omitempty" json:"documentVersion,omitempty"`
	ExpiryDate			string		`firestore:"expiryDate,omitempty" json:"expiryDate,omitempty"`
	ValidFrom			string		`firestore:"validFrom,omitempty" json:"validFrom,omitempty"`
}
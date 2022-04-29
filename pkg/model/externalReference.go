package model

import (
	"time"
)

type ExternalReference struct {
	ID         string `firestore:"id" json:"id,omitempty"`
	DocumentOriginator	Company		`firestore:"documentOriginator,omitempty" json:"documentOriginator,omitempty"`
	Location			Location	`firestore:"location,omitempty" json:"location,omitempty"`
	DocumentChecksum	string		`firestore:"documentChecksum,omitempty" json:"documentChecksum,omitempty"`
	DocumentId			string		`firestore:"documentId,omitempty" json:"documentId,omitempty"`
	DocumentLink		string		`firestore:"documentLink,omitempty" json:"documentLink,omitempty"`
	DocumentName		string		`firestore:"documentName,omitempty" json:"documentName,omitempty"`
	DocumentType		string		`firestore:"documentType,omitempty" json:"documentType,omitempty"`
	DocumentVersion		string		`firestore:"documentVersion,omitempty" json:"documentVersion,omitempty"`
	ExpiryDate			time.Time	`firestore:"expiryDate,omitempty" json:"expiryDate,omitempty"`
	ValidFrom			time.Time	`firestore:"validFrom,omitempty" json:"validFrom,omitempty"`
}
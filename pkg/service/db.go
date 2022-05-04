package service

import "github.com/chi-deutschland/one-record-server/pkg/model"

type DBService interface {
	GetCompanies(projectID, role string) ([]model.Company, error)

	GetCompany(projectID, companyID string) (model.Company, error)
	AddCompany(projectID string, company model.Company) (companyID string, err error)
	UpdateCompany(projectID, companyID string, company model.Company) (error)
	DeleteCompany(projectID, companyID string) (error)
	DeleteCompanyFields(projectID, companyID string, fields []string) (error)

	GetPieces(projectID, companyID string) ([]model.Piece, error)

	GetPiece(projectID, companyID, pieceID string) (model.Piece, error)
	AddPiece(projectID, companyID string, piece model.Piece) (pieceID string, err error)
	UpdatePiece(projectID, companyID, pieceID string, piece model.Piece) (error)
	DeletePiece(projectID, companyID, pieceID string) (error)
	DeletePieceFields(projectID, companyID, pieceID string, fields []string) (error)

	GetEvents(projectID, companyID, pieceID string) ([]model.Event, error)

	GetEvent(projectID, companyID, pieceID, eventID string) (model.Event, error)
	AddEvent(projectID, companyID, pieceID string, event model.Event) (eventID string, err error)
	UpdateEvent(projectID, companyID, pieceID, eventID string, event model.Event) (error)
	DeleteEvent(projectID, companyID, pieceID, eventID string) (error)
	DeleteEventFields(projectID, companyID, pieceID, eventID string, fields []string) (error)

	GetExternalReferences(projectID, companyID, pieceID string) ([]model.ExternalReference, error)

	GetExternalReference(projectID, companyID, pieceID, externalReferenceID string) (model.ExternalReference, error)
	AddExternalReference(projectID, companyID, pieceID string, externalReference model.ExternalReference) (externalReferenceID string, err error)
	UpdateExternalReference(projectID, companyID, pieceID, externalReferenceID string, externalReference model.ExternalReference) (error)
	DeleteExternalReference(projectID, companyID, pieceID, externalReferenceID string) (error)
	DeleteExternalReferenceFields(projectID, companyID, pieceID, externalReferenceID string, fields []string) (error)

	GetSecurityDeclaration(projectID, companyID, pieceID string) (model.SecurityDeclaration, error)
	AddSecurityDeclaration(projectID, companyID, pieceID string, piece model.SecurityDeclaration) (error)
	UpdateSecurityDeclaration(projectID, companyID, pieceID string, piece model.SecurityDeclaration) (error)
	DeleteSecurityDeclaration(projectID, companyID, pieceID string) (error)
	DeleteSecurityDeclarationFields(projectID, companyID, pieceID string, fields []string) (error)
}

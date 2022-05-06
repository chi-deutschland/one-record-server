package service

import (
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/utils/conv"
)

type DBService interface {
	GetCompanies(projectID, role, colPath string) ([]model.Company, error)

	GetCompany(projectID, role, docPath string) (model.Company, error)
	AddCompany(projectID, role, colPath, id string, company model.Company) (ID string, err error)
	UpdateCompany( projectID, role, docPath string, updates utils.FirestoreMap) (error)
	DeleteCompany( projectID, role, docPath string) (error)
	DeleteCompanyFields( projectID, role, docPath string, fields []string) (error)

	GetPieces(projectID, role, colPath string) ([]model.Piece, error)

	GetPiece(projectID, role, docPath string) (model.Piece, error)
	AddPiece(projectID, role, colPath, id string, piece model.Piece) (ID string, err error)
	UpdatePiece( projectID, role, docPath string, updates utils.FirestoreMap) (error)
	DeletePiece( projectID, role, docPath string) (error)
	DeletePieceFields( projectID, role, docPath string, fields []string) (error)

	GetEvents(projectID, role, colPath string) ([]model.Event, error)

	GetEvent(projectID, role, docPath string) (model.Event, error)
	AddEvent(projectID, role, colPath, id string, event model.Event) (ID string, err error)
	UpdateEvent( projectID, role, docPath string, updates utils.FirestoreMap) (error)
	DeleteEvent( projectID, role, docPath string) (error)
	DeleteEventFields( projectID, role, docPath string, fields []string) (error)

	GetExternalReferences(projectID, role, colPath string) ([]model.ExternalReference, error)

	GetExternalReference(projectID, role, docPath string) (model.ExternalReference, error)
	AddExternalReference(projectID, role, colPath, id string, externalReference model.ExternalReference) (ID string, err error)
	UpdateExternalReference( projectID, role, docPath string, updates utils.FirestoreMap) (error)
	DeleteExternalReference( projectID, role, docPath string) (error)
	DeleteExternalReferenceFields( projectID, role, docPath string, fields []string) (error)

	GetSecurityDeclarations(projectID, role, colPath string) ([]model.SecurityDeclaration, error)

	GetSecurityDeclaration(projectID, role, docPath string) (model.SecurityDeclaration, error)
	AddSecurityDeclaration(projectID, role, colPath, id string, securityDeclaration model.SecurityDeclaration) (ID string, err error)
	UpdateSecurityDeclaration( projectID, role, docPath string, updates utils.FirestoreMap) (error)
	DeleteSecurityDeclaration( projectID, role, docPath string) (error)
	DeleteSecurityDeclarationFields( projectID, role, docPath string, fields []string) (error)

	GetShipments(projectID, role, colPath string) ([]model.Shipment, error)

	GetShipment(projectID, role, docPath string) (model.Shipment, error)
	AddShipment(projectID, role, colPath, id string, shipment model.Shipment) (ID string, err error)
	UpdateShipment( projectID, role, docPath string, updates utils.FirestoreMap) (error)
	DeleteShipment( projectID, role, docPath string) (error)
	DeleteShipmentFields( projectID, role, docPath string, fields []string) (error)

	GetTransportMovements(projectID, role, colPath string) ([]model.TransportMovement, error)

	GetTransportMovement(projectID, role, docPath string) (model.TransportMovement, error)
	AddTransportMovement(projectID, role, colPath, id string, transportMovement model.TransportMovement) (ID string, err error)
	UpdateTransportMovement( projectID, role, docPath string, updates utils.FirestoreMap) (error)
	DeleteTransportMovement( projectID, role, docPath string) (error)
	DeleteTransportMovementFields( projectID, role, docPath string, fields []string) (error)
}

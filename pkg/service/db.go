package service

import "github.com/chi-deutschland/one-record-server/pkg/model"

type DBService interface {
	GetCompanies(projectID, role string) ([]model.Company, error)
	GetCompany(projectID, companyID string) (model.Company, error)
	AddCompany(projectID string, company model.Company) (companyID string, err error)
	UpdateCompany(projectID, companyID string, company model.Company) (error)
	DeleteCompany(projectID, companyID string) (error)
	GetPieces(projectID, companyID string) ([]model.Piece, error)
	GetPiece(projectID, companyID string, pieceID string) (model.Piece, error)
	// GetLogisticsObject(projectID, companyID string, logisticsObjectID string) (model.LogisticsObject, error)
	// GetEvents(projectID, companyID string, pieceID string) ([]model.Event, error)
	// GetShipment(projectID, companyID string, pieceID string) (model.Shipment, error)
}

package service

import "github.com/chi-deutschland/one-record-server/pkg/model"

type DBService interface {
	GetCompanies(projectID, role string) ([]model.Company, error)
	GetCompany(projectID, companyID string) (model.Company, error)
}

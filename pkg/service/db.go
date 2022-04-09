package service

import "github.com/chi-deutschland/one-record-server/pkg/model"

type DBService interface {
	GetCompanies(projectID string) ([]model.Company, error)
}

package gcp

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	"google.golang.org/api/iterator"
)

type FirestoreService struct {
}

func (f *FirestoreService) GetCompanies(projectID, role string) ([]model.Company, error) {
	var companies []model.Company
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return companies, err
	}
	defer client.Close()
	iter := client.Collection("companies").
		Where("role", "==", role).
		Documents(ctx)
	for {
		var company model.Company

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return companies, err
		}

		if err := doc.DataTo(&company); err != nil {
			return companies, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}

func NewFirestoreService() (*FirestoreService, error) {
	var f FirestoreService
	return &f, nil
}

var _ service.DBService = (*FirestoreService)(nil)

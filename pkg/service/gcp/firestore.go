package gcp

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	"github.com/chi-deutschland/one-record-server/pkg/utils"
	"google.golang.org/api/iterator"
)

type FirestoreService struct {
}

func (f *FirestoreService) GetCompanies(
	projectID,
	role string,
) (
	[]model.Company,
	error,
) {

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

func (f *FirestoreService) GetCompany(
	projectID,
	companyID string,
) (
	model.Company,
	error,
) {

	var company model.Company
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return company, err
	}
	defer client.Close()
	doc, err := client.Collection("companies").
		Doc(companyID).
		Get(ctx)
	if err != nil {
		return company, err
	}

	if err := doc.DataTo(&company); err != nil {
		return company, err
	}

	return company, nil
}

func (f *FirestoreService) AddCompany(
	projectID string,
	company model.Company,
) (	
	companyID string,
	err error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return companyID, err
	}
	defer client.Close()
	
	if company.ID == "" {
		var ref, _, err = client.Collection("companies").
			Add(ctx, company)
		company.ID = ref.ID

		client.Collection("companies").
		Doc(company.ID).
		Set(ctx, map[string]interface{}{
        "id": company.ID}, firestore.MergeAll)

		if err != nil {
			return companyID, err
		}
	} else {
		_, err = client.Collection("companies").
			Doc(company.ID).
			Set(ctx, company)
		if err != nil {
			return companyID, err
		}
	}

	return company.ID, nil
}

func (f *FirestoreService) UpdateCompany(
	projectID,
	companyID string,
	company model.Company,
) (
	error,
) {
	var mapCompany = utils.ToFirestoreMap(company)
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()
	_, err = client.Collection("companies").
		Doc(companyID).
		Set(ctx, mapCompany, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteCompany(
	projectID,
	companyID string,
) (
	error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()
	_, err = client.Collection("companies").
		Doc(companyID).
		Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) GetPieces(
	projectID,
	companyID string,
) (
	[]model.Piece,
	error,
) {
	var pieces []model.Piece
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return pieces, err
	}
	defer client.Close()
	iter := client.
		Collection("companies").Doc(companyID).
		Collection("pieces").
		Documents(ctx)
	for {
		var piece model.Piece

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return pieces, err
		}

		if err := doc.DataTo(&piece); err != nil {
			return pieces, err
		}

		pieces = append(pieces, piece)
	}

	return pieces, nil
}

func (f *FirestoreService) GetPiece(
	projectID,
	companyID string,
	pieceID string,
) (
	model.Piece,
	error,
) {
	var piece model.Piece
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return piece, err
	}
	defer client.Close()
	doc, err := client.
		Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Get(ctx)
	if err != nil {
		return piece, err
	}

	if err := doc.DataTo(&piece); err != nil {
		return piece, err
	}

	return piece, nil
}

// func (f *FirestoreService) GetShipment(
// 	projectID,
// 	companyID string,
// 	pieceID string,
// ) (
// 	model.Shipment,
// 	error,
// ) {
// 	var piece model.Piece
// 	var shipment model.Shipment
// 	ctx := context.Background()
// 	client, err := firestore.NewClient(ctx, projectID)
// 	if err != nil {
// 		return shipment, err
// 	}
// 	defer client.Close()
// 	doc, err := client.
// 		Collection("companies").Doc(companyID).
// 		Collection("pieces").Doc(pieceID).
// 		Get(ctx)
// 	if err != nil {
// 		return shipment, err
// 	}

// 	if err := doc.DataTo(&piece); err != nil {
// 		return shipment, err
// 	}

// 	shipment := piece.

// 	return shipment, nil
// }

func NewFirestoreService() (*FirestoreService, error) {
	var f FirestoreService
	return &f, nil
}

var _ service.DBService = (*FirestoreService)(nil)

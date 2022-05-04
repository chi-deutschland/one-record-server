package gcp

import (
	"context"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	"github.com/chi-deutschland/one-record-server/pkg/utils"
	"google.golang.org/api/iterator"
	"golang.org/x/exp/slices"
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
	ref := client.Collection("companies").
		Doc(companyID)
	err = deleteDocumentAndSubcollections(ctx, client, ref)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteCompanyFields(
	projectID,
	companyID string,
	fields []string,
) (
	error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	updates := []firestore.Update{}

	for _, field := range fields {
		if !slices.Contains(strings.Split(field, "."), "id") {
			updates = append(updates, firestore.Update{Path: field, Value: firestore.Delete})
		}
	}

	_, err = client.Collection("companies").
		Doc(companyID).
		Update(ctx, updates)
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

func (f *FirestoreService) AddPiece(
	projectID,
	companyID string,
	piece model.Piece,
) (	
	pieceID string,
	err error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return pieceID, err
	}
	defer client.Close()
	
	if piece.ID == "" {
		var ref, _, err = client.Collection("companies").Doc(companyID).
			Collection("pieces").
			Add(ctx, piece)
		piece.ID = ref.ID

		client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(piece.ID).
		Set(ctx, map[string]interface{}{
        "id": piece.ID}, firestore.MergeAll)

		if err != nil {
			return pieceID, err
		}
	} else {
		_, err = client.Collection("companies").Doc(companyID).
			Collection("pieces").Doc(piece.ID).
			Set(ctx, piece)
		if err != nil {
			return pieceID, err
		}
	}

	return piece.ID, nil
}

func (f *FirestoreService) UpdatePiece(
	projectID,
	companyID,
	pieceID string,
	piece model.Piece,
) (
	error,
) {
	var mapPiece = utils.ToFirestoreMap(piece)
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()
	_, err = client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Set(ctx, mapPiece, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeletePiece(
	projectID,
	companyID,
	pieceID string,
) (
	error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()
	ref := client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID)
	err = deleteDocumentAndSubcollections(ctx, client, ref)
	if err != nil {
		return err
	}

	return nil
}



func (f *FirestoreService) DeletePieceFields(
	projectID,
	companyID,
	pieceID string,
	fields []string,
) (
	error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	updates := []firestore.Update{}

	for _, field := range fields {
		if !slices.Contains(strings.Split(field, "."), "id") {
			updates = append(updates, firestore.Update{Path: field, Value: firestore.Delete})
		}
	}

	_, err = client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Update(ctx, updates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) GetEvents(
	projectID,
	companyID,
	pieceID string,
) (
	[]model.Event,
	error,
) {
	var events []model.Event
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return events, err
	}
	defer client.Close()
	iter := client.
		Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("events").
		Documents(ctx)
	for {
		var event model.Event

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return events, err
		}

		if err := doc.DataTo(&event); err != nil {
			return events, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (f *FirestoreService) GetEvent(
	projectID,
	companyID,
	pieceID,
	eventID string,
) (
	model.Event,
	error,
) {
	var event model.Event
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return event, err
	}
	defer client.Close()
	doc, err := client.
		Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("events").Doc(eventID).
		Get(ctx)
	if err != nil {
		return event, err
	}

	if err := doc.DataTo(&event); err != nil {
		return event, err
	}

	return event, nil
}

func (f *FirestoreService) AddEvent(
	projectID,
	companyID,
	pieceID string,
	event model.Event,
) (	
	eventID string,
	err error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return eventID, err
	}
	defer client.Close()
	
	if event.ID == "" {
		var ref, _, err = client.Collection("companies").Doc(companyID).
			Collection("pieces").Doc(pieceID).
			Collection("events").
			Add(ctx, event)
		event.ID = ref.ID

		client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("events").Doc(event.ID).
		Set(ctx, map[string]interface{}{
        "id": event.ID}, firestore.MergeAll)

		if err != nil {
			return eventID, err
		}
	} else {
		_, err = client.Collection("companies").Doc(companyID).
			Collection("pieces").Doc(pieceID).
			Collection("events").Doc(event.ID).
			Set(ctx, event)
		if err != nil {
			return eventID, err
		}
	}

	return event.ID, nil
}

func (f *FirestoreService) UpdateEvent(
	projectID,
	companyID,
	pieceID,
	eventID string,
	event model.Event,
) (
	error,
) {
	var mapEvent = utils.ToFirestoreMap(event)
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, eventID)
	if err != nil {
		return err
	}
	defer client.Close()
	_, err = client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("events").Doc(eventID).
		Set(ctx, mapEvent, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteEvent(
	projectID,
	companyID,
	pieceID,
	eventID string,
) (
	error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, eventID)
	if err != nil {
		return err
	}
	defer client.Close()
	ref := client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("events").Doc(eventID)
	err = deleteDocumentAndSubcollections(ctx, client, ref)
	if err != nil {
		return err
	}

	return nil
}



func (f *FirestoreService) DeleteEventFields(
	projectID,
	companyID,
	pieceID,
	eventID string,
	fields []string,
) (
	error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	updates := []firestore.Update{}

	for _, field := range fields {
		if !slices.Contains(strings.Split(field, "."), "id") {
			updates = append(updates, firestore.Update{Path: field, Value: firestore.Delete})
		}
	}

	_, err = client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("events").Doc(eventID).
		Update(ctx, updates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) GetExternalReferences(
	projectID,
	companyID,
	pieceID string,
) (
	[]model.ExternalReference,
	error,
) {
	var externalReferences []model.ExternalReference
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return externalReferences, err
	}
	defer client.Close()
	iter := client.
		Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("externalReferences").
		Documents(ctx)
	for {
		var externalReference model.ExternalReference

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return externalReferences, err
		}

		if err := doc.DataTo(&externalReference); err != nil {
			return externalReferences, err
		}

		externalReferences = append(externalReferences, externalReference)
	}

	return externalReferences, nil
}

func (f *FirestoreService) GetExternalReference(
	projectID,
	companyID,
	pieceID,
	externalReferenceID string,
) (
	model.ExternalReference,
	error,
) {
	var externalReference model.ExternalReference
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return externalReference, err
	}
	defer client.Close()
	doc, err := client.
		Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("externalReferences").Doc(externalReferenceID).
		Get(ctx)
	if err != nil {
		return externalReference, err
	}

	if err := doc.DataTo(&externalReference); err != nil {
		return externalReference, err
	}

	return externalReference, nil
}

func (f *FirestoreService) AddExternalReference(
	projectID,
	companyID,
	pieceID string,
	externalReference model.ExternalReference,
) (	
	externalReferenceID string,
	err error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return externalReferenceID, err
	}
	defer client.Close()
	
	if externalReference.ID == "" {
		var ref, _, err = client.Collection("companies").Doc(companyID).
			Collection("pieces").Doc(pieceID).
			Collection("externalReferences").
			Add(ctx, externalReference)
		externalReference.ID = ref.ID

		client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("externalReferences").Doc(externalReference.ID).
		Set(ctx, map[string]interface{}{
        "id": externalReference.ID}, firestore.MergeAll)

		if err != nil {
			return externalReferenceID, err
		}
	} else {
		_, err = client.Collection("companies").Doc(companyID).
			Collection("pieces").Doc(pieceID).
			Collection("externalReferences").Doc(externalReference.ID).
			Set(ctx, externalReference)
		if err != nil {
			return externalReferenceID, err
		}
	}

	return externalReference.ID, nil
}

func (f *FirestoreService) UpdateExternalReference(
	projectID,
	companyID,
	pieceID,
	externalReferenceID string,
	externalReference model.ExternalReference,
) (
	error,
) {
	var mapExternalReference = utils.ToFirestoreMap(externalReference)
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, externalReferenceID)
	if err != nil {
		return err
	}
	defer client.Close()
	_, err = client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("externalReferences").Doc(externalReferenceID).
		Set(ctx, mapExternalReference, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteExternalReference(
	projectID,
	companyID,
	pieceID,
	externalReferenceID string,
) (
	error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, externalReferenceID)
	if err != nil {
		return err
	}
	defer client.Close()
	ref := client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("externalReferences").Doc(externalReferenceID)
	err = deleteDocumentAndSubcollections(ctx, client, ref)
	if err != nil {
		return err
	}

	return nil
}



func (f *FirestoreService) DeleteExternalReferenceFields(
	projectID,
	companyID,
	pieceID,
	externalReferenceID string,
	fields []string,
) (
	error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	updates := []firestore.Update{}

	for _, field := range fields {
		if !slices.Contains(strings.Split(field, "."), "id") {
			updates = append(updates, firestore.Update{Path: field, Value: firestore.Delete})
		}
	}

	_, err = client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Collection("externalReferences").Doc(externalReferenceID).
		Update(ctx, updates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) GetSecurityDeclaration(
	projectID,
	companyID,
	pieceID string,
) (
	model.SecurityDeclaration,
	error,
) {
	var securityDeclaration model.SecurityDeclaration
	var piece model.Piece
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return securityDeclaration, err
	}
	defer client.Close()
	doc, err := client.
		Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Get(ctx)
	if err != nil {
		return securityDeclaration, err
	}

	if err := doc.DataTo(&piece); err != nil {
		return securityDeclaration, err
	}

	return piece.SecurityDeclaration, nil
}

func (f *FirestoreService) AddSecurityDeclaration(
	projectID,
	companyID,
	pieceID string,
	securityDeclaration model.SecurityDeclaration,
) (	
	err error,
) {
	var piece model.Piece

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()
	
	doc, err := client.
		Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Get(ctx)
	if err != nil {
		return err
	}

	if err := doc.DataTo(&piece); err != nil {
		return err
	}

	piece.SecurityDeclaration = securityDeclaration
	
	_, err = client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Set(ctx, piece)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) UpdateSecurityDeclaration(
	projectID,
	companyID,
	pieceID string,
	securityDeclaration model.SecurityDeclaration,
) (
	error,
) {
	var piece model.Piece

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()
	
	doc, err := client.
		Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Get(ctx)
	if err != nil {
		return err
	}

	if err := doc.DataTo(&piece); err != nil {
		return err
	}

	piece.SecurityDeclaration = securityDeclaration
	var mapPiece = utils.ToFirestoreMap(piece)
	
	_, err = client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Set(ctx, mapPiece, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteSecurityDeclaration(
	projectID,
	companyID,
	pieceID string,
) (
	error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Update(ctx, []firestore.Update{{Path: "securityDeclaration", Value: firestore.Delete}})
	if err != nil {
		return err
	}

	return nil
}



func (f *FirestoreService) DeleteSecurityDeclarationFields(
	projectID,
	companyID,
	pieceID string,
	fields []string,
) (
	error,
) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	updates := []firestore.Update{}

	for _, field := range fields {
		if !slices.Contains(strings.Split(field, "."), "id") {
			updates = append(updates, firestore.Update{Path: "securityDeclaration." + field, Value: firestore.Delete})
		}
	}

	_, err = client.Collection("companies").Doc(companyID).
		Collection("pieces").Doc(pieceID).
		Update(ctx, updates)
	if err != nil {
		return err
	}

	return nil
}

func NewFirestoreService() (*FirestoreService, error) {
	var f FirestoreService
	return &f, nil
}

var _ service.DBService = (*FirestoreService)(nil)

func deleteDocumentAndSubcollections(ctx context.Context, client *firestore.Client,
	ref *firestore.DocumentRef) error {

	iter := ref.Collections(ctx)
	for {
		collRef, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		deleteCollection(ctx, client, collRef)
	}

	_, err := ref.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func deleteCollection(ctx context.Context, client *firestore.Client,
	ref *firestore.CollectionRef) error {

	for {
		// Get a batch of documents
		iter := ref.Limit(100).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}
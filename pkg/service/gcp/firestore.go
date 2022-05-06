package gcp

import (
	"context"
	"errors"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	"github.com/chi-deutschland/one-record-server/pkg/utils/conv"
	"golang.org/x/exp/slices"
	"google.golang.org/api/iterator"
)

type FirestoreService struct {
}

func NewFirestoreService() (*FirestoreService, error) {
	var f FirestoreService
	return &f, nil
}

var _ service.DBService = (*FirestoreService)(nil)

func (f *FirestoreService) GetCompanies(
	projectID,
	role,
	colPath string,
) (
	companies []model.Company,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return companies, err
	}
	defer client.Close()

	iter, err := GetDocuments(ctx, client, role, colPath)
	if err != nil {
		return companies, err
	}
	
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
	role,
	docPath string,
) (
	company model.Company,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return company, err
	}
	defer client.Close()

	doc, err := GetDocument(ctx, client, role, docPath)
	if err != nil {
		return company, err
	}

	if err := doc.DataTo(&company); err != nil {
		return company, err
	}

	return company, nil
}

func (f *FirestoreService) AddCompany(
	projectID,
	role,
	colPath,
	id string,
	company model.Company,
) (
	ID string,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return ID, err
	}
	defer client.Close()

	ID, err = AddDocument(ctx, client, role, colPath, id, company)
	if err != nil {
		return ID, err
	}

	return ID, nil
}

func (f *FirestoreService) UpdateCompany(
	projectID,
	role,
	docPath string,
	updates utils.FirestoreMap,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = UpdateDocument(ctx, client, role, docPath, updates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteCompany(
	projectID,
	role,
	docPath string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentAndSubcollections(ctx, client, role, docPath)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteCompanyFields(
	projectID,
	role,
	docPath string,
	fields []string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentFields(ctx, client, role, docPath, fields)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) GetPieces(
	projectID,
	role,
	colPath string,
) (
	pieces []model.Piece,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return pieces, err
	}
	defer client.Close()

	iter, err := GetDocuments(ctx, client, role, colPath)
	if err != nil {
		return pieces, err
	}
	
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
	role,
	docPath string,
) (
	piece model.Piece,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return piece, err
	}
	defer client.Close()

	doc, err := GetDocument(ctx, client, role, docPath)
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
	role,
	colPath,
	id string,
	piece model.Piece,
) (
	ID string,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return ID, err
	}
	defer client.Close()

	ID, err = AddDocumentWithCollection(ctx, client, role, colPath, id, piece)
	if err != nil {
		return ID, err
	}

	return ID, nil
}

func (f *FirestoreService) UpdatePiece(
	projectID,
	role,
	docPath string,
	updates utils.FirestoreMap,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = UpdateDocument(ctx, client, role, docPath, updates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeletePiece(
	projectID,
	role,
	docPath string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentAndSubcollections(ctx, client, role, docPath)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeletePieceFields(
	projectID,
	role,
	docPath string,
	fields []string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentFields(ctx, client, role, docPath, fields)
	if err != nil {
		return err
	}

	return nil
}


func (f *FirestoreService) GetEvents(
	projectID,
	role,
	colPath string,
) (
	events []model.Event,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return events, err
	}
	defer client.Close()

	iter, err := GetDocuments(ctx, client, role, colPath)
	if err != nil {
		return events, err
	}
	
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
	role,
	docPath string,
) (
	event model.Event,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return event, err
	}
	defer client.Close()

	doc, err := GetDocument(ctx, client, role, docPath)
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
	role,
	colPath,
	id string,
	event model.Event,
) (
	ID string,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return ID, err
	}
	defer client.Close()

	ID, err = AddDocumentWithCollection(ctx, client, role, colPath, id, event)
	if err != nil {
		return ID, err
	}

	return ID, nil
}

func (f *FirestoreService) UpdateEvent(
	projectID,
	role,
	docPath string,
	updates utils.FirestoreMap,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = UpdateDocument(ctx, client, role, docPath, updates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteEvent(
	projectID,
	role,
	docPath string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentAndSubcollections(ctx, client, role, docPath)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteEventFields(
	projectID,
	role,
	docPath string,
	fields []string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentFields(ctx, client, role, docPath, fields)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) GetExternalReferences(
	projectID,
	role,
	colPath string,
) (
	externalReferences []model.ExternalReference,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return externalReferences, err
	}
	defer client.Close()

	iter, err := GetDocuments(ctx, client, role, colPath)
	if err != nil {
		return externalReferences, err
	}
	
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
	role,
	docPath string,
) (
	externalReference model.ExternalReference,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return externalReference, err
	}
	defer client.Close()

	doc, err := GetDocument(ctx, client, role, docPath)
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
	role,
	colPath,
	id string,
	externalReference model.ExternalReference,
) (
	ID string,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return ID, err
	}
	defer client.Close()

	ID, err = AddDocumentWithCollection(ctx, client, role, colPath, id, externalReference)
	if err != nil {
		return ID, err
	}

	return ID, nil
}

func (f *FirestoreService) UpdateExternalReference(
	projectID,
	role,
	docPath string,
	updates utils.FirestoreMap,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = UpdateDocument(ctx, client, role, docPath, updates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteExternalReference(
	projectID,
	role,
	docPath string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentAndSubcollections(ctx, client, role, docPath)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteExternalReferenceFields(
	projectID,
	role,
	docPath string,
	fields []string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentFields(ctx, client, role, docPath, fields)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) GetSecurityDeclarations(
	projectID,
	role,
	colPath string,
) (
	securityDeclarations []model.SecurityDeclaration,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return securityDeclarations, err
	}
	defer client.Close()

	iter, err := GetDocuments(ctx, client, role, colPath)
	if err != nil {
		return securityDeclarations, err
	}
	
	for {
		var securityDeclaration model.SecurityDeclaration

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return securityDeclarations, err
		}

		if err := doc.DataTo(&securityDeclaration); err != nil {
			return securityDeclarations, err
		}

		securityDeclarations = append(securityDeclarations, securityDeclaration)
	}

	return securityDeclarations, nil
}

func (f *FirestoreService) GetSecurityDeclaration(
	projectID,
	role,
	docPath string,
) (
	securityDeclaration model.SecurityDeclaration,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return securityDeclaration, err
	}
	defer client.Close()

	doc, err := GetDocument(ctx, client, role, docPath)
	if err != nil {
		return securityDeclaration, err
	}

	if err := doc.DataTo(&securityDeclaration); err != nil {
		return securityDeclaration, err
	}

	return securityDeclaration, nil
}

func (f *FirestoreService) AddSecurityDeclaration(
	projectID,
	role,
	colPath,
	id string,
	securityDeclaration model.SecurityDeclaration,
) (
	ID string,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return ID, err
	}
	defer client.Close()

	ID, err = AddDocumentWithCollection(ctx, client, role, colPath, id, securityDeclaration)
	if err != nil {
		return ID, err
	}

	return ID, nil
}

func (f *FirestoreService) UpdateSecurityDeclaration(
	projectID,
	role,
	docPath string,
	updates utils.FirestoreMap,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = UpdateDocument(ctx, client, role, docPath, updates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteSecurityDeclaration(
	projectID,
	role,
	docPath string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentAndSubcollections(ctx, client, role, docPath)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteSecurityDeclarationFields(
	projectID,
	role,
	docPath string,
	fields []string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentFields(ctx, client, role, docPath, fields)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) GetShipments(
	projectID,
	role,
	colPath string,
) (
	shipments []model.Shipment,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return shipments, err
	}
	defer client.Close()

	iter, err := GetDocuments(ctx, client, role, colPath)
	if err != nil {
		return shipments, err
	}
	
	for {
		var shipment model.Shipment

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return shipments, err
		}

		if err := doc.DataTo(&shipment); err != nil {
			return shipments, err
		}

		shipments = append(shipments, shipment)
	}

	return shipments, nil
}

func (f *FirestoreService) GetShipment(
	projectID,
	role,
	docPath string,
) (
	shipment model.Shipment,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return shipment, err
	}
	defer client.Close()

	doc, err := GetDocument(ctx, client, role, docPath)
	if err != nil {
		return shipment, err
	}

	if err := doc.DataTo(&shipment); err != nil {
		return shipment, err
	}

	return shipment, nil
}

func (f *FirestoreService) AddShipment(
	projectID,
	role,
	colPath,
	id string,
	shipment model.Shipment,
) (
	ID string,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return ID, err
	}
	defer client.Close()

	ID, err = AddDocumentWithCollection(ctx, client, role, colPath, id, shipment)
	if err != nil {
		return ID, err
	}

	return ID, nil
}

func (f *FirestoreService) UpdateShipment(
	projectID,
	role,
	docPath string,
	updates utils.FirestoreMap,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = UpdateDocument(ctx, client, role, docPath, updates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteShipment(
	projectID,
	role,
	docPath string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentAndSubcollections(ctx, client, role, docPath)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteShipmentFields(
	projectID,
	role,
	docPath string,
	fields []string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentFields(ctx, client, role, docPath, fields)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) GetTransportMovements(
	projectID,
	role,
	colPath string,
) (
	transportMovements []model.TransportMovement,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return transportMovements, err
	}
	defer client.Close()

	iter, err := GetDocuments(ctx, client, role, colPath)
	if err != nil {
		return transportMovements, err
	}
	
	for {
		var transportMovement model.TransportMovement

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return transportMovements, err
		}

		if err := doc.DataTo(&transportMovement); err != nil {
			return transportMovements, err
		}

		transportMovements = append(transportMovements, transportMovement)
	}

	return transportMovements, nil
}

func (f *FirestoreService) GetTransportMovement(
	projectID,
	role,
	docPath string,
) (
	transportMovement model.TransportMovement,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return transportMovement, err
	}
	defer client.Close()

	doc, err := GetDocument(ctx, client, role, docPath)
	if err != nil {
		return transportMovement, err
	}

	if err := doc.DataTo(&transportMovement); err != nil {
		return transportMovement, err
	}

	return transportMovement, nil
}

func (f *FirestoreService) AddTransportMovement(
	projectID,
	role,
	colPath,
	id string,
	transportMovement model.TransportMovement,
) (
	ID string,
	err error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return ID, err
	}
	defer client.Close()

	ID, err = AddDocumentWithCollection(ctx, client, role, colPath, id, transportMovement)
	if err != nil {
		return ID, err
	}

	return ID, nil
}

func (f *FirestoreService) UpdateTransportMovement(
	projectID,
	role,
	docPath string,
	updates utils.FirestoreMap,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = UpdateDocument(ctx, client, role, docPath, updates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteTransportMovement(
	projectID,
	role,
	docPath string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentAndSubcollections(ctx, client, role, docPath)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreService) DeleteTransportMovementFields(
	projectID,
	role,
	docPath string,
	fields []string,
) (
	error,
) {
	// Set up client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	err = DeleteDocumentFields(ctx, client, role, docPath, fields)
	if err != nil {
		return err
	}

	return nil
}

func GetDocuments(
	ctx context.Context,
	client *firestore.Client,
	role,
	colPath string,
) (
	iter *firestore.DocumentIterator,
	err error,
) {
	colRef := client.Collection(colPath)

	// Get document iterator
	iter = colRef.Documents(ctx)

	return iter, nil
}

func GetDocument(
	ctx context.Context,
	client *firestore.Client,
	role,
	docPath string,
) (
	doc *firestore.DocumentSnapshot,
	err error,
) {
	docRef := client.Doc(docPath)
	// Get document
	doc, err = docRef.Get(ctx)
	if err != nil {
		return doc, err
	}

	return doc, nil
}

func AddDocumentWithCollection(
	ctx context.Context,
	client *firestore.Client,
	role,
	colPath,
	id string,
	data interface{},
) (
	ID string,
	err error,
) {
	splitPath := strings.Split(colPath, "/")
	last := len(splitPath) - 1
	parentDocPath, colName := strings.Join(splitPath[:last], "/"), splitPath[last]
	parentDocRef := client.Doc(parentDocPath)
	// Check whether parent document exists
	_, err = parentDocRef.Get(ctx)
	if err != nil {
		return ID, err
	}

	colRef := parentDocRef.Collection(colName)

	// If no id is passed
	if id == "" {
		// Add data and generate ID
		var ref, _, err = colRef.Add(ctx, data)
		if err != nil {
			return ID, err
		}

		newDocRef := colRef.Doc(ref.ID)

		// Add id to new data doc
		_, err = newDocRef.Set(ctx, map[string]interface{}{"id": ref.ID}, firestore.MergeAll)
		if err != nil {
			return ID, err
		}

		// Return new ID
		return ref.ID, nil

	// If id is passed
	} else {
		docRef := colRef.Doc(id)
		// Check whether document exists
		_, err := docRef.Get(ctx)
		if err == nil {
			return ID, errors.New("document already exists")
		}

		// Create the document
		_, err = colRef.Doc(id).Set(ctx, data)
		if err != nil {
			return ID, err
		}

		// Return ID
		return id, nil
	}
}

func AddDocument(
	ctx context.Context,
	client *firestore.Client,
	role,
	colPath,
	id string,
	data interface{},
) (
	ID string,
	err error,
) {
	colRef := client.Collection(colPath)

	// If no id is passed
	if id == "" {
		// Add data and generate ID
		var ref, _, err = colRef.Add(ctx, data)
		if err != nil {
			return ID, err
		}

		newDocRef := colRef.Doc(ref.ID)

		// Add id to new data doc
		_, err = newDocRef.Set(ctx, map[string]interface{}{"id": ref.ID}, firestore.MergeAll)
		if err != nil {
			return ID, err
		}

		// Return new ID
		return ref.ID, nil

	// If id is passed
	} else {
		docRef := colRef.Doc(id)
		// Check whether document exists
		_, err := docRef.Get(ctx)
		if err == nil {
			return ID, errors.New("document already exists")
		}

		// Create the document
		_, err = colRef.Doc(id).Set(ctx, data)
		if err != nil {
			return ID, err
		}

		// Return ID
		return id, nil
	}
}

func UpdateDocument(
	ctx context.Context,
	client *firestore.Client,
	role,
	docPath string,
	updates utils.FirestoreMap,
) (
	error,
) {
	docRef := client.Doc(docPath)
	// Check whether document exists
	_, err := docRef.Get(ctx)
	if err != nil {
		return err
	}

	// Overwrite passed fields
	_, err = docRef.Set(ctx, updates, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func DeleteDocumentAndSubcollections(
	ctx context.Context,
	client *firestore.Client,
	role,
	docPath string,
) (
	error,
) {
	docRef := client.Doc(docPath)
	// Check whether document exists
	_, err := docRef.Get(ctx)
	if err != nil {
		return err
	}

	iter := docRef.Collections(ctx)
	for {
		collRef, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		DeleteCollection(ctx, client, role, collRef)
	}

	_, err = docRef.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCollection(
	ctx context.Context,
	client *firestore.Client,
	role string,
	ref *firestore.CollectionRef,
) (
	error,
) {
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

func DeleteDocumentFields(
	ctx context.Context,
	client *firestore.Client,
	role,
	docPath string,
	fields []string,
) (
	error,
) {
	docRef := client.Doc(docPath)
	// Check whether document exists
	_, err := docRef.Get(ctx)
	if err != nil {
		return err
	}

	updates := []firestore.Update{}
	// Filter  all fields that contain "id" after splitting by "."
	for _, field := range fields {
		if !slices.Contains(strings.Split(field, "."), "id") {
			updates = append(updates, firestore.Update{Path: field, Value: firestore.Delete})
		}
	}

	// Delete the fields
	_, err = docRef.Update(ctx, updates)
	if err != nil {
		return err
	}

	return nil
}
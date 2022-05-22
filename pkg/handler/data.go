package handler

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
    "io/ioutil"
	"github.com/Meschkov/jsonld"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type DataHandler struct {
	Service *service.Service
}

func (h *DataHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPOST Data")
		logger.Infof("Received request with params %#v", r.URL.Path)

		err := CreateData(h)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	case "DELETE":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nDELETE Data")
		logger.Infof("Received request with params %#v", r.URL.Path)
		
		err := h.Service.DBService.DeleteCollectionRecursiveGivenPath(h.Service.Env.ProjectId, h.Service.Env.ServerRole, "companies")
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func NewDataHandler(svc *service.Service) *DataHandler {
	return &DataHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*DataHandler)(nil)

func GetFiles(dir, ext string) (files []string) {
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(d.Name()) == ext {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func CreateData(h *DataHandler) error {
	err := CreateCompanies(h)
	if err != nil {
		return err
	}

	err = CreatePieces(h)
	if err != nil {
		return err
	}

	err = CreateEvents(h)
	if err != nil {
		return err
	}

	err = CreateExternalReferences(h)
	if err != nil {
		return err
	}

	err = CreateSecurityDeclarations(h)
	if err != nil {
		return err
	}

	err = CreateShipments(h)
	if err != nil {
		return err
	}

	err = CreateTransportMovements(h)
	if err != nil {
		return err
	}

	err = CreateItems(h)
	if err != nil {
		return err
	}

	err = CreateRegulatedEntities(h)
	if err != nil {
		return err
	}

	return nil
}

func CreateCompanies(h *DataHandler) error {
	files := GetFiles("data/companies", ".json")

	for _, file := range files {
        f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		var objmap map[string]json.RawMessage
		data, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(data, &objmap)
		if err != nil {
			return err
		}

		var path string
		err = json.Unmarshal(objmap["path"], &path)
		if err != nil {
			return err
		}

		var company model.Company
		err = jsonld.UnmarshalCompacted(objmap["obj"], &company)
		if err != nil {
			return err
		}

		_, err = h.Service.DBService.AddCompany(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, company.ID, company)
		if err != nil {
			return err
		}
    }

	return nil
}

func CreatePieces(h *DataHandler) error {
	files := GetFiles("data/pieces", ".json")

	for _, file := range files {
        f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		var objmap map[string]json.RawMessage
		data, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(data, &objmap)
		if err != nil {
			return err
		}

		var path string
		err = json.Unmarshal(objmap["path"], &path)
		if err != nil {
			return err
		}

		var piece model.Piece
		err = jsonld.UnmarshalCompacted(objmap["obj"], &piece)
		if err != nil {
			return err
		}

		_, err = h.Service.DBService.AddPiece(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, piece.ID, piece)
		if err != nil {
			return err
		}
    }

	return nil
}

func CreateEvents(h *DataHandler) error {
	files := GetFiles("data/events", ".json")

	for _, file := range files {
        f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		var objmap map[string]json.RawMessage
		data, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(data, &objmap)
		if err != nil {
			return err
		}

		var path string
		err = json.Unmarshal(objmap["path"], &path)
		if err != nil {
			return err
		}

		var event model.Event
		err = jsonld.UnmarshalCompacted(objmap["obj"], &event)
		if err != nil {
			return err
		}

		_, err = h.Service.DBService.AddEvent(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, event.ID, event)
		if err != nil {
			return err
		}
    }

	return nil
}

func CreateExternalReferences(h *DataHandler) error {
	files := GetFiles("data/externalReferences", ".json")

	for _, file := range files {
        f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		var objmap map[string]json.RawMessage
		data, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(data, &objmap)
		if err != nil {
			return err
		}

		var path string
		err = json.Unmarshal(objmap["path"], &path)
		if err != nil {
			return err
		}

		var externalReference model.ExternalReference
		err = jsonld.UnmarshalCompacted(objmap["obj"], &externalReference)
		if err != nil {
			return err
		}

		_, err = h.Service.DBService.AddExternalReference(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, externalReference.ID, externalReference)
		if err != nil {
			return err
		}
    }

	return nil
}

func CreateSecurityDeclarations(h *DataHandler) error {
	files := GetFiles("data/securityDeclarations", ".json")

	for _, file := range files {
        f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		var objmap map[string]json.RawMessage
		data, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(data, &objmap)
		if err != nil {
			return err
		}

		var path string
		err = json.Unmarshal(objmap["path"], &path)
		if err != nil {
			return err
		}

		var securityDeclaration model.SecurityDeclaration
		err = jsonld.UnmarshalCompacted(objmap["obj"], &securityDeclaration)
		if err != nil {
			return err
		}

		_, err = h.Service.DBService.AddSecurityDeclaration(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, securityDeclaration.ID, securityDeclaration)
		if err != nil {
			return err
		}
    }

	return nil
}

func CreateShipments(h *DataHandler) error {
	files := GetFiles("data/shipments", ".json")

	for _, file := range files {
        f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		var objmap map[string]json.RawMessage
		data, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(data, &objmap)
		if err != nil {
			return err
		}

		var path string
		err = json.Unmarshal(objmap["path"], &path)
		if err != nil {
			return err
		}

		var shipment model.Shipment
		err = jsonld.UnmarshalCompacted(objmap["obj"], &shipment)
		if err != nil {
			return err
		}

		_, err = h.Service.DBService.AddShipment(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, shipment.ID, shipment)
		if err != nil {
			return err
		}
    }

	return nil
}

func CreateTransportMovements(h *DataHandler) error {
	files := GetFiles("data/transportMovements", ".json")

	for _, file := range files {
        f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		var objmap map[string]json.RawMessage
		data, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(data, &objmap)
		if err != nil {
			return err
		}

		var path string
		err = json.Unmarshal(objmap["path"], &path)
		if err != nil {
			return err
		}

		var transportMovement model.TransportMovement
		err = jsonld.UnmarshalCompacted(objmap["obj"], &transportMovement)
		if err != nil {
			return err
		}

		_, err = h.Service.DBService.AddTransportMovement(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, transportMovement.ID, transportMovement)
		if err != nil {
			return err
		}
    }

	return nil
}

func CreateRegulatedEntities(h *DataHandler) error {
	files := GetFiles("data/regulatedEntities", ".json")

	for _, file := range files {
        f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		var objmap map[string]json.RawMessage
		data, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(data, &objmap)
		if err != nil {
			return err
		}

		var path string
		err = json.Unmarshal(objmap["path"], &path)
		if err != nil {
			return err
		}

		var regulatedEntity model.RegulatedEntity
		err = jsonld.UnmarshalCompacted(objmap["obj"], &regulatedEntity)
		if err != nil {
			return err
		}

		_, err = h.Service.DBService.AddRegulatedEntity(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, regulatedEntity.ID, regulatedEntity)
		if err != nil {
			return err
		}
    }

	return nil
}

func CreateItems(h *DataHandler) error {
	files := GetFiles("data/items", ".json")

	for _, file := range files {
        f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		var objmap map[string]json.RawMessage
		data, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(data, &objmap)
		if err != nil {
			return err
		}

		var path string
		err = json.Unmarshal(objmap["path"], &path)
		if err != nil {
			return err
		}

		var item model.Item
		err = jsonld.UnmarshalCompacted(objmap["obj"], &item)
		if err != nil {
			return err
		}

		_, err = h.Service.DBService.AddItem(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, item.ID, item)
		if err != nil {
			return err
		}
    }

	return nil
}
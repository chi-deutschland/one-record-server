package handler

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

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

type ObjPath struct {
	Path string `json:"path"`
}

func CreateData(h *DataHandler) error {
	err := CreateCompanies(h)
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

		decoder := json.NewDecoder(f)
		var company model.Company
		var objPath ObjPath
		err = decoder.Decode(&company)
		if err != nil {
			return err
		}
		err = decoder.Decode(&objPath)
		if err != nil {
			return err
		}

		_, err = h.Service.DBService.AddCompany(h.Service.Env.ProjectId, h.Service.Env.ServerRole, objPath.Path, company.ID, company)
		if err != nil {
			return err
		}
    }

	return nil
}

package dinoapi

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shunsukw/golang-practice/dino/databaselayer"
)

func RunApi(endpoint string, db databaselayer.DinoDBHandler) error {
	r := mux.NewRouter()
	RunAPIOnRouter(r, db)
	return http.ListenAndServe(endpoint, r)
}

func RunAPIOnRouter(r *mux.Router, db databaselayer.DinoDBHandler) {
	handler := newDinoRESTAPIHandler(db)

	apirouter := r.PathPrefix("/api/dinos").Subrouter()

	apirouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.searchHandler)
	apirouter.Methods("POST").PathPrefix("/{Operation}").HandlerFunc(handler.editsHandler)
}

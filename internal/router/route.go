package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {

	route := mux.NewRouter().StrictSlash(true)

	h := &Handler{store: db}

	route.HandleFunc("/", h.getItems).Methods("GET")
	route.HandleFunc("/", h.addItems).Methods("POST")
	route.HandleFunc("/", h.deleteItems).Methods("DELETE")
	route.HandleFunc("/", h.updateItems).Methods("PUT")
	route.NotFoundHandler = http.HandlerFunc(h.notFound)

	return route
}

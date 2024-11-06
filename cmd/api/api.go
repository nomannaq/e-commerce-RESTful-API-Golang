package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nomannaq/e-commerce-restfulAPI-go/cmd/services/user"
)

type APIserver struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIserver {
	return &APIserver{
		addr: addr,
		db:   db,
	}
}

func (s *APIserver) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	return http.ListenAndServe(s.addr, router)

}

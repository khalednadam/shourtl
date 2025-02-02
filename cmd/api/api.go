package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func (s *APIServer) HandleFunc(param1 string, f func(w http.ResponseWriter, r *http.Request)) {
	panic("unimplemented")
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}
	fmt.Println("Listening on port", server.Addr)
	return server.ListenAndServe()
}

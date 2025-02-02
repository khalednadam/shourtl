package api

import (
	"fmt"
	"net/http"
)

type APIServer struct {
	addr   string
	router *http.ServeMux
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr:   addr,
		router: http.NewServeMux(),
	}
}

func (s *APIServer) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.router.HandleFunc(pattern, handler)
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

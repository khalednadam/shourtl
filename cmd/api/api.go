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
	server := http.Server{
		Addr:    s.addr,
		Handler: s.router,
	}
	fmt.Println("Listening on port", s.addr)
	return server.ListenAndServe()
}

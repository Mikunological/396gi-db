package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	router     *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() error {
	s.router = s.Route()
	return nil
}

func (s *Server) Run(port int) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	router := mux.NewRouter()
	r.Methods("GET").Path("/").HandlerFunc(echoHello)
}

func echoHello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello World</h1>")
}
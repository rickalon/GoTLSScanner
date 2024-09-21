package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	addr string
}

func NewServer(addr ...string) *Server {
	primary := &Server{addr: ":8080"}
	for _, val := range addr {
		primary.addr = val
	}
	return primary
}

func (server *Server) Run(router ...*mux.Router) error {
	if len(router) == 0 {
		return http.ListenAndServe(server.addr, nil)
	} else {
		return http.ListenAndServe(server.addr, router[0])
	}
}

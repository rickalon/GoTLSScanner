package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Rt *mux.Router
}

func NewRouter() *Router {
	return &Router{Rt: mux.NewRouter()}
}

func (router *Router) CreateSubrouter(prefix string) *mux.Router {
	return router.Rt.PathPrefix(prefix).Subrouter()
}

func (router *Router) NewHandlerGet(rt *mux.Router, rtPath string, fn http.HandlerFunc) {
	rt.HandleFunc(rtPath, fn).Methods("GET")
}

func (router *Router) NewHandlerPost(rt *mux.Router, rtPath string, fn http.HandlerFunc) {
	rt.HandleFunc(rtPath, fn).Methods("POST")
}

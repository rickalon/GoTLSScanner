package main

import (
	"log"

	"github.com/rickalon/GoWebScraper/handlers"
	"github.com/rickalon/GoWebScraper/server"
)

func main() {
	srv := server.NewServer()
	router := server.NewRouter()
	subroute := router.CreateSubrouter("/api/v1")
	router.NewHandlerGet(subroute, "/start", handlers.MainHandler)
	log.Fatal(srv.Run(router.Rt))
}

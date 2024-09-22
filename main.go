package main

import (
	"context"
	"log"

	"github.com/rickalon/GoWebScraper/db"
	"github.com/rickalon/GoWebScraper/handlers"
	"github.com/rickalon/GoWebScraper/server"
	"github.com/rickalon/GoWebScraper/settings"
)

func main() {
	mongoDBConfig := settings.GetEnvMongo()
	database, err := db.NewMongoDB(context.Background(), mongoDBConfig)
	database.CreateDB()
	if err != nil {
		log.Println("Persistance not avaible")
		database.IsOn = false
	} else {
		database.IsOn = true
	}
	srv := server.NewServer()
	router := server.NewRouter()
	subroute := router.CreateSubrouter("/api/v1")
	router.NewHandlerGet(subroute, "/test", handlers.TestHandler)
	router.NewHandlerPost(subroute, "/analyze", handlers.URLDBHandler(database))
	log.Fatal(srv.Run(router.Rt))
}

package main

import (
	api "app/api/handlers"
	config "app/cofig"
	"app/pkg/db"
	"log"


)

func main() {

	cfg := config.NewConfig()
	log.Println("successfully loaded config")

	con, err := db.Conn(cfg)

	if err != nil {

		log.Println("error to connect db", err)
		return
	}

	log.Println("successfully connected to db")

	storage := storage.NewStorage(con)

	api.Api(storage)

}

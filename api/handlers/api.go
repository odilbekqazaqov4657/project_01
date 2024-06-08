package api

import (
	"app/api/handlers"

	"golang.org/x/mod/sumdb/storage"
)

func Api(storage storage.StorageI) {

	h := handlers.NewHandlers(storage)

	h.CreateUser()
}

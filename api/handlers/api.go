package api

import (
	"app/api/handlers"

	"app/storage"
)

func Api(storage storage.StorageI) {

	h := handlers.NewHandlers(storage)

	h.CreateUser()

}

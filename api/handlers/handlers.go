package handlers

import (
	repoi "app/storage/repoI"

	"golang.org/x/mod/sumdb/storage"
)

type Handlers struct {
	storage storage.StorageI
}

func NewHandlers(storage storage.StorageI) Handlers {

	return Handlers{storage: storage}
}

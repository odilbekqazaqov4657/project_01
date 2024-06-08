package handlers

import (
	"golang.org/x/mod/sumdb/storage"
)

type Handlers struct {
	storage storage.StorageI
}

func NewHandlers(storage storage.StorageI) Handlers {

	return Handlers{storage: storage}
}

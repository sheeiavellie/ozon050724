package graph

import "github.com/sheeiavellie/ozon050724/storage"

type resolver struct {
	storage storage.Storage
}

func NewResolver(storage storage.Storage) *resolver {
	return &resolver{
		storage: storage,
	}
}

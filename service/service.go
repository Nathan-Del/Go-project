package service

import "github.com/estiam/ticketing/db"

type Service struct {
	db db.Store
}

func New(db db.Store) *Service {
	return &Service{
		db: db,
	}
}

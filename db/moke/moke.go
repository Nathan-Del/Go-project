package moke

import (
	"sync"

	"github.com/estiam/ticketing/db"
	"github.com/estiam/ticketing/model"
)

var _ db.Store = &DB{}

// DB contient la connexion à la BDD
type DB struct {
	mx         *sync.Mutex
	UserList   map[string]*model.User
	TicketList map[string]*model.Ticket
}

// New est la création d'une nouvelle connexion à la base de données
func New() *DB {
	return &DB{
		UserList:   make(map[string]*model.User),
		TicketList: make(map[string]*model.Ticket),
	}
}

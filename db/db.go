// Le package db contient la connexion db et les CRUDs associés.
package db

import (
	"github.com/estiam/ticketing/model"
)

type Store interface {
	StoreUser
	StoreTicket
}

type StoreUser interface {
	GetAllUser() ([]*model.User, error)
	CreateUser(u *model.User) (*model.User, error)
	DeleteUser(id string) error
	UpdateUser(id string, data map[string]interface{}) (*model.User, error)
	GetUserById(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

type StoreTicket interface {
	GetAllTicket() ([]*model.Ticket, error)
	CreateTicket(u *model.Ticket) (*model.Ticket, error)
	DeleteTicket(id string) error
	UpdateTicket(id string, data map[string]interface{}) (*model.Ticket, error)
	GetTicketById(id string) (*model.Ticket, error)
	GetTicketsByUserId(id string) ([]model.Ticket, error)
}

package moke

import (
	"errors"

	"github.com/estiam/ticketing/model"
	"github.com/google/uuid"
)

// GetAllTicket récupère tous les tickets de la BDD
func (db *DB) GetAllTicket() ([]*model.Ticket, error) {
	list := make([]*model.Ticket, len(db.TicketList))
	i := 0
	for k := range db.TicketList {
		list[i] = db.TicketList[k]
		i++
	}
	return list, nil
}

// CreateTicket créer un nouveau ticket dans la BDD
func (db *DB) CreateTicket(u *model.Ticket) (*model.Ticket, error) {
	u.ID = uuid.New().String()
	db.TicketList[u.ID] = u
	return u, nil
}

// DeleteTicket supprime un ticket avec son id dans la BDD
func (db *DB) DeleteTicket(id string) error {
	_, ok := db.TicketList[id]
	if !ok {
		return errors.New("db: user not found")
	}
	delete(db.TicketList, id)
	return nil
}

// UpdateTicket metà jour un ticket avec son id
func (db *DB) UpdateTicket(id string, data map[string]interface{}) (*model.Ticket, error) {
	//Permet de crer une transaction
	db.mx.Lock()
	//defer s'éxécute à la fin de la fonction
	defer db.mx.Unlock()

	u, err := db.GetTicketById(id)
	if err != nil {
		return nil, err
	}
	setStringFromMap(data, "title", &u.Title)
	setStringFromMap(data, "description", &u.Description)

	// Si le status est à "Ouvert" il ne peut être changer que en "En cours"
	if u.Status == "Ouvert" && (data["status"] != "En cours" || data["status"] == "Ouvert") {
		return nil, errors.New("db : changement de status non autorisé")
	}
	// Si le status est à "En cours" il ne peut être changer que en "Terminer"
	if u.Status == "En cours" && (data["status"] != "Terminer" || data["status"] == "En cours") {
		return nil, errors.New("db : changement de status non autorisé")
	}
	// Si le status est à "Terminer" il ne peut plus être changer
	if u.Status == "Terminer" && data["status"] != "Terminer" {
		return nil, errors.New("db : changement de status non autorisé")
	}
	setStringFromMap(data, "status", &u.Status)
	setStringFromMap(data, "id_user", &u.IdUser)
	return nil, nil
}

// GetTicketById récupère à partir de l'id donne le ticket de la BDD
func (db *DB) GetTicketById(id string) (*model.Ticket, error) {
	u, ok := db.TicketList[id]
	if !ok {
		return nil, errors.New("db: not found")
	}
	return u, nil
}

// GetTicketById à partir de l'id du user donne le ticket de la BDD
func (db *DB) GetTicketsByUserId(id string) ([]model.Ticket, error) {
	var result []model.Ticket
	for k := range db.TicketList {
		if db.TicketList[k].IdUser == id {
			result = append(result, *db.TicketList[k])

		}
	}
	return result, nil
	//return nil, errors.New("db: not found")
}

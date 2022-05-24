package sqlite

import (
	"github.com/estiam/ticketing/model"
	"github.com/google/uuid"
)

// utilise gorm pour trouver tous les tickets avec Find()
func (db *DB) GetAllTicket() ([]*model.Ticket, error) {
	var us []*model.Ticket
	err := db.conn.Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

// utilise gorm pour créer le ticket avec Create()
func (db *DB) CreateTicket(u *model.Ticket) (*model.Ticket, error) {
	u.ID = uuid.New().String()
	err := db.conn.Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// utilise gorm pour trouver avec Where() et Delete() pour supprimer l'élément trouver
func (db *DB) DeleteTicket(id string) error {
	var u model.Ticket
	return db.conn.Where("id = ?", id).Delete(&u).Error
}

// utilise gorm pour trouver avec Where() et Update() pour mettre à jour l'élément trouver
func (db *DB) UpdateTicket(id string, data map[string]interface{}) (*model.Ticket, error) {
	err := db.conn.Model(&model.Ticket{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return db.GetTicketById(id)
}

// utilise gorm pour trouver le premier élément qui correspond a l'id demander avec First()
func (db *DB) GetTicketById(id string) (*model.Ticket, error) {
	var u model.Ticket
	err := db.conn.First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// utiliser pour trouver les tickets qui ont le "id_user" demander en paramètre
func (db *DB) GetTicketsByUserId(id string) ([]model.Ticket, error) {
	var ts []model.Ticket
	err := db.conn.Where("id_user = ?", id).Find(&ts).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

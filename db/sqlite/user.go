package sqlite

import (
	"github.com/estiam/ticketing/model"
	"github.com/google/uuid"
)

// utilise gorm pour trouver tous les users avec Find()
func (db *DB) GetAllUser() ([]*model.User, error) {
	var us []*model.User
	err := db.conn.Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

// utilise gorm pour créer le user demander avec Create()
func (db *DB) CreateUser(u *model.User) (*model.User, error) {
	u.ID = uuid.New().String()
	err := db.conn.Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// utilise gorm pour trouver avec Where() et Delete() pour supprimer le user demander avec l'id
func (db *DB) DeleteUser(id string) error {
	var u model.User
	return db.conn.Where("id = ?", id).Delete(&u).Error
}

// utilise gorm pour trouver avec Where() et Update() pour mettre à jour le user
func (db *DB) UpdateUser(id string, data map[string]interface{}) (*model.User, error) {
	err := db.conn.Model(&model.User{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return db.GetUserById(id)
}

// utilise gorm pour trouver le premier user qui correspond a l'ID demander avec First()
func (db *DB) GetUserById(id string) (*model.User, error) {
	var u model.User
	err := db.conn.First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// utilise gorm pour trouver le premier user qui correspond a l'EMAIL demander avec First()
func (db *DB) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	err := db.conn.First(&u, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

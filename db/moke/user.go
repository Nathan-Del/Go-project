package moke

import (
	"errors"

	"github.com/estiam/ticketing/model"
	"github.com/google/uuid"
)

// GetAllUser récupère tous les users de la BDD
func (db *DB) GetAllUser() ([]*model.User, error) {
	list := make([]*model.User, len(db.UserList))
	i := 0
	for k := range db.UserList {
		list[i] = db.UserList[k]
		i++
	}
	return list, nil
}

// CreateUser créer un nouvel user dans la BDD
func (db *DB) CreateUser(u *model.User) (*model.User, error) {
	u.ID = uuid.New().String()
	db.UserList[u.ID] = u
	return u, nil
}

// DeleteUser supprime un user avec son id dans la BDD
func (db *DB) DeleteUser(id string) error {
	_, ok := db.UserList[id]
	if !ok {
		return errors.New("db: user not found")
	}
	delete(db.UserList, id)
	return nil
}

// UpdateUser met à jour un user avec son id
func (db *DB) UpdateUser(id string, data map[string]interface{}) (*model.User, error) {
	u, err := db.GetUserById(id)
	if err != nil {
		return nil, err
	}
	setStringFromMap(data, "first_name", &u.FirstName)
	setStringFromMap(data, "last_name", &u.LastName)
	setStringFromMap(data, "email", &u.Email)
	return nil, nil
}

// GetUserById à partir de l'id donné le user de la BDD
func (db *DB) GetUserById(id string) (*model.User, error) {
	u, ok := db.UserList[id]
	if !ok {
		return nil, errors.New("db: not found")
	}
	return u, nil
}

// GetUserById à partir de l'email donné le user de la BDD
func (db *DB) GetUserByEmail(email string) (*model.User, error) {
	for k := range db.UserList {
		if db.UserList[k].Email == email {
			return db.UserList[k], nil
		}
	}
	return nil, errors.New("db: user not found")

}

// setStringFromMap fait correspondre la map et la key.
func setStringFromMap(m map[string]interface{}, key string, fieldPt *string) {
	if v, ok := m[key]; ok {
		value, ok := v.(string)
		if ok {
			*fieldPt = value
		}
	}
}

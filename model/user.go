package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// User represente un client du service
type User struct {
	// ID est l'identifiant unique du user dans la BDD
	ID string `json:"id"`
	// FirstName est le prénom du user
	FirstName string `json:"first_name"`
	// LastName est le nom de famille du user
	LastName string `json:"last_name"`
	// Email est utilisé pour la connexion et la communication avec le user
	Email string `json:"email" gorm:"index:,unique"`
	// Password est haché en sha256 du user
	Password Password `json:"password"`
}

// UnmarshalJSON implémente une interface de contrat Unmarshaler sur le user
func (u User) MarshalJSON() ([]byte, error) {
	aux := struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}

	return json.Marshal(aux)
}

func hashPassword(value string) (res Password) {
	h := sha256.New()
	h.Write([]byte(value))
	res = Password(fmt.Sprintf("%x", h.Sum(nil)))
	return res
}

// Password is a type for containing password.
type Password string

// MarshalJSON implémente une interface de contrat Marshaler sur le password
func (p *Password) UnmarshalJSON(data []byte) error {

	var s string = ""
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*p = hashPassword(s)

	return nil
}

type LoginPayload struct {
	Email string `json:"email"`
	// Password est haché en sha256 du user
	Password Password `json:"password"`
}

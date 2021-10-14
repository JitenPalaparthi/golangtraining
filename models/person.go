package models

import (
	"example/models/common"
)

type Person struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Mobile       string `json:"mobile"`
	Status       string `json:"status"`
	LastModified string `json:"lastModified"`
}

func (p *Person) ValidatePerson() (err error) {
	if !common.IsEmailValid(p.Email) {
		return ErrInvalidEmail
	}
	return nil
}

// gob struct to []byte or string

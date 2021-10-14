package database

import (
	"example/models"
	"fmt"

	"gorm.io/gorm"
)

// CreatePerson(person)(result,err)
// UpdatePersonByID(id,data map)(result,err)

type PersonDB struct {
	DB *Database
}

func (p *PersonDB) CreatePerson(person *models.Person) (result string, err error) {
	p.DB.Client.(*gorm.DB).AutoMigrate(person)
	res := p.DB.Client.(*gorm.DB).Create(person)
	if res.Error != nil {
		return "", res.Error
	}
	return fmt.Sprintln(person.ID), nil
}

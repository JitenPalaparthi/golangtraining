package interfaces

import (
	"example/models"
)

type IPerson interface {
	CreatePerson(person *models.Person) (result string, err error)
}

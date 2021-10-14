package filedb

import (
	"example/models"
	"fmt"
	"io/ioutil"
)

type PersonFileDB struct{}

func (p *PersonFileDB) CreatePerson(person *models.Person) (result string, err error) {
	err = ioutil.WriteFile("person.txt", []byte(fmt.Sprintln(person)), 0755)
	if err != nil {
		return "", err
	}
	return fmt.Sprintln("person.txt"), nil
}

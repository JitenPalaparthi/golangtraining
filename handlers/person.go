package handlers

import (
	"encoding/json"
	"example/interfaces"
	"example/messaging"
	"example/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Interface interfaces.IPerson
}

func (p *Person) CreatePerson() func(*gin.Context) {
	var err error
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			person := &models.Person{}
			err = json.NewDecoder(c.Request.Body).Decode(person)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failure",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			err = person.ValidatePerson()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failure",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			result, err := p.Interface.CreatePerson(person)
			err = person.ValidatePerson()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failure",
					"message": err.Error(),
				})
				c.Abort()
				return
			}

			message := messaging.Message{}
			message.Subject = "person.create"
			buf, err := person.ToBytes()
			// This error will not delete the record from the db.This has to be handled at a later stage
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failure",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			message.Data = buf
			messaging.ChMessage <- message

			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": result,
			})
		}
	}
}

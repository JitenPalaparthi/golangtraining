package models

import (
	"bytes"
	"encoding/gob"
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

func (p *Person) ToBytes() ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	//	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes(), nil
}

// gob struct to []byte or string

package database

import (
	"fmt"
	"os"
	"time"

	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_MAX_TRIES    int = 10
	DB_MAX_DURATION int = 5
)

type Database struct {
	Client interface{}
	DBName string
}

func GetConnection(conStr, dbName string) (*Database, error) {
	var (
		db  *gorm.DB
		err error
	)

	if os.Getenv("DB_MAX_TRIES") != "" {
		DB_MAX_TRIES, err = strconv.Atoi(os.Getenv("DB_MAX_TRIES"))
		return nil, err
	}
	if os.Getenv("DB_MAX_DURATION") != "" {
		DB_MAX_DURATION, err = strconv.Atoi(os.Getenv("DB_MAX_DURATION"))
		return nil, err
	}

	db, err = gorm.Open(mysql.Open(conStr), &gorm.Config{})
	if err != nil {
		for i := 1; i < DB_MAX_TRIES; i++ {
			time.Sleep(time.Duration(DB_MAX_DURATION) * time.Second)
			fmt.Println("Trying to connect to the database-", i, " times")
			db, err = gorm.Open(mysql.Open(conStr), &gorm.Config{})
			if err == nil {
				return &Database{Client: db, DBName: dbName}, nil
			}
		}
		return nil, err
	}
	return &Database{Client: db, DBName: dbName}, nil
}

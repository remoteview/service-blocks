package common

import (
	"fmt"

	"github.com/gobuffalo/pop"
)

// Database -
type Database struct {
	*pop.Connection
}

// DB -
var DB *pop.Connection

// Init - Opening a database and save the reference to `Database` struct.
func Init() *pop.Connection {
	db, err := pop.Connect("development")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	DB = db
	return DB
}

// GetDB - Using this function to get a connection, you can create your connection pool here.
func GetDB() *pop.Connection {
	return DB
}

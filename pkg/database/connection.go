package database

import (
	"database/sql"
	"fmt"
)

type Connection struct {
	databaseURL      string
	databasePort     string
	databaseDriver   string
	databaseName     string
	databaseUser     string
	databasePassword string
}

func NewConnection(
	databaseURL string,
	databasePort string,
	databaseDriver string,
	databaseName string,
	databaseUser string,
	databasePassword string) Connection {
	return Connection{
		databaseURL:      databaseURL,
		databasePort:     databasePort,
		databaseDriver:   databaseDriver,
		databaseName:     databaseName,
		databaseUser:     databaseUser,
		databasePassword: databasePassword,
	}
}

func (c Connection) Open() (*sql.DB, error) {
	return sql.Open(c.databaseDriver, fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		c.databaseURL, c.databasePort, c.databaseName, c.databaseUser, c.databasePassword))
}

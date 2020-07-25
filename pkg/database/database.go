package database

import (
	"Photon_Server/pkg/config"
	"database/sql"
	"fmt"
	"net/url"
)

// DB Database
type DB struct {
	Host       string
	Port       string
	Username   string
	Password   string
	DBName     string
	Connection *sql.DB
}

// NewDB DataBase create
func NewDB(c *config.Config) *DB {
	return newDB(&DB{
		Host:     c.DataBase.Host,
		Port:     c.DataBase.Port,
		Username: c.DataBase.User,
		Password: c.DataBase.Password,
		DBName:   c.DataBase.Database,
	})
}

func newDB(d *DB) *DB {
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.DBName)

	option := url.Values{}
	option.Add("charset", "utf8")
	option.Add("parseTime", "True")
	option.Add("loc", "Local")

	connection := fmt.Sprintf("%s?%s", connectionInfo, option.Encode())

	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err.Error())
	}

	d.Connection = db
	return d
}

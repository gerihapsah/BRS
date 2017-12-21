package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var SQL *sqlx.DB

// MySQLInfo is the details for the database connection
type Info struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Conect(i Info) {
	var err error
	dsn := i.Username + ":" + i.Password + "@tcp(" + i.Hostname + ":" + fmt.Sprintf("%d", i.Port) + ")/" + i.Name
	if SQL, err = sqlx.Connect("mysql", dsn); err != nil {
		log.Println("SQL Driver Error", err)
	}
	// Check if is alive
	if err = SQL.Ping(); err != nil {
		log.Println("Database Error", err)
	}
}

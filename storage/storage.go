package storage

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MySQL"
	Postgres Driver = "Postgres"
)

// New create the connection with database
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(
		func() {
			var err error
			// db, err = gorm.Open("postgres", "postgres://openpg:openpgpwd@localhost:5432/godb?sslmode=disable")
			db, err = gorm.Open("postgres", "postgres://openpg:openpgpwd@localhost:5432/gorm?sslmode=disable")
			if err != nil {
				log.Fatalf("can't open database: %v", err)
			}
			// no vamos a cerrar la conexion porque es unica
			// defer db.Close()

			fmt.Println("database postgresql connected")
		},
	)
}

func newMySQLDB() {
	once.Do(
		func() {
			var err error
			db, err = gorm.Open("mysql", "root:leandro@tcp(localhost:3306)/godb")
			if err != nil {
				log.Fatalf("can't open database: %v", err)
			}
			// no vamos a cerrar la conexion porque es unica
			// defer db.Close()

			fmt.Println("database mysql connected")
		},
	)
}

// Pool return a unique instance od db
func DB() *gorm.DB {
	return db
}

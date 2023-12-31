package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/cart_system?sslmode=disable"
)

var testQueris *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueris = New(testDB)
	os.Exit(m.Run())
}

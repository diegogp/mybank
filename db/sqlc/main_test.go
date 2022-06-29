package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/mybank?sslmode=disable"
)

var testQueries *Queries

//entrypoint for all tests
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}

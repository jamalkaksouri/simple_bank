package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/stretchr/testify/require"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@127.0.0.1:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}

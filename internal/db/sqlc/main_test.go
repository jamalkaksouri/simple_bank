package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/jamalkaskouri/simple_bank/internal/db/util"
	_ "github.com/lib/pq"
	_ "github.com/stretchr/testify/require"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testDB, err = sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}

package main

import (
	"database/sql"
	"log"

	"github.com/jamalkaskouri/simple_bank/internal/api"
	db "github.com/jamalkaskouri/simple_bank/internal/db/sqlc"
	"github.com/jamalkaskouri/simple_bank/internal/db/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

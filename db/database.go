package db

import (
	"log"

	"github.com/go-pg/pg/v10"
)

func ConnectDB() *pg.DB {
	opt, err := pg.ParseURL("postgres://amendez:@localhost:5432/test?sslmode=disable")

	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	log.Println("Conexion Realizada")
	return db
}

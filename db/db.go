package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)


const (
	host = "localhost"
	port = 5432
	user = "kylemartinelli"
	dbname = "surflinedata"
)

func NewPostgresStorage () (*sql.DB, error){
	psqlconn := fmt.Sprintf("host= %s port= %d user= %s dbname= %s sslmode=disable", host, port, user, dbname)

db, err := sql.Open("postgres", psqlconn)
if err != nil {
	log.Fatal(err)
}
return db, nil

}


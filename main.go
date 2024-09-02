package main

import (
	"database/sql"
	"log"

	"github.com/kylemartinelli/surflineData/db"
	"github.com/kylemartinelli/surflineData/parser"
	"github.com/kylemartinelli/surflineData/requests"
	_ "github.com/lib/pq"
)


func main() {
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}
initStorage(db)

data, err := parser.ReadCsv("/Users/kylemartinelli/surflineData/surfHrefs2.csv")
if err != nil {
	log.Fatal(err, "error parsing csv")
}

requests.PingSurflineServices(data, db)


}

func initStorage(db *sql.DB)  {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Succesfully Connected")
}
package main

import (
	"log"

	"github.com/lylemartinelli/surflineData/parser"
	"github.com/lylemartinelli/surflineData/requests"
)



func main() {
data, err := parser.ReadCsv("/Users/kylemartinelli/surflineData/surfHrefs2.csv")
if err != nil {
	log.Fatal(err, "error parsing csv")
}

requests.PingSurflineServices(data)






}
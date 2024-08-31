package main

import (
	"fmt"
	"log"

	"github.com/lylemartinelli/surflineData/parser"
)



func main() {
data, err := parser.ReadCsv("/Users/kylemartinelli/surflineData/surfHrefs2.csv")
if err != nil {
	log.Fatal(err, "error parsing csv")
}

fmt.Println(data)






}
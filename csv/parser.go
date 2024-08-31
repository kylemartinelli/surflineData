package csv

import (
	"log"
	"os"
)

func ReadCsv() {
	data, err := os.ReadFile("surfHrefs.csv")
	if err != nil {
		log.Fatal(err)
	}
	println(data)

}
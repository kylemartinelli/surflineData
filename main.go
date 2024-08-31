package main

import (
	"encoding/csv"
	"fmt"
	"os"
)


func main() {

		file, err := os.Open("/Users/kylemartinelli/surflineData/surfHrefs2.csv")
		if err != nil {
			fmt.Println("Error opening file", err)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Error reading csv", err)
		}
		for i := 0 ;i < 10; i++{
			fmt.Println(records[i][1])
		}



}
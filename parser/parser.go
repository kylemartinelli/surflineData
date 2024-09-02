package parser

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsv(fileName string) ([][]string, error){



	file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file", err)
		}
		defer file.Close()

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			return records, error(err)
		}

		return records, nil

}
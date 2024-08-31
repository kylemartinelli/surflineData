package parser

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsv(fileName string) ([]string, error){
	var result []string
	file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file", err)
		}
		defer file.Close()

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			return result, error(err)
		}

		for i := 1 ;i < 5; i++{
			result = append(result, records[i][1])
		}
		return result, nil

}
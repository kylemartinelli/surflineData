package csvParser

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsv() []string{
	var result []string
	file, err := os.Open("/Users/kylemartinelli/surflineData/surfHrefs2.csv")
		if err != nil {
			fmt.Println("Error opening file", err)
		}
		defer file.Close()

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Error reading csv", err)
		}
		for i := 0 ;i < 10; i++{
			result = append(result, records[i][1])
		}
		return result

}
package main

import (
	"fmt"

	"github.com/lylemartinelli/surflineData/parser"
)



func main() {
data := parser.ReadCsv()
fmt.Print(data)




}
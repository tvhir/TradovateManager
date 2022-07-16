package main

import (
	"fmt"

	ETL "github.com/tvhir/TradovateManager/ETL"
)

func main() {
	fmt.Println("Tradovate Manager")

	//Extract data from CSV, recieve struct with data
	trades := ETL.Extract()

	for _, trade := range trades {
		fmt.Printf("%+v\n", *trade)
	}
}

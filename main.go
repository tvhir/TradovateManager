package main

import (
	"fmt"

	ETL "github.com/tvhir/TradovateManager/ETL"
)

func main() {
	fmt.Println("Tradovate Manager")

	//Extract data from CSV, recieve struct with data
	trades := ETL.Extract()

	// for _, trade := range trades {
	// 	fmt.Printf("%+v\n", *trade)
	// }

	transformedTrades := ETL.Transform(trades)

	for _, trade := range transformedTrades {
		fmt.Printf("%+v\n", *trade)
	}

	ETL.Load(trades)

}

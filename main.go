package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Tradovate Manager")

	//Open CSV
	performanceCSV, err := os.Open("Performance.07.11-15.22.csv")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Opened CSV")
	defer performanceCSV.Close()

	// read and print CSV file
	fileReader := csv.NewReader(performanceCSV)
	trades, error := fileReader.ReadAll()
	if error != nil {
		fmt.Println(error)
	}
	for trade := range trades {
		fmt.Printf("%s\n", trades[trade])
	}
}

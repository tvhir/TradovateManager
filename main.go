package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Trade struct {
	Symbol          string  `csv:"symbol"`
	PriceFormat     int     `csv:"_priceFormat"`
	PriceFormatType int     `csv:"_priceFormatType"`
	TickSize        float32 `csv:"_tickSize"`
	BuyFillId       int     `csv:"buyFillId"`
	SellFillId      int     `csv:"sellFillId"`
	Qty             int     `csv:"qty"`
	BuyPrice        float32 `csv:"buyPrice"`
	SellPrice       float32 `csv:"sellPrice"`
	PnL             string  `csv:"pnL"`
	BoughtTimestamp string  `csv:"boughtTimestamp"`
	SellTimestamp   string  `csv:"sellTimestamp"`
	Duration        string  `csv:"duration"`
}

func main() {
	fmt.Println("Tradovate Manager")

	//Open CSV
	performanceCSV, err := os.Open("Performance.07.11-15.22.csv")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Opened CSV")
	defer performanceCSV.Close()

	// // read and print CSV file
	// fileReader := csv.NewReader(performanceCSV)
	// trades, error := fileReader.ReadAll()
	// if error != nil {
	// 	fmt.Println(error)
	// }
	// for trade := range trades {
	// 	fmt.Printf("%s\n", trades[trade])
	// }

	trades := []*Trade{}

	if err := gocsv.UnmarshalFile(performanceCSV, &trades); err != nil {
		panic(err)
	}
	for _, trade := range trades {
		fmt.Printf("%+v\n", *trade)
	}
}

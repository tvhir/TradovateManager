package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type Trade struct {
	Symbol          string             `csv:"symbol"`
	TickSize        float32            `csv:"_tickSize"`
	Qty             int                `csv:"qty"`
	BuyPrice        float32            `csv:"buyPrice"`
	SellPrice       float32            `csv:"sellPrice"`
	PnL             string             `csv:"pnl"`
	BoughtTimestamp TradovateTimestamp `csv:"boughtTimestamp"`
	SoldTimestamp   TradovateTimestamp `csv:"soldTimestamp"`
	Duration        string             `csv:"duration"`
}

type TradovateTimestamp struct {
	time.Time
}

// Convert the CSV string as internal date
func (date *TradovateTimestamp) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("01/02/2006 15:04:05", csv)
	return err
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

	// read and print CSV file
	trades := []*Trade{}

	if err := gocsv.UnmarshalFile(performanceCSV, &trades); err != nil {
		panic(err)
	}
	for _, trade := range trades {
		fmt.Printf("%+v\n", *trade)
	}
}

package ETL

import (
	"fmt"
	"os"
	"strings"
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
	Duration        TradovateDuration  `csv:"duration"`
}

type TradovateTimestamp struct {
	time.Time
}

// Convert the CSV string as internal date
func (date *TradovateTimestamp) UnmarshalCSV(timestampString string) (err error) {
	date.Time, err = time.Parse("01/02/2006 15:04:05", timestampString)
	return err
}

type TradovateDuration struct {
	time.Duration
}

// Convert the CSV string as internal date
func (duration *TradovateDuration) UnmarshalCSV(durationString string) (err error) {
	durationArray := strings.Fields(durationString)
	for i, duration := range durationArray {
		if strings.Contains(duration, "min") {
			durationArray[i] = strings.Replace(duration, "min", "m", 1)
		}
		if strings.Contains(duration, "sec") {
			durationArray[i] = strings.Replace(duration, "sec", "s", 1)
		}
	}
	duration.Duration, err = time.ParseDuration(strings.Join(durationArray, ""))
	if err != nil {
		panic(err)
	}
	return err
}

//Function that will extract data from a Tradovate Performance CSV and insert into a struct
func Extract() []*Trade {

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

	return trades
}

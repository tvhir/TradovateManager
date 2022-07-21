package ETL

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

type Trade struct {
	Symbol          string             `csv:"symbol"`
	Direction       string             `csv:"direction"`
	TickSize        float64            `csv:"_tickSize"`
	Quantity        float64            `csv:"qty"`
	BuyPrice        float64            `csv:"buyPrice"`
	SellPrice       float64            `csv:"sellPrice"`
	PriceSpread     float64            `csv:"priceSpread"`
	GrossPnL        Currency           `csv:"pnl"`
	Commisions      float64            `csv:"commisions"`
	NetPnL          float64            `csv:"netPNL"`
	RunningNetPNL   float64            `csv:"runningNetPNL"`
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
	return err
}

type Currency struct {
	float64
}

// Convert the CSV string as internal date
func (dollars *Currency) UnmarshalCSV(dollarString string) (err error) {
	//Remove the $ char
	dollarString = dollarString[1:]

	//Negative Value
	if dollarString[0] == '(' && dollarString[len(dollarString)-1] == ')' {
		dollarString = "-" + dollarString[1:len(dollarString)-1]
	}

	//Remove any commas
	dollarString = strings.ReplaceAll(dollarString, ",", "")

	dollars.float64, err = strconv.ParseFloat(dollarString, 64)
	return err
}

//Function that will extract data from a Tradovate Performance CSV and insert into a struct
func Extract() []*Trade {

	//Open CSV
	performanceCSV, err := os.Open("performance.csv")
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

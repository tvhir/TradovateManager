package ETL

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func Load(trades []*Trade) {

	csvFile, err := os.Create("detailedPerformance.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()
	gocsv.MarshalFile(&trades, csvFile)
}

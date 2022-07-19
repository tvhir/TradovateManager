package ETL

import "fmt"

func Transform(trades []*Trade) {
	runningNet := 0.0
	for _, trade := range trades {
		//Direction
		if trade.BoughtTimestamp.Time.Before(trade.SoldTimestamp.Time) {
			trade.Direction = "Long"
		}
		if trade.BoughtTimestamp.Time.After(trade.SoldTimestamp.Time) {
			trade.Direction = "Short"
		}

		//Price Spread
		trade.PriceSpread = trade.SellPrice - trade.BuyPrice

		//Commissions and NetPNL
		if trade.Symbol[:1] == "M" {
			trade.Commisions = trade.Quantity * -1.32
			trade.NetPnL = trade.GrossPnL.float64 - trade.Commisions
		}
		if trade.Symbol[:1] == "N" {
			trade.Commisions = trade.Quantity * -4.39
			trade.NetPnL = trade.GrossPnL.float64 - trade.Commisions
		}
		runningNet += trade.NetPnL
		trade.RunningNetPNL = runningNet
		fmt.Printf("%+v\n", *trade)
	}
}

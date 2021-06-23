package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var listPrice []float64
var Price_buy, Qty_buy_usd, TotalUnitBitcoin, Expected_profit, Expected_to_sell, Sellprice float64
var Presentage_profit, RateUSDIDR, running_price_bitcoin_usd, Expected_profit_max_sell, SellNow float64
var IsProfit bool

func main() {

	var url string = "https://gist.githubusercontent.com/Jekiwijaya/c72c2de532203965bf818e5a4e5e43e3/raw/2631344d08b044a4b833caeab8a42486b87cc19a/gistfile1.txt"
	data, _ := getPriceIDRMarket(url)
	dataMarket := strings.Split(data, " ")

	RateUSDIDR = 14465.15                        //USD to ID
	Price_buy = 33936.84                         // Price 1 BIT COIN TO USD
	Qty_buy_usd = 1                              // Qty($) Buy
	TotalUnitBitcoin = (Qty_buy_usd / Price_buy) //Total Unit Buy

	Expected_profit = 6 // $ Profit expected
	Presentage_profit = (Qty_buy_usd + Expected_profit) / Qty_buy_usd
	Expected_to_sell = Presentage_profit * TotalUnitBitcoin //Price for release

	fmt.Println("Begining Unit = ", TotalUnitBitcoin)
	fmt.Println(Expected_to_sell)

	listPriceMarket := ListConvertFloatValue(dataMarket)

	for i := 0; i < len(listPriceMarket); i++ {
		running_price_bitcoin_usd = listPriceMarket[i] / RateUSDIDR //Conversi Bitcoin IDR to Bitcoin USD
		IsProfit = IsProfitExpected(TotalUnitBitcoin, Expected_to_sell, running_price_bitcoin_usd)
		if IsProfit {
			Sellprice = running_price_bitcoin_usd
			totalsell := Sellprice * TotalUnitBitcoin
			//Do Sell now
			fmt.Printf("Hours (%v) = (%v) Sell Now (%v X %v = %v)", (i + 1), running_price_bitcoin_usd, TotalUnitBitcoin, Sellprice, totalsell)
		} else {
			fmt.Printf("Hours (%v) = (%v)", (i + 1), running_price_bitcoin_usd)
		}

		fmt.Println()
	}
}

/*
	SellProfit is
*/
func IsProfitExpected(TUnitBitcoin float64, profit_expected float64, current_price float64) bool {
	unit_runing := current_price * TUnitBitcoin
	//fmt.Println("Unit Run  = ", unit_runing)
	if unit_runing >= profit_expected {
		return true
	}
	return false
}

/*
	ListConvertFloatValue is
*/

func ListConvertFloatValue(dMarket []string) []float64 {
	var newlistPrice []float64
	for i := 0; i < len(dMarket); i++ {
		if n, err := strconv.ParseFloat(dMarket[i], 64); err == nil {
			newlistPrice = append(newlistPrice, n)
		}

	}
	return newlistPrice
}

/*
	getPriceIDRMarket is
*/

func getPriceIDRMarket(url string) (string, error) {
	//return "3 2 1 5 6 2", nil
	resp, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil

}

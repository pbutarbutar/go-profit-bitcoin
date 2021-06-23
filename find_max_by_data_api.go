package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var listPrice []float64

func main() {

	var url string = "https://gist.githubusercontent.com/Jekiwijaya/c72c2de532203965bf818e5a4e5e43e3/raw/2631344d08b044a4b833caeab8a42486b87cc19a/gistfile1.txt"
	data, _ := getPriceIDRMarket(url)
	dataMarket := strings.Split(data, " ")

	listPriceMarket := ListConvertFloatValue(dataMarket)

	min, max := findMinAndMax(listPriceMarket)
	fmt.Printf("Min: %.4f", min)
	fmt.Println()
	fmt.Printf("Max: %.4f", max)
}

func findMinAndMax(a []float64) (min float64, max float64) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
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

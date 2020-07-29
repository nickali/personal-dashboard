package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const url = "https://www.alphavantage.co/query"

// StockQuote is wrapper for quote data.
type StockQuote struct {
	QuoteDetail *QuoteDetail `json:"Global Quote"`
}

// QuoteDetail provides details about stock.
type QuoteDetail struct {
	QuoteSymbol    string `json:"01. symbol"`
	QuoteOpen      string `json:"02. open"`
	QuoteHigh      string `json:"03. high"`
	QuoteLow       string `json:"04. low"`
	QuotePrice     string `json:"05. price"`
	QuoteVol       string `json:"06. volume"`
	QuoteLastDate  string `json:"07. latest trading day"`
	QuotePrevClose string `json:"08. previous close"`
	QuoteChange    string `json:"09. change"`
	QuoteChangePer string `json:"10. change percent"`
}

// StockPrint just outputs a string.
// See https://www.alphavantage.co/documentation/#latestprice
func StockPrint(stSymbol string, stAPI string) {
	stURL := url + "?function=GLOBAL_QUOTE&symbol=" + stSymbol + "&apikey=" + stAPI
	//	fmt.Println("Stock URL: %s", stURL)
	response, err := http.Get(stURL)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		jsonData := &StockQuote{
			QuoteDetail: &QuoteDetail{},
		}
		err := json.Unmarshal([]byte(data), jsonData)

		if err != nil {
			fmt.Printf("Something failed with failed with error %s\n", err)
		} else {
			fQuote, _ := strconv.ParseFloat(jsonData.QuoteDetail.QuotePrice, 64)
			fmt.Println(jsonData.QuoteDetail.QuoteSymbol + ":")
			fmt.Print("\t")
			fmt.Println(fQuote)
		}
	}
	return
}

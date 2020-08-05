package stocks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
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

// Print just outputs a string.
// See https://www.alphavantage.co/documentation/#latestprice
func Print(stSymbol string, stAPI string) ([]string, []text.WriteOption, []text.WriteOption) {
	stURL := url + "?function=GLOBAL_QUOTE&symbol=" + stSymbol + "&apikey=" + stAPI
	wrappedText := make([]string, 0)
	wrappedOpt := make([]text.WriteOption, 0)
	wrappedState := make([]text.WriteOption, 0)

	response, err := http.Get(stURL)
	if err != nil {
		wrappedText = append(wrappedText, "The HTTP request failed with error with stock")
		wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
		wrappedState = append(wrappedState, text.WriteReplace())
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		jsonData := &StockQuote{
			QuoteDetail: &QuoteDetail{},
		}
		err := json.Unmarshal([]byte(data), jsonData)

		if err != nil {
			wrappedText = append(wrappedText, "Problems unmarshalling stock")
			wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
			wrappedState = append(wrappedState, text.WriteReplace())

		} else {
			wrappedText = append(wrappedText, jsonData.QuoteDetail.QuoteSymbol+": "+jsonData.QuoteDetail.QuotePrice)
			wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorGreen)))
			wrappedState = append(wrappedState, text.WriteReplace())
		}
	}
	return wrappedText, wrappedOpt, wrappedState
}

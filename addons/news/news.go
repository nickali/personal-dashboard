package news

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/tidwall/gjson"
)

var maxItems = 20
var maxLength = 300
var fmtDate string

// News is wrapper for news data.

// NewsPrint just outputs a string.

func NewsPrint(url string, stAPI string) ([]string, []text.WriteOption, []text.WriteOption) {
	wrappedText := make([]string, 0)
	wrappedOpt := make([]text.WriteOption, 0)
	wrappedState := make([]text.WriteOption, 0)
	dt := time.Now()

	stURL := url + stAPI

	response, err := http.Get(stURL)
	if err != nil {
		panic(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		var jsonData = gjson.Parse(string(data)).Value().(map[string]interface{})

		if jsonData["status"] != "ok" {
			panic(err)
		} else {
			wrappedText = append(wrappedText, dt.Format("Updated (01-02-2006 15:04:05)")+"("+jsonData["status"].(string)+")\n\n")
			wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorYellow)))
			wrappedState = append(wrappedState, text.WriteReplace())

			numArticles, _ := strconv.Atoi(gjson.Get(string(data), "totalResults").Array()[0].String())
			resultTitle := gjson.Get(string(data), "articles.#.title").Array()
			resultDesc := gjson.Get(string(data), "articles.#.description").Array()
			resultURL := gjson.Get(string(data), "articles.#.url").Array()

			for i := 0; i < numArticles/2-1; i++ {

				descLength := len(resultDesc[i].String())
				if descLength < 300 {
					if descLength == 0 {
						maxLength = 0
					} else {
						maxLength = len(resultDesc[i].String()) - 1
					}
				} else {
					maxLength = 300
				}

				scrubbedDesc := strings.ReplaceAll(resultDesc[i].String(), "\u00a0", "")
				wrappedText = append(wrappedText, resultTitle[i].String()+"\n")
				wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorGreen)))
				wrappedState = append(wrappedState, nil)

				wrappedText = append(wrappedText, scrubbedDesc[:maxLength]+"\n")
				wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
				wrappedState = append(wrappedState, nil)

				wrappedText = append(wrappedText, resultURL[i].String()+"\n\n")
				wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorBlue)))
				wrappedState = append(wrappedState, nil)
			}
		}
	}

	return wrappedText, wrappedOpt, wrappedState

}

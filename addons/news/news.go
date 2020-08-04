package news

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	//	"util/textTransform"

	"github.com/tidwall/gjson"
)

var maxItems = 20
var maxLength = 100
var fmtDate string

// News is wrapper for news data.

// NewsPrint just outputs a string.

func NewsPrint(url string, stAPI string) string {
	dt := time.Now()
	var b strings.Builder

	stURL := url + stAPI

	response, err := http.Get(stURL)
	if err != nil {
		//		fmt.Printf("The HTTP request failed with error %s\n", err)
		fmt.Fprintf(&b, "The HTTP request failed with error %s %s\n", stURL, err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		//var jsonData map[string]interface{}

		var jsonData = gjson.Parse(string(data)).Value().(map[string]interface{})

		//err := json.Unmarshal([]byte(data), &jsonData)

		if jsonData["status"] != "ok" {
			fmt.Printf("Something failed with failed with error %s\n", err)
			fmt.Fprintf(&b, "Something failed with failed with error %s\n", err)
			panic(err)
		} else {
			//			fmt.Println("Weather:")
			//			fmt.Println("\t" + fmt.Sprintf("%2.0f", jsonData.WeatherDetail.WeatherTemp) + "F")
			fmt.Fprintf(&b, "Updated (%s) (%s): \n\n", dt.Format("01-02-2006 15:04:05"), jsonData["status"])

			numArticles, _ := strconv.Atoi(gjson.Get(string(data), "totalResults").Array()[0].String())
			resultTitle := gjson.Get(string(data), "articles.#.title").Array()
			resultDesc := gjson.Get(string(data), "articles.#.description").Array()
			resultURL := gjson.Get(string(data), "articles.#.url").Array()

			for i := 0; i < numArticles/2-1; i++ {

				descLength := len(resultDesc[i].String())
				if descLength < 100 {
					if descLength == 0 {
						maxLength = 0
					} else {
						maxLength = len(resultDesc[i].String()) - 1
					}
				} else {
					maxLength = 100
				}
				scrubbedDesc := strings.ReplaceAll(resultDesc[i].String(), "\u00a0", "")

				fmt.Fprintf(&b, "%s:\n%s\n%s\n\n", resultTitle[i], scrubbedDesc[:maxLength], resultURL[i].String())
			}
		}
	}

	return b.String()

}

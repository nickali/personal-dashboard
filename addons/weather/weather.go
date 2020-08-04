package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
)

const url = "https://api.openweathermap.org/data/2.5/weather?"

// Weather is wrapper for weather data.
type Weather struct {
	WeatherDetail *WeatherDetail `json:"main"`
}

// WeatherDetail provides details about the current weather
type WeatherDetail struct {
	WeatherTemp      float32 `json:"temp"`
	WeatherFeelsLike float32 `json:"feels_like"`
	WeatherTempMin   float32 `json:"temp_min"`
	WeatherTempMax   float32 `json:"temp_max"`
	WeatherPressure  float32 `json:"pressure"`
	WeatherHumidity  float32 `json:"humidity"`
}

// WeatherPrint just outputs a string.
// See https://openweathermap.org/current#one
func WeatherPrint(stZip string, stAPI string) ([]string, []text.WriteOption, []text.WriteOption) {
	dt := time.Now()
	var b strings.Builder
	wrappedText := make([]string, 0)
	wrappedOpt := make([]text.WriteOption, 0)
	wrappedState := make([]text.WriteOption, 0)

	stURL := url + "zip=" + stZip + "&appid=" + stAPI + "&units=imperial"

	response, err := http.Get(stURL)
	if err != nil {

		fmt.Fprintf(&b, "The HTTP request failed with error %s\n", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)
		jsonData := &Weather{
			WeatherDetail: &WeatherDetail{},
		}
		err := json.Unmarshal([]byte(data), jsonData)

		if err != nil {
			panic(err)
		} else {
			wrappedText = append(wrappedText, "Weather ("+dt.Format("01-02-2006 15:04:05")+"): ")
			wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorBlue)))
			wrappedState = append(wrappedState, text.WriteReplace())

			wrappedText = append(wrappedText, fmt.Sprintf("%2.0f", jsonData.WeatherDetail.WeatherTemp)+" F")
			wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
			wrappedState = append(wrappedState, nil)
		}
	}
	return wrappedText, wrappedOpt, wrappedState
}

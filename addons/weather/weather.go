package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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
func WeatherPrint(stZip string, stAPI string) string {
	dt := time.Now()
	var b strings.Builder

	stURL := url + "zip=" + stZip + "&appid=" + stAPI + "&units=imperial"
	//fmt.Println("Weather URL: %s", stURL)
	response, err := http.Get(stURL)
	if err != nil {
		//		fmt.Printf("The HTTP request failed with error %s\n", err)
		fmt.Fprintf(&b, "The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		jsonData := &Weather{
			WeatherDetail: &WeatherDetail{},
		}
		err := json.Unmarshal([]byte(data), jsonData)

		if err != nil {
			fmt.Printf("Something failed with failed with error %s\n", err)
			fmt.Fprintf(&b, "Something failed with failed with error %s\n", err)
		} else {
			//			fmt.Println("Weather:")
			//			fmt.Println("\t" + fmt.Sprintf("%2.0f", jsonData.WeatherDetail.WeatherTemp) + "F")
			fmt.Fprintf(&b, "Weather (%s): ", dt.Format("01-02-2006 15:04:05"))
			fmt.Fprintf(&b, fmt.Sprintf("%2.0f", jsonData.WeatherDetail.WeatherTemp)+"F")
		}
	}

	//	b.WriteString("ignition")
	//	fmt.Println(b.String())

	return b.String()
}

package main

import (
	"addons/stocks"
	"addons/weather"
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

func readConfig() bool {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("can't find config file: %s", err)
		} else {
			log.Fatalf("something wrong with config file: %s", err)
		}
	}

	/*
		// Print out config file content
		c := viper.AllSettings()
		bs, err := yaml.Marshal(c)
		if err != nil {
			log.Fatalf("unable to marshal config to YAML: %v", err)
		}
		fmt.Println("Printing imported config ---")
		fmt.Println(string(bs))
		fmt.Println("Done printing imported config ---")
	*/

	return true
}

func main() {
	fmt.Println("Starting the application...")

	/*
		var ip = flag.Int("flagname", 1234, "help message for flagname")
		flag.Parse()
		fmt.Println("ip has value ", *ip)
	*/

	readConfig()

	// Call and wait till all are finished.
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		weather.WeatherPrint(viper.GetString("weather.zip"), viper.GetString("weather.api_key"))
		wg.Done()
	}()

	go func() {
		stocks.StockPrint(viper.GetString("stocks.symbol"), viper.GetString("stocks.api_key"))
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Terminating the application...")
}

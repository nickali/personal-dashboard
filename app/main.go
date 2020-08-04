package main

import (
	"addons/news"
	"addons/newsreader"
	"addons/stocks"
	"addons/weather"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

const rootID = "root"

func printConfig() {
	// Print out config file content
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	fmt.Println("Printing imported config ---")
	fmt.Println(string(bs))
	fmt.Println("Done printing imported config ---")

}

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

	return true
}

func writeUpdate(ctx context.Context, TopLeft *text.Text, BottomLeft *text.Text, TopRight *text.Text, delay time.Duration) {

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			weatherOutput := weather.WeatherPrint(viper.GetString("weather.zip"), viper.GetString("weather.api_key"))
			if err := TopLeft.Write(fmt.Sprintf("%s\n", weatherOutput), text.WriteReplace()); err != nil {
				panic(err)
			}
			newsreaderOutput := newsreader.NewsReaderPrint(viper.GetString("newsreader.url"))
			if err := BottomLeft.Write(fmt.Sprintf("%s\n", newsreaderOutput), text.WriteReplace()); err != nil {
				panic(err)
			}
			newsOutput := news.NewsPrint(viper.GetString("news.url"), viper.GetString("news.api_key"))
			if err := TopRight.Write(fmt.Sprintf("%s\n", newsOutput), text.WriteReplace()); err != nil {
				panic(err)
			}

		case <-ctx.Done():
			return
		}
	}
}

func main() {
	fmt.Println("Starting the application...")

	/*
		var ip = flag.Int("flagname", 1234, "help message for flagname")
		flag.Parse()
		fmt.Println("ip has value ", *ip)
	*/

	readConfig()
	//printConfig()

	var weatherOutput string
	var newsreaderOutput string
	var newsOutput string

	//	var stocksOutput string

	// Call and wait till all are finished.
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		weatherOutput = weather.WeatherPrint(viper.GetString("weather.zip"), viper.GetString("weather.api_key"))
		wg.Done()
	}()

	go func() {
		stocks.StockPrint(viper.GetString("stocks.symbol"), viper.GetString("stocks.api_key"))
		wg.Done()
	}()

	go func() {
		newsreaderOutput = newsreader.NewsReaderPrint(viper.GetString("newsreader.url"))
		wg.Done()
	}()

	go func() {
		newsOutput = news.NewsPrint(viper.GetString("news.url"), viper.GetString("news.api_key"))
		wg.Done()
	}()

	wg.Wait()

	t, err := termbox.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())

	borderlessTopLeft, err := text.New()
	if err != nil {
		panic(err)
	}
	if err := borderlessTopLeft.Write(weatherOutput, text.WriteReplace()); err != nil {
		panic(err)
	}

	borderlessBottomLeft, err := text.New()
	if err != nil {
		panic(err)
	}
	if err := borderlessBottomLeft.Write(newsreaderOutput, text.WriteReplace()); err != nil {
		panic(err)
	}

	borderlessTopRight, err := text.New()
	if err != nil {
		panic(err)
	}
	if err := borderlessTopRight.Write(newsOutput, text.WriteReplace()); err != nil {
		panic(err)
	}

	go writeUpdate(ctx, borderlessTopLeft, borderlessBottomLeft, borderlessTopRight, 300*time.Second)

	c, err := container.New(
		t, container.ID(rootID),
		container.Border(linestyle.Light),
		container.BorderTitle("PRESS Q TO QUIT"),
		container.SplitVertical(
			container.Left(
				container.SplitHorizontal(container.Top(container.PlaceWidget(borderlessTopLeft)), container.Bottom(container.PlaceWidget(borderlessBottomLeft)), container.SplitPercent(5))),
			container.Right(container.PlaceWidget(borderlessTopRight))))

	if err != nil {
		panic(err)
	}

	if err := c.Update(rootID); err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter)); err != nil {
		panic(err)
	}

	fmt.Println("Terminating the application...")
}

package newsreader

import (
	"github.com/mmcdole/gofeed"
)

// NewsReaderPrint just outputs a string.
func NewsReaderPrint(stURL string) string {

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(stURL)

	if err != nil {
		return string("The HTTP request failed with error with feed")
	} else {
		return feed.Title
	}
}

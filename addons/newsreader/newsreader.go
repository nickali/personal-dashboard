package newsreader

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
)

var maxItems = 100
var stOutput strings.Builder

// NewsReaderPrint just outputs a string.
func NewsReaderPrint(stURL string) string {

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(stURL)

	if err != nil {
		stOutput.WriteString("The HTTP request failed with error with feed")
	} else {
		stOutput.WriteString(feed.Title + "\n\n")

		var items = feed.Items
		var maxItemsInFeed = len(items)
		p := bluemonday.StrictPolicy()
		p.AllowElements("p")

		for i := 0; i <= (maxItemsInFeed-1) && i <= (maxItems-1); i++ {
			stOutput.WriteString(items[i].Title + "\n")
			//strippedHTMLDesc := strip.StripTags(items[i].Description[:100])

			strippedNewlines2Desc := strings.Replace(items[i].Description, "\r", " ", -1)
			strippedNewlines3Desc := strings.Replace(strippedNewlines2Desc, "<!-- SC_OFF --><div>", "", -1)
			strippedNewlines4Desc := strings.Replace(strippedNewlines3Desc, "<div>", "", -1)
			strippedNewlines5Desc := strings.Replace(strippedNewlines4Desc, "</div>", "", -1)
			strippedNewlines6Desc := strings.Replace(strippedNewlines5Desc, "<div>", "", -1)
			strippedNewlines7Desc := strings.Replace(strippedNewlines6Desc, "<p>", "", -1)
			strippedNewlines8Desc := strings.Replace(strippedNewlines7Desc, "</p>", "", -1)
			strippedNewlines9Desc := strings.Replace(strippedNewlines8Desc, "&lt;", "", -1)
			strippedDesc := strippedNewlines9Desc[:100]

			stOutput.WriteString(p.Sanitize(strippedDesc) + "\n" + items[i].Link + "\n-----\n")
		}

	}

	return stOutput.String()

}

package newsreader

import (
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
)

var maxItems = 100
var maxLength = 300
var fmtDate string

// Print outputs what needs to be drawn on the screen.
func Print(stURL string) ([]string, []text.WriteOption, []text.WriteOption) {
	wrappedText := make([]string, 0)
	wrappedOpt := make([]text.WriteOption, 0)
	wrappedState := make([]text.WriteOption, 0)

	stOutput := strings.Builder{}
	dt := time.Now()
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(stURL)

	if err != nil {
		wrappedText = append(wrappedText, "The HTTP request failed with error with feed")
		wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
		wrappedState = append(wrappedState, text.WriteReplace())

	} else {
		fmtDate = "Updated (" + dt.Format("01-02-2006 15:04:05"+")\n\n")
		wrappedText = append(wrappedText, fmtDate)
		wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorYellow)))
		wrappedState = append(wrappedState, text.WriteReplace())
		stOutput.WriteString(fmtDate)

		var items = feed.Items
		var maxItemsInFeed = len(items)
		p := bluemonday.StrictPolicy()
		p.AllowElements("p")

		for i := 0; i <= (maxItemsInFeed-1) && i <= (maxItems-1); i++ {
			// sanitize description
			strippedNewlines2Desc := strings.Replace(items[i].Description, "\r", " ", -1)
			strippedNewlines3Desc := strings.Replace(strippedNewlines2Desc, "<!-- SC_OFF --><div>", "", -1)
			strippedNewlines4Desc := strings.Replace(strippedNewlines3Desc, "<div>", "", -1)
			strippedNewlines5Desc := strings.Replace(strippedNewlines4Desc, "</div>", "", -1)
			strippedNewlines6Desc := strings.Replace(strippedNewlines5Desc, "<div>", "", -1)
			strippedNewlines7Desc := strings.Replace(strippedNewlines6Desc, "<p>", "", -1)
			strippedNewlines8Desc := strings.Replace(strippedNewlines7Desc, "</p>", "", -1)
			strippedNewlines9Desc := strings.Replace(strippedNewlines8Desc, "&lt;", "", -1)
			strippedNewlines10Desc := strings.ReplaceAll(strippedNewlines9Desc, "\u200b", "")

			descLength := len(strippedNewlines10Desc)
			if descLength < 300 {
				if descLength == 0 {
					maxLength = 0
				} else {
					maxLength = len(strippedNewlines9Desc) - 1
				}
			} else {
				maxLength = 300
			}

			strippedDesc := p.Sanitize(strippedNewlines10Desc[:maxLength])

			wrappedText = append(wrappedText, items[i].Title+"\n")
			wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorGreen)))
			wrappedState = append(wrappedState, nil)

			if len(strippedDesc) > 5 {
				wrappedText = append(wrappedText, strippedDesc+"\n")
			} else {
				wrappedText = append(wrappedText, "")
			}

			wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
			wrappedState = append(wrappedState, nil)

			wrappedText = append(wrappedText, items[i].Link+"\n\n")
			wrappedOpt = append(wrappedOpt, text.WriteCellOpts(cell.FgColor(cell.ColorBlue)))
			wrappedState = append(wrappedState, nil)

		}

	}

	return wrappedText, wrappedOpt, wrappedState

}

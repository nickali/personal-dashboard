package util

import (
	"fmt"
	"strings"
)

func UseString(s []string) []string {
	return s
}

func SanitizeNonAlphaNum(s string) string {
	var b strings.Builder

	var bad []string

	bad = append(bad, "\u00a0")
	//bad = UseString(bad)

	for _, badChar := range bad {
		tString := b.String()
		b.Reset()
		b.WriteString(strings.ReplaceAll(tString, badChar, ""))
	}

	fmt.Fprintf(&b, s)

	return b.String()

}

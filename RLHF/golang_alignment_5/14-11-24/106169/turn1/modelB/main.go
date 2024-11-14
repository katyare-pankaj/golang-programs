package main

import (
	"regexp"
)

var validChars = regexp.MustCompile("[^a-zA-Z ]")

func scrubName(name string) string {
	return validChars.ReplaceAllString(name, "")
}

func main() {
	testing.Run("Example", func(t *testing.T) {
		t.Run("ScrubNames", testScrubNames)
	})
}

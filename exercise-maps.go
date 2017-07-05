package main

import (
  "strings"
	"golang.org/x/tour/wc"
)

// returns map which includes each word in `s` is occured.
func WordCount(s string) map[string]int {
	counts := make(map[string]int)
  for _, v := range strings.Fields(s) {
    counts[v] += 1
  }

  return counts
}

func main() {
	wc.Test(WordCount)
}

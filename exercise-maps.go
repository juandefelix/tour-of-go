package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  wordArray := strings.Fields(s)
  wordCountMap := make(map[string]int)

  
  for _, value := range wordArray {
    wordCountMap[value] += 1
  }


  return wordCountMap
}

func main() {
  wc.Test(WordCount)
}

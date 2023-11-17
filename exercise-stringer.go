package main

import (
  "fmt"
  "strings"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
  result := make([]string, len(i))
  
  for index, item := range i  {
    result[index] = fmt.Sprint(item)
  }
  return strings.Join(result, ".")
}

func main() {
  hosts := map[string]IPAddr{
    "loopback":  {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }
  for name, ip := range hosts {
    fmt.Printf("!!! %v: %v\n", name, ip)
  }
}

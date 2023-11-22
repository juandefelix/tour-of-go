package main

import "golang.org/x/tour/reader"


type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
  value := []byte("A")[0]
  
  for index, _ := range b {
    b[index] = value
  }
  return len(b), nil
}

func main() {
  reader.Validate(MyReader{})
}

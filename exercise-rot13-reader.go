package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (int, error) {
	length, error := reader.r.Read(b)

	if error == io.EOF {
		return 0, io.EOF
	} else {
		for index, item := range b { b[index] = rot13(item) }
		return length, nil
	}
}

func rot13(b byte) byte {
  var offset byte

  if 'a' <= b && b <= 'z' {
  	offset = 'a'
  } else if 'A' <= b && b <= 'Z' {
  	offset = 'A'
  } else {
  	return b
  }

  byteInAlphabet := b - offset
  rotatedByte := byteInAlphabet + 13
  adjustedRotatedByte := rotatedByte%26
  adjustedRotatedByteWithOffset := adjustedRotatedByte + offset

  return adjustedRotatedByteWithOffset
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, r)
}


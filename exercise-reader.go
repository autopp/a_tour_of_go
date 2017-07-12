package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// MyReader.Read generate infinite 'A'
func (MyReader) Read(b []byte) (int, error) {
  b[0] = 'A'
  return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(buf []byte) (int, error) {
  n, err := reader.r.Read(buf)

  if err != nil {
    return 0, err
  }

  for i := 0; i < n; i++ {
    b := buf[i]
    if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
      buf[i] = b + 13
    } else if (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z') {
      buf[i] = b - 13
    }
  }

  return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

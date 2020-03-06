package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {

	if len(b) == 0 {
		return 0, errors.New("Can't write on empty byte slice")
	}

	n, err := rot.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i, v := range b {
		switch {
		case v >= 'A' && v <= 'Z':
			if v < 'N' {
				b[i] = v + 13
			} else {
				b[i] = v - 13
			}
		case v >= 'a' && v <= 'z':
			if v < 'n' {
				b[i] = v + 13
			} else {
				b[i] = v - 13
			}
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

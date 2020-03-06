package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct {
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(p []byte) (int, error) {

	if len(p) == 0 {
		//You can return error here in place of nil as len(p) is 0 and writting 'A' is not not possible in empty slice
		return 0, nil
	}
	i := 0
	for ; i < len(p); i = i + 1 {
		p[i] = 'A'
	}
	return i, nil
}

func main() {
	reader.Validate(MyReader{})
}

package main

import "golang.org/x/tour/reader"

type myReader struct {
}

func (r *myReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(&myReader{})
}

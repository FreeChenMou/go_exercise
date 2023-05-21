package main

import (
	"io"
)

type Reader struct {
	s string
	n int
}

func (r *Reader) Read(p []byte) (int, error) {
	i := copy(p, r.s)
	r.n = i
	return i, nil
}

func NewReader(s string) io.Reader {
	return &Reader{s, 0}
}

//exercise7.4
func main() {
	reader := NewReader("wo ai nicccc ccccc")

	b := make([]byte, 10)

	reader.Read(b)
	println(string(b))
}

package main

import (
	"io"
	"os"
)

type LimitR struct {
	R io.Reader
	N int64
}

func (l *LimitR) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > l.N {
		p = p[:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitR{r, n}
}

func main() {
	file, _ := os.Open(os.Args[1])

	reader := LimitReader(file, 1024)
	b := make([]byte, 1024)
	reader.Read(b)
	println(string(b))
}

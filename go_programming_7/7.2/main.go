package main

import (
	"fmt"
	"io"
	"os"
)

var count int64

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	return w, &count
}

//exercise7.2
func main() {

	writer, i := CountingWriter(os.Stdout)
	n, err := writer.Write([]byte("wo zhen cnm\n"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write file err:%v", err)
	}
	*i = int64(n)
	println(*i)
}

package main

import (
	"io"
	"log"
	"net"
	"os"
)

//exercise8.12
func main() {
	addr := &net.TCPAddr{Port: 62482}
	dialer := net.Dialer{LocalAddr: addr}
	conn, err := dialer.Dial("tcp", "localhost:8000")
	//conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

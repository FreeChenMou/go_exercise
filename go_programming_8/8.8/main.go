package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//!-
// 练习 8.8： 使用select来改造8.3节中的echo服务器，为其增加超时，
func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	scanner := bufio.NewScanner(c)
	inputChan := make(chan string)
	go func() {
		ticker := time.NewTicker(1 * time.Second)

		for countdown := 10; countdown > 0; countdown-- {
			fmt.Printf("countdown: %d \n", countdown)
			select {
			case <-ticker.C:
				// do nothing
			case str := <-inputChan:
				countdown = 10
				go echo(c, str, 1*time.Second)
			}
		}
		ticker.Stop()
		c.Close()
	}()

	for scanner.Scan() {
		inputChan <- scanner.Text()
	}
}

//!+
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

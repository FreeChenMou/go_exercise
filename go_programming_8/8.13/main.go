// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

var record = make(map[string]bool)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering  = make(chan client)
	leaving   = make(chan client)
	messages  = make(chan string) // all incoming client messages
	inputchan = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	//新用户加个判断即可
	if !record[who] {
		messages <- who + " has arrived"
		record[who] = true
	}

	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		inputchan <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//exercise8.13 重置计时器即可
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
		go control(conn)
	}
}

func control(conn net.Conn) {
	ticker := time.NewTicker(5 * time.Minute)

	for {
		select {
		case <-ticker.C:
			conn.Close()
			break
		case str := <-inputchan:
			ticker.Reset(10 * time.Second)
			messages <- str
		}
	}
}

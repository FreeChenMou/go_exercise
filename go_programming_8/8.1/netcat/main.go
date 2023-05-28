package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		location, err := time.LoadLocation(*timezone)
		if err != nil {
			return // set localtion_time fail
		}
		_, err = io.WriteString(c, time.Now().In(location).Format("15:04:05\n")) // 返回相应时区的时间
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

// 练习 8.1： 修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，这个程序可以同时与多个clock服务器通信，

// go run .\eval.go Tokyo=localhost:8010 London=localhost:8020
// go run .\eval.go -port 8020 -timezone Europe/London
// go run .\eval.go -port 8010

var port = flag.String("port", "8000", "服务端口")
var timezone = flag.String("timezone", "Asia/Tokyo", "时区") //默认时区是Tokyo

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}

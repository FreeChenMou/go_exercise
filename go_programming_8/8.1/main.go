package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 练习 8.1： 修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，这个程序可以同时与多个clock服务器通信，

// go run .\eval.go Tokyo=localhost:8010 London=localhost:8020
// go run .\eval.go -port 8020 -timezone Europe/London
// go run .\eval.go -port 8010

var zoneMap = map[string]int{"NewYork": -5, "Tokyo": +9, "London": 0, "Beijing": +8}

func main() {
	for _, arg := range os.Args[1:] {
		zoneAddress := strings.Split(arg, "=")
		go getServerTime(zoneAddress[1], zoneAddress[0])
	}
	for {
		time.Sleep(1)
	}
}

func getServerTime(address string, zoneName string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var sb strings.Builder
	buf := make([]byte, 1)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		char := string(buf[:n])
		if char == "\n" { // 读到换行符，说明读取到一次时间输入结束
			timeStr := sb.String()
			fmt.Printf("%9.9s  %s %s \n", zoneName, address, timeStr)
			sb.Reset() // 读取成功一次，清空builder
			continue
		}
		sb.Write(buf[:n])
	}

}

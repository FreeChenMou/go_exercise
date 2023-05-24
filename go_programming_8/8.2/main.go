package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

var dir, _ = os.Getwd()

//exercise8.2 错漏百出版虚假ftp
func main() {
	listen, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handlerConn(conn) // handle connections concurrently
	}

}

func handlerConn(conn net.Conn) {
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
		str := string(buf[:n])
		if str == "\n" { // 读到换行符，说明读取到一次时间输入结束
			s := sb.String()
			s = s[:len(s)-1]
			split := strings.Split(s, " ")
			switch split[0] {
			case "cd":
				cd(split[1])
			case "ls":
				bytes := list()
				io.WriteString(conn, string(bytes))
			case "get":
				bytes := get(split)
				io.WriteString(conn, string(bytes))
			case "close":
				close(conn)
				return
			}
			sb.Reset()
			continue
		}
		sb.Write(buf[:n])
	}
}

func close(conn net.Conn) {
	conn.Close()
}

func get(commond []string) []byte {
	if len(commond) > 2 {
		return []byte("您的输入不规范：请按照get filename 格式使用\n")
	}
	file, err := os.Open(commond[1]) //默认按照约束输入用户输入
	if err != nil {
		fmt.Fprintf(os.Stderr, "open file err:%v\n", err)
		return []byte("file open err or file not exist\n")
	}

	all, _ := io.ReadAll(file)
	all = append(all, "\n"...)
	return all
}

func cd(commond string) []byte {
	if strings.Contains(commond, "..") {
		index := strings.LastIndex(dir, "\\")
		dir = dir[:index]
		return nil
	}

	if IsDir(commond) {
		dir = commond
		return nil
	} else {
		return []byte("require parameter is not dir")
	}
}

func IsDir(path string) bool {
	println(dir)
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func list() []byte {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd.exe", "/C", "dir", "/a", dir)

	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("无法获取命令的标准输出管道", err.Error())
		return nil
	}
	// 执行Linux命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Linux命令执行失败，请检查命令输入是否有误", err.Error())
		return nil
	}
	// 读取输出
	bytes, err := ioutil.ReadAll(output)
	if err != nil {
		fmt.Println("打印异常，请检查")
		return nil
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait", err.Error())
		return nil
	}
	return bytes
}

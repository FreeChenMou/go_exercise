package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd.exe", "/C", "dir", "/a")

	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("无法获取命令的标准输出管道", err.Error())
		return
	}
	// 执行Linux命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Linux命令执行失败，请检查命令输入是否有误", err.Error())
		return
	}
	// 读取输出
	bytes, err := ioutil.ReadAll(output)
	if err != nil {
		fmt.Println("打印异常，请检查")
		return
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait", err.Error())
		return
	}
	println(string(bytes))
}

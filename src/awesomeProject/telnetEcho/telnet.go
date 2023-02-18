package main

import (
	"fmt"
	"strings"
)

/**
 * @Author: admin
 * @Description:
 * @File: telnet
 * @Version: 1.0.0
 * @Date: 2023/1/15 20:23
 */
func processTelnetCommand(str string, exitChan chan int) bool {
	// @close 指令表示终止会话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		// 告诉外部需要断开连接
		return false
		// @shutdown 指令表示终止进程
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		// 往通道中写入0，阻塞等待接受方处理
		exitChan <- 0
		// 告诉外部需要断开连接
		return false
	}
	// 打印输入的字符串
	fmt.Println(str)
	return true
}

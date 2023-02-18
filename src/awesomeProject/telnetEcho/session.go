package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

/**
 * @Author: admin
 * @Description:
 * @File: session
 * @Version: 1.0.0
 * @Date: 2023/1/15 20:00
 */

// 连接会话逻辑
func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started:")
	// 创建一个网络连接数据的读取器
	readr := bufio.NewReader(conn)

	// 接受数据的循环
	for {
		// 读取字符串，知道碰到回车结束
		str, err := readr.ReadString('\n')
		// 数据读取正确
		if err == nil {
			// 去掉字符串尾部的回车
			str = strings.TrimSpace(str)
			// 处理 Telnet 命令
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}
			// echo逻辑，发生什么数据，原因返回
			conn.Write([]byte(str + "\r\n"))
		} else {
			//发生错误
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}

}

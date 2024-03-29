package main

import (
	"fmt"
	"net"
)

/**
 * @Author: admin
 * @Description:  telnet服务处理
 * @File: server
 * @Version: 1.0.0
 * @Date: 2023/1/15 19:40
 */

// 服务逻辑，传入地址和退出的通道
func server(address string, exitChan chan int) {
	// 根据给定地址进行监听
	l, err := net.Listen("tcp", address)
	// 如果发生错误，打印错误并退出
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}
	// 打印监听地址
	fmt.Println("listen: " + address)
	//延迟关闭监听器
	defer l.Close()
	// 侦听循环
	for {
		// 新链接没有到来时，Accept是阻塞的
		conn, err := l.Accept()
		// 发生任何侦听错误，打印错误并退出服务器
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// 根据连接会话，这个过程需要并行执行
		go handleSession(conn, exitChan)
	}
}

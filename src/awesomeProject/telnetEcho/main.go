package main

import "os"

/**
 * @Author: admin
 * @Description:
 * @File: maain
 * @Version: 1.0.0
 * @Date: 2023/2/4 17:04
 */

func main() {
	// 创建一个程序结束码的通道
	exitChan := make(chan int)

	// 将服务器并运行
	go server("127.0.0.1:7001", exitChan)

	// 通道阻塞，等待接受返回值
	code := <-exitChan

	// 标记程序返回值并退出
	os.Exit(code)
}

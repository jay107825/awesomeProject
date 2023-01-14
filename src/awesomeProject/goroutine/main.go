package main

import (
	"fmt"
	"time"
)

/**
 * @Author: admin
 * @Description:  goroutine 并发
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/1/8 10:16
 */

func running() {
	var times int
	// 无限循环
	for {
		times++

		fmt.Println("回车键终止程序！", times)

		// 延时1秒
		time.Sleep(time.Second)
	}
}

func main() {
	// 并发执行程序
	go running()

	// 接受命令行输入，不做任何事情
	var input string
	fmt.Scanln(&input)

}

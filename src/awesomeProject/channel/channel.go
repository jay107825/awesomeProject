package main

import "fmt"

/**
 * @Author: admin
 * @Description: 通道的输出
 * @File: channel
 * @Version: 1.0.0
 * @Date: 2023/1/8 11:51
 */

func main() {
	ch := make(chan int)

	// 开启一个匿名函数的并发
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine
		ch <- 0
		fmt.Println("exit goroutine")
	}()

	//先执行
	fmt.Println("wait goroutine")
	//等带匿名gorouine
	<-ch

	fmt.Println("all done")
}

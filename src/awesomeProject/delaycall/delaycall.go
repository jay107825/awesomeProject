package main

import (
	"fmt"
	"time"
)

/**
 * @Author: admin
 * @Description: 使用通道响应计时器的事件  延时call
 * @File: delaycall
 * @Version: 1.0.0
 * @Date: 2023/1/14 19:14
 */

func main() {
	// 声明一个退出用的通道
	exit := make(chan int)

	// 打印开始
	fmt.Println("start")
	// 过一秒后，调用匿名函数
	time.AfterFunc(time.Second*12, func() {
		// 过一秒后打印
		fmt.Println("one second after")

		// 通知main() 的 gorouttine 已经结束
		exit <- 0
	})
	// 等待结束
	<-exit

}

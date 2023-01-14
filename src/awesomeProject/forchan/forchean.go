package main

import (
	"fmt"
	"time"
)

/**
 * @Author: admin
 * @Description: 使用for从通道中接受数据
 * @File: forchean
 * @Version: 1.0.0
 * @Date: 2023/1/8 13:49
 */

func main() {
	// 构造一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数

	go func() {
		// 从3循环到0
		for i := 3; i >= 0; i-- {
			//发送3到0之间的数值
			ch <- i
			//每次发送完时等待
			time.Sleep(time.Second)
		}
	}()

	// 遍历接受通道数据
	for data := range ch {

		// 打印通道数据
		fmt.Println(data)

		// 当遇到数据0时。退出接受数据循环
		if data == 1 {
			break
		}
	}
}

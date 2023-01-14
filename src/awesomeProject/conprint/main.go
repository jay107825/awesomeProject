package main

import (
	"fmt"
)

/**
 * @Author: admin
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/1/8 14:06
 */

func printer(c chan interface{}) {
	//先无限循环等待数据
	for {
		// 从channel 中获取一个数据
		data := <-c
		// 将0 视为数据结束
		if data == 0 { //出口 停止for循环
			break
		}
		//打印数据
		fmt.Println("打印数据：", data)
	}
	// 通知main已经结束循环 （我搞定了！）
	c <- 0
}

func main() {
	// 创建一个channel
	c := make(chan interface{})

	// 并行执行printer，传入channel
	go printer(c)

	for i := 1; i < 10; i++ {
		// 将数据通过channel 投给printer
		c <- i
	}

	//通知并发的print结束循环（没数据啦！）
	c <- 0
	// 等待printer结束（搞定喊我！）
	<-c
}

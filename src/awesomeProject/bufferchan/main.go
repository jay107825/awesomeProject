package main

import "fmt"

/**
 * @Author: admin
 * @Description: 带缓冲的通道
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/1/8 15:09
 */

func main() {
	// 创建一个3个元素缓冲大小的整形通道
	ch := make(chan int, 3) // 3为缓冲的大小  大于缓冲， 会报错

	//  查看当前通道的大小
	fmt.Println("start len(ch):", len(ch))

	// 发送3个整形元素到通道
	ch <- 1
	ch <- 2
	ch <- 3
	// 添加后当前通道的大小
	fmt.Println("end len(ch):", len(ch))
}

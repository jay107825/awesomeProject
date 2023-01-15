package main

import "fmt"

/**
 * @Author: admin
 * @Description:  从已关闭的通道接受数据时将不会阻塞
 * @File: closedChenel
 * @Version: 1.0.0
 * @Date: 2023/1/15 17:32
 */

func main() {
	// 创建一个整型带两个缓冲的通道
	ch := make(chan int, 2)

	// 给通到放入两个数据 这时，通道就满了
	ch <- 1
	ch <- 2

	// 关闭缓冲
	close(ch)

	// 遍历缓冲中所有的数据，且多遍历1个
	for i := 0; i < cap(ch)+1; i++ {
		// 从通道中取出数据
		v, ok := <-ch
		// 打印取出数据的
		fmt.Println(v, ok)
	}

}

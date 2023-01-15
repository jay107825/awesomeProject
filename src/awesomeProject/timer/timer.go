package main

import (
	"fmt"
	"time"
)

/**
 * @Author: admin
 * @Description: 定点计时
 * @File: timer
 * @Version: 1.0.0
 * @Date: 2023/1/14 19:36
 */

func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result

}

func main() {
	// 创建一个打点器 ，每500 毫秒触发一次
	ticker := time.NewTicker(time.Microsecond * 500)

	// 创建一个计时器，2秒后触发
	stopper := time.NewTicker(time.Second * 2)

	// 声明计数变量
	var i int

	// 不断地检查通道
	for {
		select {
		case <-stopper.C: //计时器到时了
			fmt.Println("stop")
			// 跳出循环
			goto StopHere // 打点器触发了
		case <-ticker.C:
			// 记录触发多少次
			i++
			fmt.Println("tick", i)
		}
	}

	// 退出标签，使用goto触发
StopHere:
	fmt.Println("done")
}

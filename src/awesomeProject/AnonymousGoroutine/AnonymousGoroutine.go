package main

import (
	"fmt"
	"time"
)

/**
 * @Author: admin
 * @Description: 匿名执行goroutine
 * @File: AnonymousGoroutine
 * @Version: 1.0.0
 * @Date: 2023/1/8 10:42
 */

func main() {
	go func() { // go 后面接匿名函数启动
		var times int
		for true {
			times++
			fmt.Println("回车键结束！", times)
			time.Sleep(time.Second)
		}
	}() // 匿名韩素
	var input string
	fmt.Scanln(&input)
}

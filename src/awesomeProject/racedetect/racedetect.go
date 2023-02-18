package main

import (
	"fmt"
	"sync/atomic"
)

/**
 * @Author: admin
 * @Description: 同步- 保证并发环境下数据访问的正确性
 * @File: racedetect
 * @Version: 1.0.0
 * @Date: 2023/2/18 15:45
 */

var (
	// 序列号
	sep int64 // 序列号生成器中保存上次序列号的变量
)

// 序列号生成器
func GenID() int64 {
	// 尝试原子的增加序列号
	atomic.AddInt64(&sep, 1) // 使用原子操作函数atomic.Addint64() 对seq()函数加1操作。
	return atomic.AddInt64(&sep, 1)
}

func main() {
	// 生成10个并发序列号
	for i := 0; i < 10; i++ {
		go GenID()
	}
	fmt.Println(GenID())
}

// 运行程序时 使用go run -race *.go

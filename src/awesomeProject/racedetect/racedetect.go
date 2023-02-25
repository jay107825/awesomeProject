package main

import (
	"fmt"
	"sync"
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

var (
	//逻辑中使用的某个变量
	count int
	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
)

func GetCount() int {
	// 锁定
	countGuard.Lock()
	//在函数退出时解除锁定
	defer countGuard.Unlock()
	return count
}
func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func main() {
	// 生成10个并发序列号
	//for i := 0; i < 10; i++ {
	//	go GenID()
	//}
	//fmt.Println(GenID())

	// 可以进行并发安全的设置
	SetCount(2)
	// 可以进行并发安全的获取
	fmt.Println(GetCount())

}

// 运行程序时 使用go run -race *.go

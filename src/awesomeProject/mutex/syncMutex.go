package main

import (
	"fmt"
	"sync"
)

/**
 * @Author: admin
 * @Description:   互斥锁  sync.Mutex - 保证同时只有一个gotoutine可以访问共享资源
 * @File: syncMutex
 * @Version: 1.0.0
 * @Date: 2023/2/18 16:13
 */
var (
	// 逻辑中使用的某个变量
	count int

	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
)

func GetCount() int {
	// 锁定
	countGuard.Lock()
	// 在函数退出时解除锁定
	defer countGuard.Unlock()
	return count
}

func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func main() {

	// 可以进行并发安全的设置
	SetCount(1)

	// 可以进行并发安全的获取
	fmt.Println(GetCount())
}

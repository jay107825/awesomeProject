package RWMtex

import "sync"

/**
 * @Author: admin
 * @Description: 读写互斥锁 - 在读写多的环境下比互斥锁更高效
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/2/18 16:33
 */

var (
	// 逻辑中使用某个变量
	count int

	// 与变量对应的使用互斥锁
	countGuard sync.RWMutex
)

func  GetCount() int {
	//锁定
	countGuard.RLocker()

	// 在函数退出的时解除锁定
	defer  countGuard.RUnlock()
	return count
}
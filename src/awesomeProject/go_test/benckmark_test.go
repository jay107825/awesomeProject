package go_test

import (
	"fmt"
	"testing"
)

/**
 * @Author: admin
 * @Description:
 * @File: benckmark
 * @Version: 1.0.0
 * @Date: 2023/5/28 20:32
 */

func Benchmark_Add(b *testing.B) {
	//var n int

	// 重制计时器
	b.ResetTimer()

	// 停止定时器
	b.StopTimer()

	// 开始定时器
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		//n++
		fmt.Sprintf("%d", i)

	}
}

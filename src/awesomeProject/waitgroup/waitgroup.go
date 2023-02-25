package main

import (
	"fmt"
	"net/http"
	"sync"
)

/**
 * @Author: admin
 * @Description: sync.WaitGroup - 保证在并发环境中
 * @File: waitgroup
 * @Version: 1.0.0
 * @Date: 2023/2/24 17:14
 */
func main() {
	// 声明一个等待组
	var wg sync.WaitGroup
	// 准备一系列的网站地址
	var urls = []string{
		"http://www.github.com",
		"https://www.qiniu.com",
		"https://translate.alibaba.com",
	}
	// 遍历这些地址
	for _, url := range urls {
		// 每一个任务开始时,将等待组增加1
		wg.Add(1)
		//开启一个并发
		go func(url string) {
			//使用defer,表示函数完成时将等待组值减1
			defer wg.Done() //等效于执行wg.add(-1)
			//使用http访问提供的地址
			_, err := http.Get(url)
			// 访问完成后，打印地址和可能发生的错误
			fmt.Println(url, err)
			//t通过参数传递url地址
		}(url)
		// 等待所有的任务完成
		wg.Wait()
		fmt.Println("over")
	}
}

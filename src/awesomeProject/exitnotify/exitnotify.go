package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

/**
 * @Author: admin
 * @Description: 避免在不必要的地方使用通道
 * @File: exitnotify
 * @Version: 1.0.0
 * @Date: 2023/7/5 22:12
 */

func socketRecv(conn net.Conn, wg *sync.WaitGroup) {

	// 创建一个接受的缓冲
	buff := make([]byte, 1024)

	// 不停地接受数据
	for {
		// 从套接字中读取数据
		ans, err := conn.Read(buff)
		fmt.Println(ans)
		//需要结束接收，退出循环
		if err != nil {
			break
		}
	}
	//函数已经结束，发出通知
	wg.Done()

}

func main() {
	conn, err := net.Dial("tcp", "www.163.com:80")

	//发生错误时打印错误退出
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建退出通道
	var wg sync.WaitGroup

	// 添加1个任务
	wg.Add(1)

	// 并发执行套件字接收
	go socketRecv(conn, &wg)

	// 在接受时，等待1秒
	time.Sleep(time.Second)

	// 主动退出
	conn.Close()

	// 等待通道退出完毕
	wg.Wait()
	fmt.Println("recv done")
}

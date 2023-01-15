package main

import (
	"errors"
	"fmt"
	"time"
)

/**
 * @Author: admin
 * @Description: 模拟远程过程调用
 * @File: rpc
 * @Version: 1.0.0
 * @Date: 2023/1/8 15:31
 */

func RPCClient(ch chan string, req string) (string, error) {

	//向服务器发送请求
	ch <- req
	// 等待服务器返回
	select {
	case ack := <-ch: //接受到服务器返回数据
		return ack, nil

	case <-time.After(time.Second * 5):
		return "", errors.New("Time out")
	}
}

func RPCServer(ch chan string) {
	for {
		// 接受客户端请求
		data := <-ch

		// 打印接受到的数据
		fmt.Println("server received", data)

		//通过睡眠函数让程序执行阻塞2秒的任务
		time.Sleep(time.Second * 10)

		// 向客户端反馈已收到
		ch <- "server：已收到"
	}
}

func main() {
	// 创建一个无缓冲字符串通道
	ch := make(chan string)

	// 并发执行服务器逻辑
	go RPCServer(ch)

	// 客户端请求数据和接受数据
	recv, err := RPCClient(ch, "向Server 发送数据")
	if err != nil {
		// 发生错误打印
		fmt.Println(err)
	} else {
		// 正常接受到数据
		fmt.Println("client received:", recv)
	}

}

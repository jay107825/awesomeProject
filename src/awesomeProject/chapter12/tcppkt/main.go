package main

import (
	"net"
	"strconv"
)

/**
 * @Author: admin
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/11/28 18:44
 */

func main() {
	// 测试次数
	const TestCount = 100000

	// 测试地址
	const address = "127.0.0.1:8800"

	// 接受到的计数器
	var recvCounter int

	// 实例化一个侦听器
	acceptor := NewAcceptor()

	// 开始侦听
	acceptor.Start(address)

	// 响应封包数据
	acceptor.OnSessionData = func(conn net.Conn, data []byte) bool {
		// 转为字符串
		str := string(data)

		// 转化为数字
		n, err := strconv.Atoi(str)

		// 任何错误或者接受错位时，报错
		if err != nil || recvCounter != n {
			panic("failed")
		}

		// 计算器增加
		recvCounter++

		// 完成计数，关闭侦听器
		if recvCounter >= TestCount {
			acceptor.Stop()
		}
		return true
	}

	// 连接器不断发送数据
	Connector(address, TestCount)

	// 等待结束
	acceptor.Wait()
}

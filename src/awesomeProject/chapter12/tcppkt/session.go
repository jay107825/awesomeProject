package main

import (
	"bufio"
	"net"
)

/**
 * @Author: admin
 * @Description: 服务器连接会话
 * @File: session
 * @Version: 1.0.0
 * @Date: 2023/11/28 17:52
 */

// 连接的会话
func handleSession(conn net.Conn, callback func(net.Conn, []byte) bool) {
	// 创建一个Socket 的读取器
	dataReader := bufio.NewReader(conn)
	// 接受数据的循环
	for {
		// 从连接读取封包
		pkt, err := readPacket(dataReader)
		//回调外部
		if err != nil || !callback(conn, pkt.Body) {
			// 回调要求退出
			conn.Close()
			break
		}
	}
}

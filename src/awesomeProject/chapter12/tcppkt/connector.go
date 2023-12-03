package main

import (
	"fmt"
	"net"
	"strconv"
)

/**
 * @Author: admin
 * @Description: 连接器
 * @File: connector
 * @Version: 1.0.0
 * @Date: 2023/11/26 20:48
 */

// 连接器，传入连接地址和发送封包次数
func Connector(address string, sendTimes int) {
	// 尝试用Socker 连接地址
	conn, err := net.Dial("tcp", address)

	// 发生错误时退出
	if err != nil {
		fmt.Println(err)
		return
	}

	// 循环指定次数
	for i := 0; i < sendTimes; i++ {
		// 将循环序号 转为字符串
		str := strconv.Itoa(i)
		// 发送封包
		if err := writePacket(conn, []byte(str)); err != nil {
			fmt.Println(err)
			break
		}
	}
}

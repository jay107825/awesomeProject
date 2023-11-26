package tcppkt

import (
	"fmt"
	"net"
	"sync"
)

/**
 * @Author: admin
 * @Description:  接受器
 * @File: acceptor
 * @Version: 1.0.0
 * @Date: 2023/11/26 21:03
 */

type Acceptor struct {
	// 保存侦听器
	i net.Listener

	// 侦听器的停止同步
	wg sync.WaitGroup

	// 连接的数据回调
	OnSessionData func(net.Conn, []byte) bool
}

// 异步开始侦听
func (a *Acceptor) Start(address string) {
	go a.listen(address)
}

func (a *Acceptor) listen(address string) {
	// 侦听开始，添加一个任务
	a.wg.Add(1)

	// 在退出函数时，结束侦听任务
	defer a.wg.Done()
	var err error

	// 根据给定地址进行侦听
	a.i, err = net.Listen("tcp", address)

	// 如果侦听发生错误。打印错误并错误
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 侦听循环
	for {
		// 新连接没有到来时,Accept()
		//conn, err := a.i.Accept()

		// 如发生任何侦听错误。打印错误并退出服务器
		if err != nil {
			break
		}

		// 根据连接开启会话
		// todo 需要定义方法
		//go handleSession(conn, a.OnSessionData)
	}
}

// 停止侦听器
func (a *Acceptor) Stop() {
	a.i.Close()
}

// 等待侦听完全停止
func (a *Acceptor) Wait() {
	a.wg.Wait()
}

// 实例化一个侦听器
func NewAcceptor() *Acceptor {

	return nil
	// return &Accepto()   // todo  346 page
}

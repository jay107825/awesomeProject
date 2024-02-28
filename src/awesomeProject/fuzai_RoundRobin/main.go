package main

/**
 * @Author: admin
 * @Description: 负载均衡  轮寻法
 * @File: f
 * @Version: 1.0.0
 * @Date: 2024/2/28 21:40
 */

import (
	"fmt"
	"sync"
)

type RoundRobin struct {
	servers []string
	index   int
	lock    sync.Mutex
}

func NewRoundRobin(servers []string) *RoundRobin {
	return &RoundRobin{
		servers: servers,
		index:   0,
	}
}

func (rr *RoundRobin) GetNextServer() string {
	rr.lock.Lock()
	defer rr.lock.Unlock()

	server := rr.servers[rr.index]
	rr.index = (rr.index + 1) % len(rr.servers)
	return server
}

func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	rr := NewRoundRobin(servers)

	for i := 0; i < 10; i++ {
		fmt.Println("Request sent to :", rr.GetNextServer())
	}
}

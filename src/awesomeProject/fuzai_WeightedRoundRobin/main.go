/**
 * @Author: admin
 * @Description: 加权轮询法
 * @File: main
 * @Version: 1.0.0
 * @Date: 2024/2/28 21:53
 */

package main

import (
	"fmt"
	"sync"
)

type WeightedRoundRobin struct {
	servers    []string
	weights    []int
	currentIdx int
	lock       sync.Mutex
}

func NewWeightedRoundRobin(servers []string, weights []int) *WeightedRoundRobin {
	return &WeightedRoundRobin{
		servers:    servers,
		weights:    weights,
		currentIdx: 0,
	}
}

func (wrr *WeightedRoundRobin) GetNextServer() string {
	wrr.lock.Lock()
	defer wrr.lock.Unlock()

	server := wrr.servers[wrr.currentIdx]
	wrr.currentIdx = (wrr.currentIdx + 1) % len(wrr.servers)
	return server
}

func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	weights := []int{2, 1, 3} // Server1权重为2，Server2权重为1，Server3权重为3

	wrr := NewWeightedRoundRobin(servers, weights)

	for i := 0; i < 10; i++ {
		fmt.Println("Request sent to:", wrr.GetNextServer())
	}
}

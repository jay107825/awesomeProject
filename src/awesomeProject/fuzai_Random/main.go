package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @Author: admin
 * @Description: 负载均衡_随机法
 * @File: main
 * @Version: 1.0.0
 * @Date: 2024/2/28 21:47
 */

type Random struct {
	servers []string
}

func NewRandom(servers []string) *Random {
	return &Random{
		servers: servers,
	}
}

func (r *Random) GetRandomServer() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(r.servers))
	return r.servers[index]
}

func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	random := NewRandom(servers)

	for i := 0; i < 10; i++ {
		fmt.Println("Request sent to:", random.GetRandomServer())
	}
}

/**
 * @Author: admin
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2024/2/28 21:56
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WeightedRandom struct {
	servers []string
	weights []int
}

func NewWeightedRandom(servers []string, weights []int) *WeightedRandom {
	return &WeightedRandom{
		servers: servers,
		weights: weights,
	}
}

func (wr *WeightedRandom) GetWeightedRandomServer() string {
	rand.Seed(time.Now().UnixNano())

	totalWeight := 0
	for _, weight := range wr.weights {
		totalWeight += weight
	}

	randWeight := rand.Intn(totalWeight)

	for i, weight := range wr.weights {
		if randWeight < weight {
			return wr.servers[i]
		}
		randWeight -= weight
	}

	return wr.servers[len(wr.servers)-1]
}

func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	weights := []int{2, 1, 3} // Server1权重为2，Server2权重为1，Server3权重为3

	wr := NewWeightedRandom(servers, weights)

	for i := 0; i < 10; i++ {
		fmt.Println("Request sent to:", wr.GetWeightedRandomServer())
	}
}

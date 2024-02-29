/**
 * @Author: admin
 * @Description: ip hash法
 * @File: main
 * @Version: 1.0.0
 * @Date: 2024/2/28 21:58
 */

package main

import (
	"fmt"
	"hash/fnv"
)

type IPHash struct {
	servers []string
}

func NewIPHash(servers []string) *IPHash {
	return &IPHash{
		servers: servers,
	}
}

func (ih *IPHash) GetServerByIP(ip string) string {
	h := fnv.New32a()
	h.Write([]byte(ip))
	index := int(h.Sum32()) % len(ih.servers)
	return ih.servers[index]
}

func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	ih := NewIPHash(servers)

	ips := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"}

	for _, ip := range ips {
		fmt.Printf("Request from IP %s sent to: %s\n", ip, ih.GetServerByIP(ip))
	}
}

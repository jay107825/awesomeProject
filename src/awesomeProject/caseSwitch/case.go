package main

import "fmt"

/**
 * @Author: admin
 * @Description: 类型分分支 = 批量判断空接口中变量的类型
 * @File: main
 * @Version: 1.0.0
 * @Date: 2022/12/24 19:55
 */

func PrintType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}

}

func main() {
	PrintType(1024)
	PrintType("pig")
	PrintType(false)

}

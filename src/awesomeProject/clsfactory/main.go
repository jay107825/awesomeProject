package main

import (
	"awesomeProject/clsfactory/base"
	_ "awesomeProject/clsfactory/cls1" // 引用cls1包  init注册
	_ "awesomeProject/clsfactory/cls2" // 引用cls2包  init注册
)

/**
 * @Author: admin
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/1/8 09:49
 */

func main() {
	// 根据字符串动态创建一个Class实例
	c1 := base.Create("Class1")
	c1.Do()

	// 根据字符串动态创建一个Class2 实例
	c2 := base.Create("Class2")
	c2.Do()
}

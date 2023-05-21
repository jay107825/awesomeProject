package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: admin
 * @Description:  反射调用函数
 * @File: redflectfunction
 * @Version: 1.0.0
 * @Date: 2023/4/22 17:40
 */

// 普通函数
func add(a, b int) int {
	return a + b
}

func main() {
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)

	//构造函数参数，传入两个整形值
	paramList := []reflect.Value{reflect.ValueOf(5), reflect.ValueOf(20)}

	// 反射调用函数
	reList := funcValue.Call(paramList)

	// 获取第一个返回值，取整数值®
	fmt.Println(reList[0].Int())
}

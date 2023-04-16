package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: admin
 * @Description: 反射值对象
 * @File: reflectValue
 * @Version: 1.0.0
 * @Date: 2023/3/27 22:01
 */
func main() {
	// 声明整形变量a并赋初值
	var a int = 1024

	// 获取 变量a的反射对象
	valueOfa := reflect.ValueOf(a)

	//获取interface{} 类型的值，通过类型断言转换
	var getA int = valueOfa.Interface().(int)

	// 获取 64位的值 强制转换int类型
	var getA2 int = int(valueOfa.Int())
	
	fmt.Println(getA, getA2)
}
package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: admin
 * @Description:  各种反射的demo
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/5/21 10:25
 */

func main() {
	/*
		通过类型创建反射实例
	*/
	var a int
	// 取变量a的反射类型对象
	typeOfa := reflect.TypeOf(a)
	// 根据反射类型对象创建类型实例
	aIns := reflect.New(typeOfa)
	// 输出Value 的类型和种类
	fmt.Println(aIns.Type(), aIns.Kind())

	// 使用反射调用函数 调用方
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	// 构造函数参数，传入俩个整型值
	ParamList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}

	// 反射调用函数
	retList := funcValue.Call(ParamList)

	// 获取第一个返回值，取整数
	fmt.Println(retList[0].Int())
}

/*
	使用反射调用函数
*/
// 普通函数
func add(a, b int) int {
	return a + b
}

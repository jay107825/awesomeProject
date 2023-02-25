package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: admin
 * @Description: 反射类型的对象
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/2/24 17:53
 */

// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

func main() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	// 声明一个空结构体
	type cat struct {
	}
	// 获取结构体实例的放射类型对
	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射类型对象的名称和种类
	fmt.Println("----------------------------")
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	// 获取 zero 常量的反射类型的对象
	typeOfA = reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

}

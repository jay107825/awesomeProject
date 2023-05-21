package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: admin
 * @Description:  反射对象的空和有效性判断
 * @File: vaildnisnli
 * @Version: 1.0.0
 * @Date: 2023/4/16 21:24
 */

func main() {

	//*int的空指针
	var a *int
	fmt.Println("var a int：", reflect.ValueOf(a).IsValid())

	// nil值  对niu值做判定
	fmt.Println("nil：", reflect.ValueOf(nil).IsValid())

	// *int 类型的空指针
	fmt.Println("（*int）(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())

	// 实例化一个结构体
	s := struct{}{}

	//尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员：", reflect.ValueOf(s).FieldByName("").IsValid())

	// 尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的结构体方法：", reflect.ValueOf(s).MethodByName("").IsValid())

	// 实例化一个map
	m := map[int]int{}

	// 尝试从map中查找一个不存在的键
	fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
}

package main

import "fmt"

/**
 * @Author: jay
 * @Description: 空接口类型- 能保存所有值的类型
 * @File: nullInterface.go
 * @Version: 1.0.0
 * @Date: 2022/12/11 18:37
 */

func main() {
	/**
	将值保存到空接口
	*/
	var any interface{}
	any = 1
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	/**
	从空接口获取值
	*/
	//声明a变量，类型int，
	var a int = 1
	var i interface{} = a
	var b int = i.(int)
	fmt.Println(b)

	/**
	空接口的比较
	*/
	var c interface{} = 100
	var d interface{} = "hi"
	fmt.Println(c == d)

	/**
	不能比较空接口的动态值
	*/
	var e interface{} = []int{10}
	var f interface{} = []int{20}
	fmt.Println(e, f)
	//fmt.Println(e == f)  //会报错

}

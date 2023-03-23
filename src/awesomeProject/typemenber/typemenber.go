package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: admin
 * @Description: 使用反射获取结构体的成员类型
 * @File: typemenber
 * @Version: 1.0.0
 * @Date: 2023/3/4 13:54
 */

func main() {
	// 声明一个空结构体
	type cat struct {
		Name string
		// 带有结构体tag的字段
		Type int `json:"type" id:"100"`
	}
	// 创建cat的实例
	ins := cat{Name: "mimi", Type: 1}

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)

	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构字段类型
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Println("name:%v tag:'v'\n", fieldType.Name, fieldType.Tag)
	}
	//通过字段名，找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		//从tag中取出需要的tag
		fmt.Println("重点：", catType.Tag.Get("json"), catType.Tag.Get("id"))

	}
}

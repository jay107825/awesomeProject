package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: admin
 * @Description: 反射结构体标签
 * @File: reflectiveStructure
 * @Version: 1.0.0
 * @Date: 2023/3/27 21:35
 */

func main() {
	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
	}

	typeOfCat := reflect.TypeOf(cat{})

	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}

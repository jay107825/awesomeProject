package main

import (
	"bytes"
	"reflect"
)

/**
 * @Author: admin
 * @Description:
 * @File: struct
 * @Version: 1.0.0
 * @Date: 2023/5/21 11:00
 */

// 将结构体序列化为json格式并输出到缓冲中
func writeStruct(buff *bytes.Buffer, value reflect.Value) error {
	// 取值的类型对象
	valueType := value.Type()

	// 写入结构体左大括号
	buff.WriteString("{")

	// 遍历结构体的所有值
	for i := 0; i < value.NumField(); i++ {
		// 获取每个字段的字段值（reflect.Value）
		fieldvalue := value.Field(i)

		// 获取每个字段的类型（reflect.StructField）
		fieldType := valueType.Field(i)

		// 写入字段名左双印号
		buff.WriteString("\"")

		// 写入字段名
		buff.WriteString(fieldType.Name)

		// 写入字段名右双引号和冒号
		buff.WriteString("\":")

		// 写入每个人字段值
		writeAny(buff, fieldvalue)

		// 每个字段尾部写入逗号，最后一个字段不添加
		if i < value.NumMethod()-1 {
			buff.WriteString(",")
		}
	}
	// 写入结构体右大括号
	return nil
}

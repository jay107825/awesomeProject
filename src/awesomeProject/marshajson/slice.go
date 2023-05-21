package main

import (
	"bytes"
	"reflect"
)

/**
 * @Author: admin
 * @Description: 切片序列化
 * @File: slice
 * @Version: 1.0.0
 * @Date: 2023/5/7 21:43
 */

func writeSlice(buff *bytes.Buffer, value reflect.Value) error {
	//fmt.Println("执行到此！！！")
	//	写入切片开始标记
	buff.WriteString("[")
	//	遍历每个元素
	for s := 0; s < value.Len(); s++ {
		sliceValue := value.Index(s)
		// 写入每个元素
		writeAny(buff, sliceValue)
		//	每个元素尾部写入逗号，最后一个字段添加
		if s < value.Len()-1 {
			buff.WriteString(",")
		}
	}
	// 写入切片结束印记
	buff.WriteString("]")
	return nil
}

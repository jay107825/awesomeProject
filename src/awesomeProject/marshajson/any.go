package main

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
)

/**
 * @Author: admin
 * @Description: 任意值序列化
 * @File: any
 * @Version: 1.0.0
 * @Date: 2023/5/7 21:19
 */

// 将任意值转换为json格式并输出到缓冲中
func writeAny(buff *bytes.Buffer, value reflect.Value) error {

	switch value.Kind() {
	case reflect.String:
		// 写入带有双引号扩起来的字符串
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		//	将整型转换为字符串并写入缓冲中
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buff, value)
	case reflect.Struct:
		return writeStruct(buff, value)
	default:
		// 遇到不认识的种类，返回错误
		return errors.New("unsupport kind:" + value.Kind().String())
	}
	return nil
}

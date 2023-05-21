package main

import (
	"bytes"
	"fmt"
	"reflect"
)

/**
 * @Author: admin
 * @Description:
 * @File: marshal
 * @Version: 1.0.0
 * @Date: 2023/5/7 21:15
 */

func Marshaljson(v interface{}) (string, error) {
	// 准备一个缓冲
	var b bytes.Buffer
	// 将任意值转换为json并输出到缓冲中
	if err := writeAny(&b, reflect.ValueOf(v)); err == nil {
		fmt.Println("Marshaljson b.String():", b.String())
		return b.String(), nil
	} else {
		return "", err
	}
}

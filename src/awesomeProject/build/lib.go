package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: admin
 * @Description:
 * @File: lib
 * @Version: 1.0.0
 * @Date: 2023/5/21 21:06
 */

func PkgFunc() {
	fmt.Println("pkg")
	fmt.Println(reflect.ValueOf(PkgFunc).Type(), reflect.ValueOf(PkgFunc).Kind())
}

package pkg2

import "fmt"

/**
 * @Author: admin
 * @Description:  理解包导入后的 init()函数初始化顺序
 * @File: pkg2
 * @Version: 1.0.0
 * @Date: 2023/1/7 21:11
 */

func Execpkg2() {
	fmt.Println("ExecPkg2")
}

func init() {
	fmt.Println("pkg2 inti")
}

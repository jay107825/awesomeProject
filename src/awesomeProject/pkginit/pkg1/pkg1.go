package pkg1

import (
	"awesomeProject/pkginit/pkg2"
	"fmt"
)

/**
 * @Author: admin
 * @Description: pkg1
 * @File: pkg1
 * @Version: 1.0.0
 * @Date: 2023/1/7 21:18
 */

func ExecPkg1() {
	fmt.Println("ExexPkg1")

	pkg2.Execpkg2()
}

func init() {
	fmt.Println("pkg1 init")
}

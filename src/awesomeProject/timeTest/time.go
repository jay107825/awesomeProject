package main

import (
	"fmt"
	"time"
)

/**
 * @Author: admin
 * @Description: timetest
 * @File: time
 * @Version: 1.0.0
 * @Date: 2023/4/13 23:04
 */

func main() {
	test := time.Now()
	testbew := test.Format("2006-01-0215:04:05.000000+08")
	fmt.Println(`updata zeus set time=` + testbew + ` where id=1;`)
}

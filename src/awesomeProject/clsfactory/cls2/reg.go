package cls2

import (
	"awesomeProject/clsfactory/base"
	"fmt"
)

/**
 * @Author: admin
 * @Description:
 * @File: reg
 * @Version: 1.0.0
 * @Date: 2023/1/8 09:44
 */

// 定义类2
type Class2 struct {
}

// 实现Class接口
func (c *Class2) Do() {
	fmt.Println("Class2  这里输出")
}

func init() {
	// 在启动时注册类2工厂
	base.Registet("Class2", func() base.Class {
		return new(Class2)
	})
}

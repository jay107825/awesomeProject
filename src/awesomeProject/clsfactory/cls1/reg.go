package cls1

import (
	"awesomeProject/clsfactory/base"
	"fmt"
)

/**
 * @Author: admin
 * @Description:
 * @File: reg
 * @Version: 1.0.0
 * @Date: 2023/1/8 09:34
 */
// 定义类1
type Class1 struct {
}

// 实现Class接口
func (c Class1) Do() {
	fmt.Println("Class1 这里输出")
}

// 初始化
func init() {
	//在启动时注册类1工厂  调用Redistet接口进行注册
	base.Registet("Class1", func() base.Class {
		return new(Class1)
	})
}

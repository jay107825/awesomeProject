package base

/**
 * @Author: admin
 * @Description: 工厂模式自动注册 into理解
 * @File: factory
 * @Version: 1.0.0
 * @Date: 2023/1/8 09:18
 */

// 类接口
type Class interface {
	Do()
}

var (
	//保存注册好的工厂信息
	factorByName = make(map[string]func() Class)
)

// 注册一个类生成工厂
func Registet(name string, factory func() Class) {
	factorByName[name] = factory
}

// 根据名称 创建对应的类
func Create(name string) Class {
	if f, ok := factorByName[name]; ok {
		return f()
	} else {
		panic("name not found !!!")
	}
}

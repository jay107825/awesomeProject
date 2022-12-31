package main

import (
	"fmt"
)

/**
 * @Author: admin
 * @Description: 使用空接口实现可以保存任意值的字典
 * @File: dict
 * @Version: 1.0.0
 * @Date: 2022/12/11 21:21
 */

// 字典结构
type Dictionary struct {
	data map[interface{}]interface{} //建值都为interface  实现了Diccionary的内部实现是一个键值均为interface的map
}

// 根据键获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key] //通过map直接获取值，键值不存在则返回nil
}

// 设置建值
func (d *Dictionary) Set(key, value interface{}) {
	d.data[key] = value
}

// 遍历所有的键值，如果回调返回值为false，停止遍历
func (d *Dictionary) visit(callback func(k, v interface{}) bool) { //定义回调，类型为func（k，v interdace{}）boool,意思是返回键值数据（k，v）bool表示遍历流程控制  返回true时继续遍历
	if callback == nil {
		return //当callback 为空时，退出遍历
	}
	for k, v := range d.data {
		if !callback(k, v) { //根据callback的返回值，决定是否继续遍历
			return
		}
	}
}

// 清空所有数据
func (d *Dictionary) clear() {
	d.data = make(map[interface{}]interface{}) // map没有独立的复位内部元素的操作，需要复位时，使用make创建新的实例；go的垃圾回收是并行的
}

// 创建一个字典
func NewDictionary() *Dictionary {
	d := &Dictionary{}
	//初始化
	d.clear()
	return d
}

func main() {
	//创建字典实例
	dict := NewDictionary()
	//添加游戏数据
	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 38)
	dict.Set("Don Hungry", 24)

	//获取值并打印值
	favorite := dict.Get("Terra Craft")
	fmt.Println("hahah favorite", favorite)

	//遍历所有元素
	dict.visit(func(k, v interface{}) bool {
		//将值转为int类型，判断是否大于40
		if v.(int) > 40 {
			fmt.Println(k, "很贵")
			return true
		}
		fmt.Println(k, "不贵")
		return true
	})
}

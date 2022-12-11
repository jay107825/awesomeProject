package dict

/**
 * @Author: admin
 * @Description: 使用空接口实现可以保存任意值的字典
 * @File: dict
 * @Version: 1.0.0
 * @Date: 2022/12/11 21:21
 */

// 字典结构
type Dictionary struct {
	data map[interface{}]interface{} //建值都为interface
}

// 根据键获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

// 设置建值
func (d *Dictionary) Set(key, value interface{}) {
	d.data[key] = value
}

// 遍历所有的键值，如果回调返回值为false，停止遍历
func (d *Dictionary) visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

// 清空所有数据
func (d *Dictionary) clear() {
	d.data = make(map[interface{}]interface{})
}

// 创建一个字典
func NewDictionary() *Dictionary {
	d := &Dictionary{}
	//初始化
	d.clear()
	return d
}

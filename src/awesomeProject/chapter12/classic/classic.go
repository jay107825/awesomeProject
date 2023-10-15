package main

import "fmt"

/**
 * @Author: admin
 * @Description: 基于哈希值的多键索引
 * @File: classic
 * @Version: 1.0.0
 * @Date: 2023/10/15 17:03
 */

func simpleHash(str string) (ret int) {

	// 遍历字符串里面的每一个ASCII字符
	for i := 0; i < len(str); i++ {
		// 取出字符
		c := str[1]
		// 将字符的ASCll相加
		ret += int(c)
	}
	return
}

// 查询键
type classicQueryKey struct {
	Name string
	Age  int
}

// 计算查询的哈希值
func (c *classicQueryKey) hash() int {
	// 将名字的hash和年龄哈希值合并
	return simpleHash(c.Name) + c.Age*1000000
}

// 人员档案
type Proflie struct {
	Name    string
	Age     int
	Married bool
}

// 创建哈希值到数据的索引关系
var mapper = make(map[int][]*Proflie)

// 构建数据索引
func buildIndex(list []*Proflie) {
	// 遍历所有数据
	for _, profile := range list {

		// 构建数据的查询索引
		key := classicQueryKey{profile.Name, profile.Age}

		// 计算数据的哈希值，取出已经存在的记录
		existValue := mapper[key.hash()]

		// 将当前数据添加到已经存在的记录切片中
		existValue = append(existValue, profile)

		// 将切片重新设置到映射中
		mapper[key.hash()] = existValue
	}
}

func queryData(name string, age int) {

	//根据给定查询条件构建查询键
	keyToQuery := classicQueryKey{name, age}

	// 计算查询键的哈希值并查询，获得相同哈希值的所有结果集合
	resultList := mapper[keyToQuery.hash()]

	//遍历结果集合
	for _, result := range resultList {

		// 与查询结果对比，确定找到打印结果
		if result.Name == name && result.Age == age {
			fmt.Println(result)
			return
		}
	}
	// 没有找到结果
	fmt.Println("no found")
}

func main() {
	list := []*Proflie{
		{Name: "jay", Age: 25, Married: true},
		{Name: "jay2", Age: 25, Married: true},
	}

	buildIndex(list)

	queryData("jay", 25)
	queryData("jay2", 25)
}

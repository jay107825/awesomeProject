package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func queryData(name string, age int) interface{} {

	//根据给定查询条件构建查询键
	keyToQuery := classicQueryKey{name, age}

	// 计算查询键的哈希值并查询，获得相同哈希值的所有结果集合
	resultList := mapper[keyToQuery.hash()]

	//遍历结果集合
	for _, result := range resultList {

		// 与查询结果对比，确定找到打印结果
		if result.Name == name && result.Age == age {
			//fmt.Println(result)
			return result
		}
	}
	// 没有找到结果
	return nil
}

func data() interface{} {
	reader := bufio.NewReader(os.Stdin)

	// 读取用户输入，直到按下 Ctrl+D（Unix 系统）或 Ctrl+Z（Windows 系统）
	fmt.Print("请输入一些文本：")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}

	// 去除输入字符串末尾的换行符
	input = input[:len(input)-1]

	fmt.Println("你要查询的是:", input)
	return input
}

func main() {
	list := []*Proflie{
		{Name: "jay", Age: 25, Married: true},
		{Name: "jay2", Age: 25, Married: true},
	}

	buildIndex(list)

	input := data()
	str := input.(string)

	inputage := data()
	ages := inputage.(string)

	num, err := strconv.Atoi(ages)
	if err != nil {
		fmt.Println("转换失败，请输入正整数:", err)
		return
	}

	result := queryData(str, num)
	if result != nil {
		fmt.Println("查询结果是：", result)
	} else {
		fmt.Println("查讯失败,查询内容不存在！")
	}
}

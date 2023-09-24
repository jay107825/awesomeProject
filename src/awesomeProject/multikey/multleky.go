package main

import "fmt"

/**
 * @Author: admin
 * @Description:  利用mao的特性多键索引及查询
 * @File: multleky
 * @Version: 1.0.0
 * @Date: 2023/9/24 18:55
 */
// 查询键
type querKey struct {
	Name string
	Age  int
}
type Profile struct {
	Name    string
	Age     int
	Married bool
}

// 创建查询键到数据的映射
var mapper = make(map[querKey]*Profile)

// 构建查询索引
func buildIndex(list []*Profile) {
	// 遍历所有数据
	for _, profile := range list {
		key := querKey{
			Name: profile.Name,
			Age:  profile.Age,
		}
		//保存查询键
		mapper[key] = profile
	}
}

func queryData(name string, age int) *Profile {
	key := querKey{name, age}

	result, ok := mapper[key]

	if ok {
		fmt.Println(result)
		return result
	} else {
		fmt.Println("on found")
		return nil
	}
}

func main() {
	lists := []*Profile{
		{Name: "jay", Age: 25, Married: false},
		{Name: "chai", Age: 24, Married: true},
		{Name: "xiao", Age: 20},
	}
	buildIndex(lists)
	data := queryData("xiao", 20)
	fmt.Println(data)
}

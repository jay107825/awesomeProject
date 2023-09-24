package mapSelect

/**
 * @Author: admin
 * @Description: map的多键索引
 * @File: mapselect
 * @Version: 1.0.0
 * @Date: 2023/9/24 17:47
 */

// 结构体 创建人员档案
//type Profile struct {
//	Name    string
//	Age     int
//	Married bool
//}
//
//// 基于哈希值的多键索引及查询
//func simpleHash(str string) (ret int) {
//	// 遍历字符串的每一个ASCII字符
//	for i := 0; i < len(str); i++ {
//		// 取出字符
//		c := str[i]
//		//将字符的AACII字符相加
//		ret += int(c)
//	}
//	return
//}

//// 查询键
//type classicQueryKey struct {
//	Name string
//	Age  int
//}

// 计算查询键的哈希值
//func (c *classicQueryKey) hash() int {
//	// 将名字的Hash和年龄哈希合并
//	return simpleHash(c.Name) + c.Age*1000000
//}

// 构建数据索引
//func buildIndex(list []*Profile) {
//	// 遍历所有数据
//	for _, profile := range list {
//		// 构建数据的查询索引
//		key := classicQueryKey{profile.Name, profile.Age}
//
//		// 计算数据的哈希值，取出已经存在的记录
//		existValue := mapper[key.hash()]
//
//		//将当前数据加到已经存在的记录切片中
//		existValue = append(existValue, profile)
//
//		// 将切片重新设置到映射中
//		mapper[key.hash()] = existValue
//	}
//}
//
//// 查询逻辑
//func queryData(name string, age int) {
//	// 根据给定查询条件构建查询键
//	keyToQuery := classicQueryKey{name, age}
//	//计算查询键哈希值并查询,获得相同哈希值的所有结果集合
//	resultList := mapper[keyToQuery.hash()]
//
//	// 遍历结果结合
//	for _, result := range resultList {
//		//与查询结果对比，确认找到打印结果
//		if result.Name == name && result.Age == age {
//			fmt.Println(result)
//			return
//		}
//	}
//	// 没有查到，打印结果
//	fmt.Println("no found")
//}
//
//func main() {
//	lists := []*Profile{
//		{Name: "jay", Age: 25, Married: false},
//		{Name: "chai", Age: 24, Married: true},
//		{Name: "xiao", Age: 20},
//	}
//
//}

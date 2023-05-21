package marshajson

/**
 * @Author: admin
 * @Description: 将结构体的数据保存为json格式的文本数据
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/5/7 20:58
 */

/**


func main() {
	// 声明技能结构
	type Skill struct {
		Name  string
		Level int
	}

	// 声明角色结构
	type Actor struct {
		Name  string
		Age   int
		Skill []Skill
	}

	//	填充基本角色数据
	a := Actor{
		Name: "cow boy",
		Age:  24,
		Skill: []Skill{
			{Name: "jay", Level: 1},
			{Name: "chai", Level: 2},
			{Name: "chaizhijie", Level: 3},
		},
	}
	if result, err := Marshajson(a); err == nil {

	}
}
*/

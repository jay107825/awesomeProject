package marshajson

import "bytes"

/**
 * @Author: admin
 * @Description:
 * @File: marshal
 * @Version: 1.0.0
 * @Date: 2023/5/7 21:15
 */

func Marshaljson(v interface{}) (string, error) {
	// 准备一个缓冲
	var b bytes.Buffer
	// 将任意值转换为json并输出到缓冲中
	if err := writeAny
}

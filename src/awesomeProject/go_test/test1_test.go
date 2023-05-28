package go_test

import (
	"fmt"
	"testing"
)

/**
 * @Author: admin
 * @Description: go test
 * @File: test1_test
 * @Version: 1.0.0
 * @Date: 2023/5/28 20:02
 */

func TestA(t *testing.T) {
	t.Logf("jay test")
}

func TestAk(t *testing.T) {
	t.FailNow()
	t.Log("hhaah")
}

func TestC(t *testing.T) {
	fmt.Println("前面")
	t.Fail()
	fmt.Println("后面！！！")
}

func TestB(t *testing.T) {
	t.Log("TestB")
}

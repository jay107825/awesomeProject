package main

import "fmt"

/**
 * @Author: admin
 * @Description: 电子支付 - 现金支付
 * @File: casehandalipay
 * @Version: 1.0.0
 * @Date: 2022/12/31 17:17
 */

// 电子支付方式
type Alipay struct {
}

// 为Alipay 添加CanUseFaceID()方法，表示现金支付，支持刷脸
func (a *Alipay) CanUseFaceID() {
}

// 现金支付方式
type Cash struct {
}

// 为Cash 添加Stuolen() 方法，表示现金支付方式出现偷窃情况
func (a Cash) Stolen() {
}

// 具备刷脸特性的接口
type CantainCanUseFaceID interface {
	CanUseFaceID()
}

// 具备被偷性的接口
type ContinStolen interface {
	Stolen()
}

// 打印支付方式具备的特点
func Print(payMethod interface{}) {
	switch payMethod.(type) {
	case CantainCanUseFaceID: //可以刷脸
		fmt.Printf("%T can use faceid\n", payMethod)
	case ContinStolen: //可能被偷
		fmt.Printf("%T may be stolen\n", payMethod)
	}
}
func main() {
	//使用电子支付判断
	Print(new(Alipay))
	// 使用现金判断
	Print(new(Cash))
}

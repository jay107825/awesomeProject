package main

import (
	"fmt"
)

/**
 * @Author: admin
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/1/1 18:06
 */

// 闲置状态
type IdleState struct {
	StateInfo // 使用Stateinfo 实现基础接口ß
}

func (i *IdleState) EnadleSameTransit() bool {
	//TODO implement me
	panic("implement me")
}

// 重新实现状态开始
func (i *IdleState) OnBegin() {
	fmt.Println("空闲状态开始")
}

// 重新实现状态结束
func (i *IdleState) OnEnd() {
	fmt.Println("空闲状态结束")
}

// 转移状态
type MoveState struct {
	StateInfo
}

// 允许移动状态互相转换
func (m *MoveState) EnadleSameTransit() bool {
	return true
}

func (m *MoveState) OnBengin() {
	fmt.Println("移动状态开始")
}

// 跳跃状态
type JumpState struct {
	StateInfo
}

// 跳跃状态不能转移到移动状态
func (j *JumpState) EnadleSameTransit() bool {
	return false
}

func (j *JumpState) OnBegin() {
	fmt.Println("跳转状态开始")
}

func main() {
	// 实例化一个状态管理器
	sm := NewStateManger()
	// 响应状态转移的通知
	sm.OnChange = func(from, to State) {
		//打印状态转移的流向
		fmt.Printf("%s ---> %s\n\n", StateName(from), StateName(to))
	}

	//添加3个状态
	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))

	// 在不同状态间转移
	transitAndReport(sm, "IdleState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "IdleState")
}

// 封装转移状态和输出日志
func transitAndReport(sm *StateManager, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("FALLED! %s --> %s,%s\n\n", sm.CurrState().Name(), target, err.Error())
	}
}

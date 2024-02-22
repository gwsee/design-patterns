package main

import "fmt"

// State 是一个接口，代表信号灯的各种状态
type State interface {
	Handle(context *TrafficLight)
	String() string
}

// RedState 表示红灯状态
type RedState struct{}

func (rs *RedState) Handle(context *TrafficLight) {
	fmt.Println("Red Light - Stop")
	// 转换到下一个状态
	context.SetState(&GreenState{})
}

func (rs *RedState) String() string {
	return "Red Light"
}

// GreenState 表示绿灯状态
type GreenState struct{}

func (gs *GreenState) Handle(context *TrafficLight) {
	fmt.Println("Green Light - Go")
	// 转换到下一个状态
	context.SetState(&YellowState{})
}

func (gs *GreenState) String() string {
	return "Green Light"
}

// YellowState 表示黄灯状态
type YellowState struct{}

func (ys *YellowState) Handle(context *TrafficLight) {
	fmt.Println("Yellow Light - Caution")
	// 转换到下一个状态
	context.SetState(&RedState{})
}

func (ys *YellowState) String() string {
	return "Yellow Light"
}

// TrafficLight 是包含状态的上下文
type TrafficLight struct {
	state State
}

// SetState 改变上下文的状态
func (tl *TrafficLight) SetState(state State) {
	tl.state = state
}

// Request 处理来自外部的请求
func (tl *TrafficLight) Request() {
	tl.state.Handle(tl)
}

// GetCurrentState 返回当前状态
func (tl *TrafficLight) GetCurrentState() State {
	return tl.state
}

func main() {
	light := &TrafficLight{
		state: &RedState{}, // 初始状态设置为红灯
	}

	// 演示状态的更改
	for i := 0; i < 6; i++ {
		fmt.Printf("The light is now %s\n", light.GetCurrentState())
		light.Request()
	}
	/**

	The light is now Red Light
	Red Light - Stop

	The light is now Green Light
	Green Light - Go

	The light is now Yellow Light
	Yellow Light - Caution


	The light is now Red Light
	Red Light - Stop

	The light is now Green Light
	Green Light - Go

	The light is now Yellow Light
	Yellow Light - Caution

	*/
}

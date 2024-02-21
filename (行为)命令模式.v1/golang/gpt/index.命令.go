package main

import "fmt"

// Command 接口声明执行命令的方法
type Command interface {
	Execute()
}

// Receiver 类，具体执行命令的类
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver: 执行操作")
}

// ConcreteCommand 实现 Command 接口，调用接收者的方法
type ConcreteCommand struct {
	receiver *Receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Invoker 类，请求者
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) PressButton() {
	i.command.Execute()
}

/**
命令模式通常涉及以下几个组成部分：

	Command 接口：定义执行操作的方法。
	ConcreteCommand 类：实现 Command 接口，并定义接收者的动作和参数。
	Receiver 类：执行与请求相关的操作；即执行具体的命令。
	Invoker 类：请求命令执行的调用者。
	Client 类：创建具体命令对象，并设置其接收者。

*/
func main() {
	receiver := &Receiver{}
	command := &ConcreteCommand{
		receiver: receiver,
	}
	invoker := &Invoker{}
	invoker.SetCommand(command)

	// 模拟按下按钮触发命令
	invoker.PressButton()
}

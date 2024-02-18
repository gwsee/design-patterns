package main

import "fmt"

/**
中介者模式（Mediator Pattern）是一种行为设计模式，它允许你将系统中对象之间的通信封装到一个中介对象中。
对象通过中介者互相交互，从而减少相互之间的直接引用。这有助于减少对象之间的耦合，并且可以更容易地改变这些对象之间的交互。

在Go语言中实现中介者模式通常涉及以下几个部分：

Mediator 接口：定义用于沟通 Colleague 对象的方法。
ConcreteMediator 类：实现 Mediator 接口，并协调几个 Colleague 对象之间的交互。
Colleague 类：一个或多个类，它们了解其 Mediator 对象。它们通过 Mediator 进行通信，而不是直接与其他 Colleague 交互
*/

// Mediator 定义中介者接口
type Mediator interface {
	Notify(sender Colleague, event string)
}

// Colleague 定义同事类接口
type Colleague interface {
	SetMediator(mediator Mediator)
}

// ConcreteMediator 实现中介者，并维护同事对象的引用
type ConcreteMediator struct {
	colleague1 *ConcreteColleague1
	colleague2 *ConcreteColleague2
}

// Notify 实现中介者的消息通知方法
func (m *ConcreteMediator) Notify(sender Colleague, event string) {
	switch sender {
	case m.colleague1:
		// 当 Colleague1 发生变化时，Mediator 通知 Colleague2
		fmt.Println("Mediator reacts on Colleague1 and triggers following operations:")
		m.colleague2.RespondToEvent()
	case m.colleague2:
		// 当 Colleague2 发生变化时，Mediator通知 Colleague1
		fmt.Println("Mediator reacts on Colleague2 and triggers following operations:")
		m.colleague1.Action()
	default:
		fmt.Println("Unknown colleague, no action taken by Mediator")
	}
}

// ConcreteColleague1 实现同事类接口
type ConcreteColleague1 struct {
	mediator Mediator
}

// SetMediator 设置中介者
func (c *ConcreteColleague1) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

// Action 执行操作时通知中介者
func (c *ConcreteColleague1) Action() {
	fmt.Println("Colleague1 takes action and notifies Mediator")
	c.mediator.Notify(c, "Colleague1")
}

// ConcreteColleague2 实现同事类接口
type ConcreteColleague2 struct {
	mediator Mediator
}

// SetMediator 设置中介者
func (c *ConcreteColleague2) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

// RespondToEvent 响应来自中介者的通知
func (c *ConcreteColleague2) RespondToEvent() {
	fmt.Println("Colleague2 responds to notification")
}

func main() {
	// 实例化中介者
	mediator := &ConcreteMediator{}

	// 初始化同事类并设置其中介者
	colleague1 := &ConcreteColleague1{}
	colleague2 := &ConcreteColleague2{}
	colleague1.SetMediator(mediator)
	colleague2.SetMediator(mediator)
	mediator.colleague1 = colleague1
	mediator.colleague2 = colleague2

	// 触发行为，行为通过中介者相互通知
	colleague1.Action()
	colleague2.RespondToEvent()
}

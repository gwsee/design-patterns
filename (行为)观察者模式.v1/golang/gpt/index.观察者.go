package main

import "fmt"

// Observer 定义观察者需要实现的接口
type Observer interface {
	Update(subject *Subject)
}

// Subject 保存观察者，并提供注册和删除观察者的接口
type Subject struct {
	observers []Observer
	state     string
}

func (s *Subject) RegisterObserver(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) RemoveObserver(o Observer) {
	var indexToRemove int
	for i, observer := range s.observers {
		if observer == o {
			indexToRemove = i
			break
		}
	}
	s.observers = append(s.observers[:indexToRemove], s.observers[indexToRemove+1:]...)
}

func (s *Subject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s)
	}
}

func (s *Subject) SetState(state string) {
	s.state = state
	s.NotifyObservers()
}

func (s *Subject) GetState() string {
	return s.state
}

// ConcreteObserver 实现了 Observer 接口
type ConcreteObserver struct {
	id int
}

func (co *ConcreteObserver) Update(subject *Subject) {
	fmt.Printf("观察者 %d: 接收到状态变化 -> %s\n", co.id, subject.GetState())
}

/**
在 Go 语言中，观察者模式通常包括以下几个组件：

	Subject 接口：该接口定义注册、删除和通知观察者的方法。
	ConcreteSubject 类：实现 Subject 接口的类，维护观察者列表，并在状态改变时通知观察者。
	Observer 接口：定义观察者需要实现的 Update 方法，用来接收 Subject 发出的通知。
	ConcreteObserver 类：实现 Observer 接口的类，定义在收到通知时要采取的行动。
*/
func main() {
	// 创建 Subject（被观察者）
	subject := &Subject{}

	// 创建观察者
	observer1 := &ConcreteObserver{id: 1}
	observer2 := &ConcreteObserver{id: 2}

	// 注册观察者
	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	// 改变 Subject 状态，观察者应该会收到通知
	subject.SetState("active")

	// 移除一个观察者
	subject.RemoveObserver(observer1)

	// 再次改变状态，移除的观察者不再接收通知
	subject.SetState("inactive")

	/**
	观察者 1: 接收到状态变化 -> active
	观察者 2: 接收到状态变化 -> active
	观察者 2: 接收到状态变化 -> inactive
	*/
}

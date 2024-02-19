package main

import "fmt"

// Memento 代表存储原发器状态的备忘录对象
type Memento struct {
	state string
}

// NewMemento 创建一个备忘录对象
func NewMemento(state string) *Memento {
	return &Memento{state: state}
}

// GetState 用于获取备忘录保存的状态
func (m *Memento) GetState() string {
	return m.state
}

// Originator 是原发器类，它有一个需要保存的状态
type Originator struct {
	state string
}

// SetState 设置原发器的状态
func (o *Originator) SetState(state string) {
	o.state = state
}

// SaveToMemento 保存当前状态到备忘录
func (o *Originator) SaveToMemento() *Memento {
	return NewMemento(o.state)
}

// RestoreFromMemento 从备忘录恢复状态
func (o *Originator) RestoreFromMemento(m *Memento) {
	o.state = m.GetState()
}

// Caretaker 代表看护者，它维护备忘录的列表，但不修改备忘录
type Caretaker struct {
	mementoList []*Memento
}

// NewCaretaker 创建一个看护者实例
func NewCaretaker() *Caretaker {
	return &Caretaker{}
}

// Add 添加备忘录至列表
func (c *Caretaker) Add(memento *Memento) {
	c.mementoList = append(c.mementoList, memento)
}

// Get 获取之前保存的备忘录
func (c *Caretaker) Get(index int) *Memento {
	if index < len(c.mementoList) && index >= 0 {
		return c.mementoList[index]
	}
	return nil
}

/**
备忘录模式（Memento Pattern）是一种行为设计模式，
它允许在不暴露对象内部状态的情况下捕获并保存当前时刻的状态，
以便可以在以后的时间恢复到这个状态。
备忘录模式通常用于实现撤销（undo）或恢复（redo）功能。

在Go语言中实现备忘录模式通常涉及以下几个角色：

Originator：原发器，它是需要保存状态的对象。
Memento：备忘录，用于存储原发器的内部状态。
Caretaker：看护者，负责保存和恢复原发器状态的对象。
*/
func main() {
	originator := &Originator{}
	caretaker := NewCaretaker()

	originator.SetState("State #1")
	originator.SetState("State #2")
	caretaker.Add(originator.SaveToMemento())

	originator.SetState("State #3")
	caretaker.Add(originator.SaveToMemento())

	originator.SetState("State #4")
	fmt.Println("Current State: ", originator.state)

	savedState := caretaker.Get(0)
	fmt.Println("First saved State: ", savedState.GetState())
	originator.RestoreFromMemento(savedState)

	fmt.Println("State after restoring: ", originator.state)
}

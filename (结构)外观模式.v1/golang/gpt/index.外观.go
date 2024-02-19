package main

import "fmt"

// SubsystemA defines an interface of a subsystem.
type SubsystemA interface {
	OperationA1() string
	OperationA2() string
}

// ConcreteSubsystemA is a concrete implementation of a subsystem.
type ConcreteSubsystemA struct{}

func (a *ConcreteSubsystemA) OperationA1() string {
	return "Subsystem A, Method A1\n"
}

func (a *ConcreteSubsystemA) OperationA2() string {
	return "Subsystem A, Method A2\n"
}

// SubsystemB defines another interface of a subsystem.
type SubsystemB interface {
	OperationB1() string
	OperationB2() string
}

// ConcreteSubsystemB is a concrete implementation of a subsystem.
type ConcreteSubsystemB struct{}

func (b *ConcreteSubsystemB) OperationB1() string {
	return "Subsystem B, Method B1\n"
}

func (b *ConcreteSubsystemB) OperationB2() string {
	return "Subsystem B, Method B2\n"
}

// Facade defines a high-level interface that
// makes the subsystems easier to use.
type Facade struct {
	subsystemA SubsystemA
	subsystemB SubsystemB
}

func NewFacade(subsystemA SubsystemA, subsystemB SubsystemB) *Facade {
	return &Facade{
		subsystemA: subsystemA,
		subsystemB: subsystemB,
	}
}

func (f *Facade) Operation() string {
	result := "Facade initializes subsystems:\n"
	result += f.subsystemA.OperationA1()
	result += f.subsystemB.OperationB1()
	result += "Facade orders subsystems to perform the operation:\n"
	result += f.subsystemA.OperationA2()
	result += f.subsystemB.OperationB2()
	return result
}

/**
在Go语言中实现设计模式时，并不强调传统意义上的类和继承，因为Go是基于接口的语言。
	外观模式（Facade Pattern）在Go中也不例外，它提供了一个统一的接口，用来访问子系统中的一群接口。
	外观定义了一个更高层次的接口，使得子系统更易于使用。

外观模式的目的是隐藏系统的复杂性，并提供一个客户端可以访问这个系统的简化接口。
	这个简化的接口使得客户端使用起来更加容易，同时也耦合了客户端对系统的依赖。
*/
func main() {
	subsystemA := &ConcreteSubsystemA{}
	subsystemB := &ConcreteSubsystemB{}
	facade := NewFacade(subsystemA, subsystemB)

	fmt.Print(facade.Operation())
}

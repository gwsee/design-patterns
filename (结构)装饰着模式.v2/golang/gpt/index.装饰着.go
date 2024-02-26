package main

import "fmt"

// Component 接口定义了业务逻辑的方法
type Component interface {
	Operation() string
}

// ConcreteComponent 是原始的、未装饰的组件
type ConcreteComponent struct{}

// Operation 是 ConcreteComponent 的方法，返回基本操作的结果
func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent Operation"
}

// Decorator 接口嵌入了 Component，允许我们构建一个可以递归嵌套的装饰器链
type Decorator interface {
	Component
	Decorate(Component) Component
}

// ConcreteDecoratorA 是一个装饰器，它实现了 Decorator 接口
type ConcreteDecoratorA struct {
	component Component
}

// Decorate 方法用来将一个 Component 设置为 Decorator 的内部属性，继而可以装饰它
func (d *ConcreteDecoratorA) Decorate(component Component) Component {
	d.component = component
	return d
}

// Operation 方法在执行基础组件的操作基础上增加了附加的行为
func (d *ConcreteDecoratorA) Operation() string {
	return "<DecoratedA>" + d.component.Operation() + "</DecoratedA>"
}

// ConcreteDecoratorB 是另一个装饰器，作用同上
type ConcreteDecoratorB struct {
	component Component
}

func (d *ConcreteDecoratorB) Decorate(component Component) Component {
	d.component = component
	return d
}

func (d *ConcreteDecoratorB) Operation() string {
	return "<DecoratedB>" + d.component.Operation() + "</DecoratedB>"
}

/**
定义一个接口，这将是你想要装饰的对象类型。
实现了这个接口的一个或多个具体对象（Concrete Component）。
创建一个装饰器结构体（Decorator），它也实现了这个接口，并包含了一个该接口类型的字段。
在装饰器中，提供方法来增强或修改具体对象的行为。
*/

func main() {
	// 创建基础组件
	var component Component = &ConcreteComponent{}

	// 装饰组件
	decoratedComponentA := &ConcreteDecoratorA{}
	decoratedComponentA.Decorate(component)

	decoratedComponentB := &ConcreteDecoratorB{}
	decoratedComponentB.Decorate(decoratedComponentA)

	// 最终操作输出
	fmt.Println(decoratedComponentB.Operation())
	// 输出: <DecoratedB><DecoratedA>ConcreteComponent Operation</DecoratedA></DecoratedB>
}

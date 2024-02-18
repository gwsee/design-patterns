package main

import "fmt"

/**
生成器模式（Builder Pattern）在 Go 语言中是一种用于构建复杂对象的设计模式，
	它允许你逐步创建对象的不同部分，并且可以为创建过程提供清晰的步骤和流程。
	生成器模式主要用于减少构造函数参数的数量，并避免构造函数中出现大量参数所导致的可读性和可维护性问题。

在 Go 中实现生成器模式通常涉及以下几个部分：

产品（Product）：是最终要构建的复杂对象。
生成器接口（Builder Interface）：定义了用于构建产品各个部件的方法。
具体生成器（Concrete Builder）：实现了生成器接口，并保留了所构建产品的状态。
指挥者（Director）：负责管理构建流程（可选）。
*/

// Product 表示最终要创建的复杂对象
type Product struct {
	partA string
	partB string
	partC string
}

// Builder 定义构建产品所需的方法
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetResult() Product
}

// ConcreteBuilder 实现了 Builder 接口，并构造和装配不同的产品部分
type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.partA = "PartA"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.partB = "PartB"
}

func (b *ConcreteBuilder) BuildPartC() {
	b.product.partC = "PartC"
}

func (b *ConcreteBuilder) GetResult() Product {
	return b.product
}

// Directors 负责构建算法，可以创建不同的产品变种
func Directors(builder Builder) {
	builder.BuildPartA()
	builder.BuildPartB()
	builder.BuildPartC()
}

func main() {
	builder := &ConcreteBuilder{}
	Directors(builder)             // 构建过程由 Director 指挥
	product := builder.GetResult() // 获取构建的产品

	fmt.Printf("Product parts: %s, %s, %s\n", product.partA, product.partB, product.partC)
}

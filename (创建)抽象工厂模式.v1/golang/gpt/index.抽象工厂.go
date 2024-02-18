package main

import "fmt"

/**
抽象工厂模式是一种创建型设计模式，它提供了创建一系列相关或依赖对象的接口，而不需要指定它们具体的类。
具体来说，抽象工厂定义了多个工厂方法，每一个工厂方法用于创建一种类型的对象。

以下是 Go 语言中实现抽象工厂模式的一个例子：
*/

// 抽象产品 A 和 B 的接口

type ProductA interface {
	UseA() string
}

type ProductB interface {
	UseB() string
}

// 具体产品
type ConcreteProductA1 struct{}

func (p *ConcreteProductA1) UseA() string { return "ProductA1" }

type ConcreteProductA2 struct{}

func (p *ConcreteProductA2) UseA() string { return "ProductA2" }

type ConcreteProductB1 struct{}

func (p *ConcreteProductB1) UseB() string { return "ProductB1" }

type ConcreteProductB2 struct{}

func (p *ConcreteProductB2) UseB() string { return "ProductB2" }

// 抽象工厂接口
type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

// 具体工厂实现
type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProductA() ProductA { return &ConcreteProductA1{} }
func (f *ConcreteFactory1) CreateProductB() ProductB { return &ConcreteProductB1{} }

type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProductA() ProductA { return &ConcreteProductA2{} }
func (f *ConcreteFactory2) CreateProductB() ProductB { return &ConcreteProductB2{} }

// 客户端代码
func main() {
	var factory AbstractFactory
	var productA ProductA
	var productB ProductB

	// 使用工厂1创建产品 A1 和 B1
	factory = &ConcreteFactory1{}
	productA = factory.CreateProductA()
	productB = factory.CreateProductB()

	fmt.Println(productA.UseA()) // Output: ProductA1
	fmt.Println(productB.UseB()) // Output: ProductB1

	// 使用工厂2创建产品 A2 和 B2
	factory = &ConcreteFactory2{}
	productA = factory.CreateProductA()
	productB = factory.CreateProductB()

	fmt.Println(productA.UseA()) // Output: ProductA2
	fmt.Println(productB.UseB()) // Output: ProductB2
}

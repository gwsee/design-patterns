package main

import "fmt"

// IProduct 定义一个产品的接口
type IProduct interface {
	Use() string
}

// ConcreteProduct1 具体产品1
type ConcreteProduct1 struct{}

func (p *ConcreteProduct1) Use() string {
	return "Product 1 being used"
}

// ConcreteProduct2 具体产品2
type ConcreteProduct2 struct{}

func (p *ConcreteProduct2) Use() string {
	return "Product 2 being used"
}

// IFactory 定义一个工厂的接口
type IFactory interface {
	CreateProduct() IProduct
}

// ConcreteFactory1 具体工厂1，生产产品1
type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProduct() IProduct {
	return &ConcreteProduct1{}
}

// ConcreteFactory2 具体工厂2，生产产品2
type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProduct() IProduct {
	return &ConcreteProduct2{}
}

func main() {
	var factory IFactory

	// 使用工厂1创建产品1
	factory = &ConcreteFactory1{}
	product1 := factory.CreateProduct()
	fmt.Println(product1.Use())

	// 使用工厂2创建产品2
	factory = &ConcreteFactory2{}
	product2 := factory.CreateProduct()
	fmt.Println(product2.Use())
}

package main

import (
	"fmt"
)

// Product 定义一个产品的接口
type Product interface {
	Use() string
}

// ConcreteProductA 具体产品A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Product A being used"
}

// ConcreteProductB 具体产品B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Product B being used"
}

// Factory 是一个简单工厂，可以根据类型创建产品
type Factory struct{}

func (f *Factory) CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}

func main() {
	factory := &Factory{}

	productA := factory.CreateProduct("A")
	fmt.Println(productA.Use())

	productB := factory.CreateProduct("B")
	fmt.Println(productB.Use())
}

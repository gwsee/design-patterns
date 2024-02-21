package main

import "fmt"

// 实现者接口
type Drawer interface {
	DrawCircle(radius int)
}

// 具体实现A
type RedCircleDrawer struct{}

func (r RedCircleDrawer) DrawCircle(radius int) {
	fmt.Println("Drawing circle of radius", radius, "in red color")
}

// 具体实现B
type BlueCircleDrawer struct{}

func (b BlueCircleDrawer) DrawCircle(radius int) {
	fmt.Println("Drawing circle of radius", radius, "in blue color")
}

// 抽象部分
type Shape struct {
	drawer Drawer
}

func NewShape(drawer Drawer) *Shape {
	return &Shape{
		drawer: drawer,
	}
}

// 抽象部分的行为（功能）
func (s *Shape) Draw(radius int) {
	s.drawer.DrawCircle(radius)
}

// 扩展抽象
type Circle struct {
	Shape
}

// 可以添加额外的方法或字段
func (c *Circle) Draw() {
	c.Shape.Draw(5) // 使用抽象部分的 Draw 方法
}

/**
Go语言（Golang）中的桥接模式（Bridge Pattern）是一种结构型设计模式，
	用于将抽象与其实现解耦，使得两者可以独立变化。
	这在Go语言中意味着通过定义接口来分离抽象和具体实现，并允许它们独立进行扩展。
	由于Go语言没有类和继承，接口和组合成为实现桥接模式的主要手段。

桥接模式的结构
桥接模式通常涉及以下几个部分：
	抽象部分（Abstraction）：定义高层控制逻辑，它会引用实现部分的接口。
	扩展抽象（Refined Abstraction）：扩展抽象部分，增加了一些额外的功能和逻辑。
	实现者接口（Implementor）：定义实现类的接口，通常是一个或一组方法声明。
	具体实现（Concrete Implementations）：实现者接口的具体类，这些类将提供具体的方法实现。
*/
func main() {
	redCircle := NewShape(RedCircleDrawer{})
	blueCircle := NewShape(BlueCircleDrawer{})

	redCircle.Draw(10)
	blueCircle.Draw(20)

	circle := Circle{Shape{RedCircleDrawer{}}}
	circle.Draw()
	/**
	Drawing circle of radius 10 in red color
	Drawing circle of radius 20 in blue color
	Drawing circle of radius 5 in red color
	*/
}

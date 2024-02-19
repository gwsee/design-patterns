package main

import "fmt"

// Component 提供了组合中对象的共享接口，既可以代表单个对象也可以代表组合对象
type Component interface {
	Perform() string
}

// Leaf 实现了 Component 接口，并代表组合中的单个对象
type Leaf struct {
	name string
}

func (l *Leaf) Perform() string {
	return l.name
}

// Composite 实现了 Component 接口，并代表一个由子部件构成的组合对象
type Composite struct {
	children []Component
	name     string
}

func (c *Composite) Add(child Component) {
	c.children = append(c.children, child)
}

func (c *Composite) Perform() string {
	result := fmt.Sprintf("Composite [%s]\n", c.name)
	for _, child := range c.children {
		result += child.Perform() + "\n"
	}

	// 移除末尾的换行符
	result = result[:len(result)-1]
	return result
}

/**
在Go语言（又称为Golang）中，组合模式不是通过继承实现的，
而是通过嵌入结构体（embedding struct）的方式来实现的。
组合模式是一种结构型设计模式，它可以让你将对象组合成树形结构以表示部分-整体的层次结构。
组合模式通过让单个对象和组合对象共享相同的接口来实现这一点。
*/
func main() {
	// 创建树形结构
	root := &Composite{name: "Root"}

	leaf1 := &Leaf{name: "Leaf 1"}
	leaf2 := &Leaf{name: "Leaf 2"}

	composite1 := &Composite{name: "Composite 1"}
	composite2 := &Composite{name: "Composite 2"}

	root.Add(leaf1)
	root.Add(composite1)

	composite1.Add(leaf2)
	composite1.Add(composite2)

	composite2.Add(&Leaf{name: "Leaf 3"})

	// 执行操作
	fmt.Println(root.Perform())
}

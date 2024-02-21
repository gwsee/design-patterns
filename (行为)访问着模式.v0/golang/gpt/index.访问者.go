package main

import "fmt"

// Visitor 接口
type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA)
	VisitConcreteElementB(*ConcreteElementB)
}

// Element 接口
type Element interface {
	Accept(Visitor)
}

// ConcreteElementA 结构体
type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(v Visitor) {
	v.VisitConcreteElementA(e)
}

// ConcreteElementB 结构体
type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(v Visitor) {
	v.VisitConcreteElementB(e)
}

// ConcreteVisitor 结构体
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Println("访问者访问元素A")
}

func (v *ConcreteVisitor) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Println("访问者访问元素B")
}

// ObjectStructure 结构体，可以遍历 Element
type ObjectStructure struct {
	elements []Element
}

func (o *ObjectStructure) Add(element Element) {
	o.elements = append(o.elements, element)
}

func (o *ObjectStructure) Accept(v Visitor) {
	for _, e := range o.elements {
		e.Accept(v)
	}
}

/**
Visitor 接口：定义一个访问方法，该方法为每一种类型的被访问元素提供了一个访问的方法签名。
Element 接口：定义一个 Accept 方法，该方法应接收一个访问者对象。
ConcreteVisitor 类：实现 Visitor 接口的类，里面包含了对每一个元素类的访问逻辑。
ConcreteElement 类：实现 Element 接口的类，代表被访问的对象。
*/
func main() {
	// 创建ObjectStructure
	objects := &ObjectStructure{}

	// 向ObjectStructure添加元素
	objects.Add(&ConcreteElementA{})
	objects.Add(&ConcreteElementB{})

	// 创建访问者
	visitor := &ConcreteVisitor{}

	// 使用访问者访问对象结构中的元素
	objects.Accept(visitor)
	/**
	访问者访问元素A
	访问者访问元素B
	*/
}

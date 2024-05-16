package main

import "fmt"

/**

在 Go 语言中，原型模式（Prototype Pattern）是一种创建型设计模式，
	它通过复制已存在的实例来创建新的实例，而不是通过新建的方式。
	这种方式特别适合于创建成本较高的实例时，
	例如对象的初始化需要消耗较多资源，或者某个实例的创建需要依赖于外部服务等。

Go 语言中并没有内置的克隆（clone）接口。
	不过，可以通过定义自己的克隆接口并实现该接口来达成原型模式的效果。

由于 Go 不支持类似于 Java 中 clone() 方法的语言级原型克隆，通常需要手动实现克隆的逻辑。
	如果需要更深层次的克隆（例如嵌套的结构体或包含指针的字段），可能需要实现更复杂的克隆逻辑以确保所有层次的数据都被正确地复制。
	在某些情况下，你可能需要使用序列化和反序列化的技术来实现一个深克隆（deep clone）功能。

使用原型模式可以避免每次创建对象时都需要经历复杂的初始化流程，
	也可以用来避免重复创建具有相同配置的对象，从而提升程序的效率。
*/

// Cloner 是一个克隆接口，任何实现了该接口的类型都能进行克隆
type Cloner interface {
	Clone() Cloner
}

// ConcretePrototype 是一个具体的原型类型
type ConcretePrototype struct {
	field string
}

// Clone 方法克隆出一个新的 ConcretePrototype 实例
func (c *ConcretePrototype) Clone() Cloner {
	// 创建一个新的 ConcretePrototype
	// 注意这里使用了 c.field 来保持字段值
	return &ConcretePrototype{field: c.field}
}

func main() {
	// 创建一个原型对象
	proto := &ConcretePrototype{field: "prototype value"}

	// 克隆原型对象
	clone := proto.Clone()

	// 类型断言获取具体的克隆对象
	cloneProto, ok := clone.(*ConcretePrototype)
	if !ok {
		fmt.Println("Clone did not return a *ConcretePrototype")
		return
	}

	fmt.Println(proto.field)         // 输出: prototype value
	fmt.Println(cloneProto.field)    // 输出: prototype value
	fmt.Println(proto == cloneProto) // 输出: false
}

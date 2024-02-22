package main

import "fmt"

type Flyweight interface {
	Operation(extrinsicState string)
}

type ConcreteFlyweight struct {
	intrinsicState string // 内部状态
}

func (c *ConcreteFlyweight) Operation(extrinsicState string) {
	fmt.Printf("ConcreteFlyweight: Intrinsic State %s, "+
		"Extrinsic State %s\n", c.intrinsicState, extrinsicState)
}

type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}
}

func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := f.flyweights[key]; ok {
		return flyweight
	}
	flyweight := &ConcreteFlyweight{intrinsicState: key}
	f.flyweights[key] = flyweight
	return flyweight
}

func main() {
	factory := NewFlyweightFactory()

	flyweight1 := factory.GetFlyweight("foo")
	flyweight1.Operation("bar1")

	flyweight2 := factory.GetFlyweight("foo") // 它应该返回相同的对象
	flyweight2.Operation("bar2")

	fmt.Printf("flyweight1 == flyweight2: %v\n", flyweight1 == flyweight2) // 输出 true, 显示它们是相同的实例
}

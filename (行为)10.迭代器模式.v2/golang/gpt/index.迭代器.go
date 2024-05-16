package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
type Collection struct {
	items []interface{}
}

func (c *Collection) CreateIterator() Iterator {
	return &ConcreteIterator{
		collection: c,
		index:      0,
	}
}

type ConcreteIterator struct {
	collection *Collection
	index      int
}

func (it *ConcreteIterator) HasNext() bool {
	return it.index < len(it.collection.items)
}

func (it *ConcreteIterator) Next() interface{} {
	if it.HasNext() {
		item := it.collection.items[it.index]
		it.index++
		return item
	}
	return nil // 或者返回错误，如果需要的话
}

/**
这个例子中，Collection类型代表聚合对象，它有一组项和一个创建迭代器的方法。
ConcreteIterator实现了Iterator接口，并可以遍历Collection。

当客户端代码运行时，它通过Collection的CreateIterator方法获取一个迭代器，
然后使用HasNext和Next方法来遍历集合中的元素。这样，客户端代码无需知道集合的内部结构，也不需要知道它是如何实现遍历的。

迭代器模式的主要优点是它支持不同的遍历策略，同时使聚合对象的结构和遍历操作相互独立，
这有助于减少聚合对象与迭代逻辑之间的耦合。在Go语言标准库中，迭代器模式的概念广泛应用于各种数据结构，比如bufio.Scanner、filepath.Walk等。
*/
func main() {
	items := []interface{}{"a", "b", "c", "d"}
	collection := Collection{items: items}
	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Println(item)
	}
}

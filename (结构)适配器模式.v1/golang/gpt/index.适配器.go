package main

import "fmt"

/**
适配器模式（Adapter Pattern）
	是一种结构型设计模式，它允许不兼容的对象能够相互协作。
	这种模式主要用于确保那些因接口不兼容而不能一起工作的类可以一起工作。
	适配器的作用就是解决接口不兼容的问题，让原本由于接口不兼容而不能一起工作的类可以一起工作。

在Go语言中，适配器模式通常涉及以下几个部分：

目标接口（Target Interface）: 定义客户端使用的与特定领域相关的接口。
要适配的类（Adaptee）: 一个已经存在的接口，这个接口需要适配。
适配器（Adapter）: 对Adaptee的接口与Target接口进行适配。
*/

// Target 是客户端期望使用的接口
type Target interface {
	Request() string
}

// Adaptee 是一个已经存在的接口，这里模拟一个旧的充电器接口
type Adaptee struct{}

// SpecificRequest 是Adapter的一个方法
// 注意：此方法与Target接口不兼容，因为返回的字符串不同
func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Adapter 是将Adaptee适配成Target接口的适配器
type Adapter struct {
	adapter *Adaptee
}

// Request 方法让Adapter实现了Target接口
func (a *Adapter) Request() string {
	// 适配器可能转换数据格式、改变请求的方式等等，以匹配目标接口期待的形式
	return "Adapter: (TRANSLATED) " + a.adapter.SpecificRequest()
}

// Client 客户端代码只知道如何使用实现Target接口的对象
func Client(target Target) {
	fmt.Println(target.Request())
}

func main() {
	// 旧的实现不兼容新的接口
	adapt := &Adaptee{}
	fmt.Println("Adapter: ", adapt.SpecificRequest())

	// 创建一个适配器，并传入一个具体的Adaptee对象
	adapter := &Adapter{
		adapter: adapt,
	}

	// 现在可以使用Adapter来调用Adaptee的方法，并且其形式是兼容Target接口的
	Client(adapter)
}

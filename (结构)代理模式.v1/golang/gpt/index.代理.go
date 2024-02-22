package main

import "fmt"

type Subject interface {
	DoAction()
}

type RealSubject struct{}

func (rs *RealSubject) DoAction() {
	fmt.Println("RealSubject: Doing action")
}

type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) DoAction() {
	if p.checkAccess() {
		p.realSubject.DoAction()
		p.logAccess()
	}
}

// 预留的钩子方法, 在这里可以进行权限检查
func (p *Proxy) checkAccess() bool {
	fmt.Println("Proxy: Checking access before firing a real request.")
	// 这里应该包含一些真实的检查逻辑
	// 如果合适，返回true
	return true
}

// 可以记录每次的请求
func (p *Proxy) logAccess() {
	fmt.Println("Proxy: Logging the time of request.")
}

/**
在这个例子中，main函数创建了一个真实对象realSubject，
	然后创建了代理对象proxy并传入了realSubject作为其依赖。
	当调用proxy.DoAction()时，proxy会先执行权限检查和日志记录，然后调用真实对象的DoAction()方法。

使用代理模式有许多好处，你可以不修改客户代码的基础上增加或修改对象的行为，
	可以非常方便地实现懒加载、权限检查、日志记录等功能，并且能够更好地遵循开闭原则。
*/
func main() {
	realSubject := &RealSubject{}
	proxy := Proxy{realSubject: realSubject}

	// 客户端使用代理，而不是直接调用真实对象
	proxy.DoAction()
	/**
	Proxy: Checking access before firing a real request.
	RealSubject: Doing action
	Proxy: Logging the time of request.
	*/
}

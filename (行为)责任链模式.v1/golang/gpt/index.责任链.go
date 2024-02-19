package main

import "fmt"

// Handler 接口定义了一个处理请求的方法和设置下一个处理者的方法。
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string) string
}

// BaseHandler 提供了处理者在责任链中的基础结构。
type BaseHandler struct {
	next Handler
}

// SetNext 设置责任链中的下一个处理者。
func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

// Handle 默认实现无操作，如果有下一个处理者则传递给下一个处理者。
func (h *BaseHandler) Handle(request string) string {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return ""
}

// ConcreteHandlerA 是责任链的一个具体处理者。
type ConcreteHandlerA struct {
	BaseHandler
}

// Handle 方法在这里实现具体逻辑。
func (a *ConcreteHandlerA) Handle(request string) string {
	if request == "A" {
		// 如果可以处理该请求，完成处理并返回。
		return fmt.Sprintf("ConcreteHandlerA handled request '%s'", request)
	}
	// 否则传递给下一个处理者。
	return a.BaseHandler.Handle(request)
}

// ConcreteHandlerB 是责任链上的另一个处理者。
type ConcreteHandlerB struct {
	BaseHandler
}

// Handle 方法在这里实现具体逻辑。
func (b *ConcreteHandlerB) Handle(request string) string {
	if request == "B" {
		return fmt.Sprintf("ConcreteHandlerB handled request '%s'", request)
	}
	return b.BaseHandler.Handle(request)
}

func main() {
	// 构建责任链。
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	// 发送请求。
	fmt.Println(handlerA.Handle("A")) // 输出: ConcreteHandlerA handled request 'A'
	fmt.Println(handlerA.Handle("B")) // 输出: ConcreteHandlerB handled request 'B'
	fmt.Println(handlerA.Handle("C")) // 输出: (因为没有处理者可以处理"C", 它将被传递层层下去，最后没有输出)
}

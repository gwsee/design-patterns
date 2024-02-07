package main

import "fmt"

func main()  {
	// 定义策略
	addStrategy := &AddOperation{}
	subtractStrategy := &SubtractOperation{}
	multiplyStrategy := &MultiplyOperation{}

	// 创建上下文并注入不同策略
	context := NewContext(addStrategy)
	fmt.Println("10 + 5 =", context.ExecuteStrategy(10, 5))

	context = NewContext(subtractStrategy)
	fmt.Println("10 - 5 =", context.ExecuteStrategy(10, 5))

	context = NewContext(multiplyStrategy)
	fmt.Println("10 * 5 =", context.ExecuteStrategy(10, 5))
}

type Strategy interface {
	DoOperation(int, int) int
}
// AddOperation 实现加法策略
type AddOperation struct{}

func (a *AddOperation) DoOperation(num1, num2 int) int {
	return num1 + num2
}
// SubtractOperation 实现减法策略
type SubtractOperation struct{}

func (s *SubtractOperation) DoOperation(num1, num2 int) int {
	return num1 - num2
}

// MultiplyOperation 实现乘法策略
type MultiplyOperation struct{}

func (m *MultiplyOperation) DoOperation(num1, num2 int) int {
	return num1 * num2
}

// Context 是使用某种策略的客户端
type Context struct {
	strategy Strategy
}
func NewContext(strategy Strategy) *Context {
	return &Context{
		strategy: strategy,
	}
}

func (c *Context) ExecuteStrategy(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}
package main

import "fmt"

// IProcess 接口定义了模板方法的步骤
type IProcess interface {
	Step1()
	Step2()
	Execute()
}

// Template 结构体
type Template struct{}

// Step1 默认实现（可以被子类覆盖）
func (t *Template) Step1() {
	fmt.Println("步骤1的默认实现")
}

// Step2 默认实现（可以被子类覆盖）
func (t *Template) Step2() {
	fmt.Println("步骤2的默认实现")
}

// Execute 方法定义了算法的骨架
func (t *Template) Execute() {
	t.Step1()
	t.Step2()
}

// ConcreteProcess 结构体继承自 Template 结构体
type ConcreteProcess struct {
	*Template
}

// Step1 方法进行了覆盖，提供了具体的实现
func (c *ConcreteProcess) Step1() {
	fmt.Println("步骤1的具体实现")
}

// Execute 方法定义了算法的骨架
func (c *ConcreteProcess) Execute() {
	c.Step1()
	c.Step2()
}

//通常模板方法模式下，我们会直接组合或者嵌入模板结构体，而不会通过接口，
//这样可以直接继承模板结构体的方法。考虑下面的改进：
func main() {
	// 创建 ConcreteProcess 结构体的实例
	concrete := ConcreteProcess{
		&Template{}, // 初始化嵌入的结构体实例
	}

	// 调用 Execute 方法
	concrete.Execute()

	/**
	步骤1的具体实现
	步骤2的默认实现
	*/
}

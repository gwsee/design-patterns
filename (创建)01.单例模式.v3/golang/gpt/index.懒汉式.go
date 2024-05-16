package main

import (
	"fmt"
	"sync"
)

// Singleton2 是单例类型
type Singleton2 struct{}

// 懒汉式单例模式中的变量
var (
	instance2 *Singleton2
	once      sync.Once
)

// GetInstance2 使用 sync.Once 确保全局仅有一个 Singleton2 实例
func GetInstance2() *Singleton2 {
	once.Do(func() {
		// 这里的代码只执行一次
		instance2 = &Singleton2{}
	})
	return instance2
}

// main 函数是程序入口
func main() {
	singleton1 := GetInstance2()
	singleton2 := GetInstance2()

	fmt.Println(singleton1 == singleton2) // 输出: true
}

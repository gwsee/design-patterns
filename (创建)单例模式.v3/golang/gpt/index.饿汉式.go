package main

import "fmt"

// Singleton 是单例类型
type Singleton struct{}

// 单例实例
var instance *Singleton = &Singleton{}

// GetInstance 提供全局访问点获取单例
func GetInstance() *Singleton {
	return instance
}

// main 函数是程序入口
func main() {
	singleton1 := GetInstance()
	singleton2 := GetInstance()

	fmt.Println(singleton1 == singleton2) // 输出: true
}

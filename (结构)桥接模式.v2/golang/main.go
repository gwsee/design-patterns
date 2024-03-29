package main

import "fmt"

/**
假设你有两台电脑： 一台 Mac 和一台 Windows。
	还有两台打印机： 爱普生和惠普。
	这两台电脑和打印机可能会任意组合使用。
	客户端不应去担心如何将打印机连接至计算机的细节问题。

如果引入新的打印机， 我们也不会希望代码量成倍增长。
	所以， 我们创建了两个层次结构， 而不是 2x2 组合的四个结构体：

抽象层： 代表计算机
实施层： 代表打印机
这两个层次可通过桥接进行沟通， 其中抽象层 （计算机） 包含对于实施层 （打印机） 的引用。
	抽象层和实施层均可独立开发， 不会相互影响。
*/

func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}

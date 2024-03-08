package main

/**
确保至少有两个类的接口不兼容：

一个无法修改 （通常是第三方、 遗留系统或者存在众多已有依赖的类） 的功能性服务类。
一个或多个将受益于使用服务类的客户端类。
声明客户端接口， 描述客户端如何与服务交互。

创建遵循客户端接口的适配器类。 所有方法暂时都为空。

在适配器类中添加一个成员变量用于保存对于服务对象的引用。 通常情况下会通过构造函数对该成员变量进行初始化， 但有时在调用其方法时将该变量传递给适配器会更方便。

依次实现适配器类客户端接口的所有方法。 适配器会将实际工作委派给服务对象， 自身只负责接口或数据格式的转换。

客户端必须通过客户端接口使用适配器。 这样一来， 你就可以在不影响客户端代码的情况下修改或扩展适配器。
*/

/**
这里有一段客户端代码， 用于接收一个对象 （Lightning 接口） 的部分功能， 不过我们还有另一个名为 adaptee 的对象 （Windows 笔记本）， 可通过不同的接口 （USB 接口） 实现相同的功能

这就是适配器模式发挥作用的场景。 我们可以创建这样一个名为 adapter 的结构体：

遵循符合客户端期望的相同接口 （Lightning 接口）。

可以适合被适配对象的方式对来自客户端的请求进行 “翻译”。 适配器能够接受来自 Lightning 连接器的信息， 并将其转换成 USB 格式的信号， 同时将信号传递给 Windows 笔记本的 USB 接口。
*/

func main() {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}

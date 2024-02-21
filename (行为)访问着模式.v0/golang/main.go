package main

import "fmt"

/**
访问者模式允许你在结构体中添加行为， 而又不会对结构体造成实际变更。 假设你是一个代码库的维护者， 代码库中包含不同的形状结构体， 如：
	方形
	圆形
	三角形
上述每个形状结构体都实现了通用形状接口。

在公司员工开始使用你维护的代码库时， 你就会被各种功能请求给淹没。 让我们来看看其中比较简单的请求： 有个团队请求你在形状结构体中添加 get­Area获取面积行为。

解决这一问题的办法有很多。

第一个选项便是将 get­Area方法直接添加至形状接口， 然后在各个形状结构体中进行实现。 这似乎是比较好的解决方案， 但其代价也比较高。
	作为代码库的管理员， 相信你也不想在每次有人要求添加另外一种行为时就去冒着风险改动自己的宝贝代码。
	不过， 你也一定想让其他团队的人还是用一用自己的代码库。

第二个选项是请求功能的团队自行实现行为。 然而这并不总是可行， 因为行为可能会依赖于私有代码。

第三个方法就是使用访问者模式来解决上述问题。 首先定义一个如下访问者接口：
	type visitor interface {
		visitForSquare(square)
		visitForCircle(circle)
		visitForTriangle(triangle)
	}

我们可以使用 visit­For­Square­(square) 、 ​ visit­For­Circle­(circle)
	以及 visit­For­Triangle­(triangle)函数来为方形、 圆形以及三角形添加相应的功能。

你可能在想， 为什么我们不再访问者接口里面使用单一的 visit­(shape)方法呢？
	这是因为 Go 语言不支持方法重载， 所以你无法以相同名称、 不同参数的方式来使用方法。

好了， 第二项重要的工作是将 accept接受方法添加至形状接口中。

func accept(v visitor)
所有形状结构体都需要定义此方法， 类似于：

func (obj *square) accept(v visitor){
    v.visitForSquare(obj)
}
等等， 我刚才是不是提到过， 我们并不想修改现有的形状结构体？ 很不幸， 在使用访问者模式时， 我们必须要修改形状结构体。 但这样的修改只需要进行一次。

如果添加任何其他行为， 比如 get­Num­Sides获取边数和 get­Middle­Coordinates获取中点坐标 ，
	我们将使用相同的 accept­(v visitor)函数， 而无需对形状结构体进行进一步的修改。

最后， 形状结构体只需要修改一次， 并且所有未来针对不同行为的请求都可以使用相同的 accept 函数来进行处理。
	如果团队成员请求 get­Area行为， 我们只需简单地定义访问者接口的具体实现， 并在其中编写面积的计算逻辑即可。
*/
func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
	/**
	Calculating area for square
	Calculating area for circle
	Calculating area for rectangle

	Calculating middle point coordinates for square
	Calculating middle point coordinates for circle
	Calculating middle point coordinates for rectangle
	*/
}

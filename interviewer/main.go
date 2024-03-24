package main

import "fmt"

// 访问者模式

/*
访问者模式允许你在结构体中添加行为， 而又不会对结构体造成实际变更。 假设你是一个代码库的维护者， 代码库中包含不同的形状结构体， 如：

方形
圆形
三角形
上述每个形状结构体都实现了通用形状接口。

在公司员工开始使用你维护的代码库时， 你就会被各种功能请求给淹没。 让我们来看看其中比较简单的请求： 有个团队请求你在形状结构体中添加 getArea获取面积行为。

解决这一问题的办法有很多。

第一个选项便是将 getArea方法直接添加至形状接口， 然后在各个形状结构体中进行实现。 这似乎是比较好的解决方案， 但其代价也比较高。 作为代码库的管理员， 相信你也不想在每次有人要求添加另外一种行为时就去冒着风险改动自己的宝贝代码。 不过， 你也一定想让其他团队的人还是用一用自己的代码库。

第二个选项是请求功能的团队自行实现行为。 然而这并不总是可行， 因为行为可能会依赖于私有代码。

第三个方法就是使用访问者模式来解决上述问题。
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
}

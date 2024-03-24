package main

import (
	"fmt"
	"log"
)

/*
让我们在一台自动售货机上使用状态设计模式。 为简单起见， 让我们假设自动售货机仅会销售一种类型的商品。 同时， 依然为了简单起见， 我们假设自动售货机可处于 4 种不同的状态中：

	有商品 （hasItem）
	无商品 （noItem）
	商品已请求 （itemRequested）
	收到纸币 （hasMoney）
同时， 自动售货机也会有不同的操作。 再一次的， 为了简单起见， 我们假设其只会执行 4 种操作：

	选择商品
	添加商品
	插入纸币
	提供商品

当对象可以处于许多不同的状态中时应使用状态设计模式， 同时根据传入请求的不同， 对象需要变更其当前状态。

在我们的例子中， 自动售货机可以有多种不同的状态， 同时会在这些状态之间持续不断地互相转换。 我们假设自动售货机处于 商品已请求状态中。 在 “插入纸币” 的操作发生后， 机器将自动转换至 收到纸币状态。

根据其当前状态， 机器可就相同请求采取不同的行为。 例如， 如果用户想要购买一件商品， 机器将在 有商品状态时继续操作， 而在 无商品状态时拒绝操作。

自动售货机的代码不会被这一逻辑污染； 所有依赖于状态的代码都存在于各自的状态实现中。
*/

func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

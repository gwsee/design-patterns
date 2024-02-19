package main

import (
	"fmt"
	"log"
)

/**
人们很容易低估使用信用卡订购披萨时幕后工作的复杂程度。
	在整个过程中会有不少的子系统发挥作用。 下面是其中的一部分：
		检查账户
		检查安全码
		借记/贷记余额
		账簿录入
		发送消息通知
在如此复杂的系统中， 可以说是一步错步步错， 很容易就会引发大的问题。
	这就是为什么我们需要外观模式， 让客户端可以使用一个简单的接口来处理众多组件。
	客户端只需要输入卡片详情、 安全码、 支付金额以及操作类型即可。
	外观模式会与多种组件进一步地进行沟通， 而又不会向客户端暴露其内部的复杂性。
*/
func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}

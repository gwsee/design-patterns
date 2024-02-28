package main

/**
中介者模式的一个绝佳例子就是火车站交通系统。 两列火车互相之间从来不会就站台的空闲状态进行通信。
stationManager车站经理可充当中介者， 让平台仅可由一列入场火车使用，
而将其他火车放入队列中等待。 离场火车会向车站发送通知， 便于队列中的下一列火车进站。
*/

func main() {
	stationManager := newStationManger()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}

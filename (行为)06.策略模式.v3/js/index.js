// 定义策略对象
const strategies = {
    OperationA: function(data) {
        // 策略 A 的具体算法实现
        console.log('Performing Operation A on', data);
    },
    OperationB: function(data) {
        // 策略 B 的具体算法实现
        console.log('Performing Operation B on', data);
    },
    OperationC: function(data) {
        // 策略 C 的具体算法实现
        console.log('Performing Operation C on', data);
    }
};

// 定义上下文，负责使用策略对象
class Context {
    constructor(strategy) {
        this.strategy = strategy;
    }

    setStrategy(strategy) {
        this.strategy = strategy;
    }

    executeStrategy(data) {
        strategies[this.strategy](data);
    }
}

// 创建上下文对象，并使用策略
const context = new Context('OperationA');
context.executeStrategy('data'); // Output: Performing Operation A on data

// 更改策略并再次执行
context.setStrategy('OperationB');
context.executeStrategy('data'); // Output: Performing Operation B on data

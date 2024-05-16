// 实现层接口
class Implementor {
    operationImpl() {
        // 具体实现
    }
}

// 具体实现层A
class ConcreteImplementorA extends Implementor {
    operationImpl() {
        console.log('ConcreteImplementorA operation implementation.');
    }
}

// 具体实现层B
class ConcreteImplementorB extends Implementor {
    operationImpl() {
        console.log('ConcreteImplementorB operation implementation.');
    }
}

// 抽象层
class Abstraction {
    constructor(implementor) {
        this.implementor = implementor;
    }
    
    operation() {
        this.implementor.operationImpl();
    }
}

// 扩展抽象层
class RefinedAbstraction extends Abstraction {
    constructor(implementor) {
        super(implementor);
    }
    
    refinedOperation() {
        this.implementor.operationImpl();
        // 执行一些额外操作...
    }
}

// 客户端代码
const implementorA = new ConcreteImplementorA();
const abstractionA = new RefinedAbstraction(implementorA);
abstractionA.operation(); // 输出: "ConcreteImplementorA operation implementation."

// 更换实现方式，客户端代码不需要改变
const implementorB = new ConcreteImplementorB();
abstractionA.implementor = implementorB;
abstractionA.operation(); // 输出: "ConcreteImplementorB operation implementation."

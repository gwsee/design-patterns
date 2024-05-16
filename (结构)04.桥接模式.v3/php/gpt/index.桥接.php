<?php

namespace php\gpt\bridge;

// 实现接口定义了实现类的方法
interface Implementor {
    public function implementation();
}

// 具体的实现A
class ConcreteImplementorA implements Implementor {
    public function implementation() {
        echo "ConcreteImplementorA operation.\n";
    }
}

// 具体的实现B
class ConcreteImplementorB implements Implementor {
    public function implementation() {
        echo "ConcreteImplementorB operation.\n";
    }
}

// 抽象类，在高层次上定义操作，并持有实现类的引用
abstract class Abstraction {
    protected $implementor;

    public function __construct(Implementor $implementor) {
        $this->implementor = $implementor;
    }

    public function operation() {
        $this->implementor->implementation();
    }
}

// 扩展抽象类的改进版本
class RefinedAbstraction extends Abstraction {
    // 你可以在这里扩展抽象类的方法
    public function refinedOperation() {
        echo "Refined operation: ";
        $this->implementor->implementation();
    }
}

// 客户端代码
$implementorA = new ConcreteImplementorA();
$abstractionA = new RefinedAbstraction($implementorA);
$abstractionA->operation(); // 输出: ConcreteImplementorA operation.

$implementorB = new ConcreteImplementorB();
$abstractionB = new RefinedAbstraction($implementorB);
$abstractionB->operation(); // 输出: ConcreteImplementorB operation.

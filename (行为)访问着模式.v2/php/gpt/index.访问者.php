<?php
namespace php\gpt\visitor;
//访问者接口（Visitor Interface）：定义了对各种元素（Element）执行操作的访问者方法。
//具体访问者（Concrete Visitor）：实现访问者接口，并定义对每种类型元素（Element）访问所需进行的具体操作。
//元素接口（Element Interface）：定义了接收访问者的接口（通常是一个 accept 方法）。
//具体元素（Concrete Element）：实现元素接口，通常包含一个 accept 方法，该方法接受一个访问者对象作为参数，并将自身传递给该访问者的对应方法。
//对象结构（Object Structure）：包含元素对象的集合，可以是复合对象或简单的列表，提供一个方法允许访问者访问其内的每个元素。
interface VisitorInterface {
    public function visitConcreteElementA(ConcreteElementA $elementA);
    public function visitConcreteElementB(ConcreteElementB $elementB);
}

interface ElementInterface {
    public function accept(VisitorInterface $visitor);
}

class ConcreteElementA implements ElementInterface {
    public function accept(VisitorInterface $visitor) {
        $visitor->visitConcreteElementA($this);
    }

    public function operationA() {
        // 具体的操作
    }
}

class ConcreteElementB implements ElementInterface {
    public function accept(VisitorInterface $visitor) {
        $visitor->visitConcreteElementB($this);
    }

    public function operationB() {
        // 具体的操作
    }
}

class ConcreteVisitor implements VisitorInterface {
    public function visitConcreteElementA(ConcreteElementA $elementA) {
        $elementA->operationA();
    }

    public function visitConcreteElementB(ConcreteElementB $elementB) {
        $elementB->operationB();
    }
}

// 创建元素对象
$elementA = new ConcreteElementA();
$elementB = new ConcreteElementB();

// 创建访问者
$visitor = new ConcreteVisitor();

// 元素接受访问者
$elementA->accept($visitor);
$elementB->accept($visitor);

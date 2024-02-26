<?php
interface Product {
    public function use(): void;
}

class ConcreteProductA implements Product {
    public function use(): void {
        echo "Using ConcreteProductA\n";
    }
}

class ConcreteProductB implements Product {
    public function use(): void {
        echo "Using ConcreteProductB\n";
    }
}

abstract class Creator {
    // 工厂方法是抽象的，由子类实现
    abstract public function factoryMethod(): Product;

    // 可能需要依赖工厂方法的其他操作
    public function doSomething(): void {
        $product = $this->factoryMethod();
        echo "Creator: Working with ";
        $product->use();
    }
}

class ConcreteCreatorA extends Creator {
    public function factoryMethod(): Product {
        return new ConcreteProductA();
    }
}

class ConcreteCreatorB extends Creator {
    public function factoryMethod(): Product {
        return new ConcreteProductB();
    }
}

function clientCode(Creator $creator): void {
    $creator->doSomething();
}

// 通过具体创建者A来创建产品A
clientCode(new ConcreteCreatorA());

// 通过具体创建者B来创建产品B
clientCode(new ConcreteCreatorB());

//    Creator: Working with Using ConcreteProductA
//    Creator: Working with Using ConcreteProductB

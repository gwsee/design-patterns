<?php
namespace php\gpt\abstract_factory;

interface Product {
    public function getName();
}

class ProductA implements Product {
    public function getName() {
        return "Product A";
    }
}

class ProductB implements Product {
    public function getName() {
        return "Product B";
    }
}

interface AbstractFactory {
    public function createProductA();
    public function createProductB();
}

class ConcreteFactory1 implements AbstractFactory {
    public function createProductA() {
        return new ProductA();
    }

    public function createProductB() {
        return new ProductB();
    }
}

class ConcreteFactory2 implements AbstractFactory {
    // 可能创建不同版本的产品
    public function createProductA() {
        // 返回产品A的不同实现或版本
    }

    public function createProductB() {
        // 返回产品B的不同实现或版本
    }
}

// 客户端代码不关心是哪个具体工厂，只关心产品满足接口要求
function clientCode(AbstractFactory $factory) {
    $productA = $factory->createProductA();
    $productB = $factory->createProductB();

    echo $productA->getName() . PHP_EOL;
    echo $productB->getName() . PHP_EOL;
}

// 通过传入具体工厂类的实例可以创建不同的产品族
clientCode(new ConcreteFactory1());
//打印输出为
//
//    Product A
//    Product B
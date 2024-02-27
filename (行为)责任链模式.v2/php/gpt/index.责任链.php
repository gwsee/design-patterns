<?php
namespace php\gpt\responsibility_chain;

abstract class Handler {
    protected $successor = null;

    public function setSuccessor($successor) {
        $this->successor = $successor;
    }

    // 定义处理请求的方法，将实现留给具体的子类。
    abstract public function handleRequest($request);
}

class ConcreteHandler1 extends Handler {
    public function handleRequest($request) {
        if ($request == "R1") {
            // 处理请求
            echo "ConcreteHandler1 handled request: R1\n";
        } elseif ($this->successor != null) {
            // 将请求传递给下一个处理者
            $this->successor->handleRequest($request);
        }
    }
}

class ConcreteHandler2 extends Handler {
    public function handleRequest($request) {
        if ($request == "R2") {
            // 处理请求
            echo "ConcreteHandler2 handled request: R2\n";
        } elseif ($this->successor != null) {
            // 将请求传递给下一个处理者
            $this->successor->handleRequest($request);
        }
    }
}


// 设置责任链中对象的先后顺序
$handler1 = new ConcreteHandler1();
$handler2 = new ConcreteHandler2();

$handler1->setSuccessor($handler2);

// 发送请求
$handler1->handleRequest("R2"); // 这将输出：ConcreteHandler2 handled request: R2
$handler1->handleRequest("R1"); // 这将输出：ConcreteHandler1 handled request: R1
$handler1->handleRequest("R3"); // 这个请求不会被任何处理者处理

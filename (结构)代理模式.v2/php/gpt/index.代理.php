<?php
namespace php\gpt\proxy;

//Subject接口：定义了RealSubject和Proxy的共同接口，这样在任何使用RealSubject的地方都可以使用Proxy。
//RealSubject类：定义了代理所代表的实际对象。
//Proxy类：维护一个对RealSubject的引用，控制对RealSubject的访问，并可以负责创建和删除它。
//Client：使用Subject接口执行它的任务，无需关心工作是由RealSubject完成还是通过Proxy完成。

interface SubjectInterface {
    public function request();
}

class RealSubject implements SubjectInterface {
    public function request() {
        echo "RealSubject: Handling request.\n";
    }
}

class Proxy implements SubjectInterface {
    private $realSubject;

    public function __construct(RealSubject $realSubject) {
        $this->realSubject = $realSubject;
    }

    public function request() {
        if ($this->checkAccess()) {
            $this->realSubject->request();
            $this->logAccess();
        }
    }

    private function checkAccess() {
        // 检查访问权限
        echo "Proxy: Checking access prior to firing a real request.\n";
        return true; // 假设所有的客户端访问都是被允许的
    }

    private function logAccess() {
        // 记录请求日志
        echo "Proxy: Logging the time of request.\n";
    }
}

// Client code
echo "Client: Executing the client code with a real subject:\n";
$realSubject = new RealSubject();
$realSubject->request();

echo "\n";

echo "Client: Executing the same client code with a proxy:\n";
$proxy = new Proxy($realSubject);
$proxy->request();

//
//    Client: Executing the client code with a real subject:
//    RealSubject: Handling request.
//
//    Client: Executing the same client code with a proxy:
//    Proxy: Checking access prior to firing a real request.
//    RealSubject: Handling request.
//    Proxy: Logging the time of request.

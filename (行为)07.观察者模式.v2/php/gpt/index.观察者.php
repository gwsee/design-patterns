<?php
namespace php\gpt\observer;

ini_set('display_errors', 1);
error_reporting(E_ALL);

// 观察者接口
interface SplObserver {
    public function update(SplSubject $subject);
}

// 主题接口
interface SplSubject {
    public function attach(SplObserver $observer);
    public function detach(SplObserver $observer);
    public function notify();
}

// 具体的主题
class ConcreteSubject implements SplSubject {
    private $observers;
    private $state;

    public function __construct() {
        //如果你在使用带有特定命名空间的PHP文件中，你仍然可以直接使用PHP标准库（SPL）中的SplObserver和SplSubject接口，因为它们处于全局命名空间中。
        //在带有命名空间的类中引入它们时，可以通过在这些类名前加上反斜杠（\）来引用全局命名空间中的类，以确保PHP解析器在当前命名空间外查找它们。
        $this->observers = new \SplObjectStorage();
    }

    public function attach(SplObserver $observer) {
        $this->observers->attach($observer);
    }

    public function detach(SplObserver $observer) {
        $this->observers->detach($observer);
    }

    public function notify() {
        foreach ($this->observers as $observer) {
            $observer->update($this);
        }
    }

    public function setState($state) {
        $this->state = $state;
        $this->notify();
    }

    public function getState() {
        return $this->state;
    }
}

// 具体的观察者
class ConcreteObserver implements SplObserver {
    public function update(SplSubject $subject) {
        echo "Observer notified. New State: " . $subject->getState() . "\n";
    }
}

// 使用示例
$subject = new ConcreteSubject();

$observer1 = new ConcreteObserver();
$observer2 = new ConcreteObserver();

$subject->attach($observer1);
$subject->attach($observer2);

$subject->setState('state 1');
// Output:
// Observer notified. New State: state 1
// Observer notified. New State: state 1

$subject->detach($observer2);

$subject->setState('state 2');
// Output:
// Observer notified. New State: state 2
